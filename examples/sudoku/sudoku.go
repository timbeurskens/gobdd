package main

import (
	"flag"
	"fmt"
	"github.com/timbeurskens/gobdd"
	"github.com/timbeurskens/gobdd/algorithm"
	"github.com/timbeurskens/gobdd/numerics"
	. "github.com/timbeurskens/gobdd/operators"
	"github.com/timbeurskens/gobdd/operators/bdd"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"time"
)

const (
	K = 2
	N = K * K
)

var (
	bits        = 1 + int(math.Ceil(math.Log2(N)))
	useBdd      = flag.Bool("bdd", false, "enable bdd solver")
	useCdcl     = flag.Bool("cdcl", false, "enable cdcl solver")
	cpuProfile  = flag.String("cpuprofile", "", "enable cpu profiler")
	heapProfile = flag.String("heapprofile", "", "enable heap profiler")
)

type SudokuHint [N][N]uint
type SudokuBoard [N][N]numerics.Number
type Row [N]numerics.Number

var (
	//HardSudokuExample = SudokuHint{
	//	{0, 2, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 6, 0, 0, 0, 0, 3},
	//	{0, 7, 4, 0, 8, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 3, 0, 0, 3},
	//	{0, 8, 0, 0, 4, 0, 0, 1, 0},
	//	{6, 0, 0, 5, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 1, 0, 7, 8, 0},
	//	{5, 0, 0, 0, 0, 9, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 4, 0},
	//}

	TwoTwoExample = SudokuHint{
		{3, 4, 1, 0},
		{0, 2, 0, 0},
		{0, 0, 2, 0},
		{0, 1, 4, 3},
	}
)

func (board *SudokuBoard) Row(i int) (result Row) {
	for j := range board[i] {
		result[j] = board[i][j]
	}
	return
}

func (board *SudokuBoard) Col(j int) (result Row) {
	for i := range board {
		result[i] = board[i][j]
	}
	return
}

func (board *SudokuBoard) Tile(k int) (result Row) {
	iStart, jStart := k/K, k%K
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			result[K*i+j] = board[iStart+i*K][jStart+j*K]
		}
	}
	return
}

func makeSingleNumber(num numerics.Number) Expression {
	exprs := make([]Expression, N)
	for i := range exprs {
		exprs[i] = numerics.Equals(num, numerics.Constant(uint(i+1), bits))
	}
	return Or(exprs...)
}

func makeUniqueRow(row *Row) Expression {
	exprs := make([]Expression, 0, 1+(N*N)/2)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			// comparing two constants does not add any information (assuming the input is valid)
			if !row[i].IsConstant() || !row[j].IsConstant() {
				exprs = append(exprs, Not(numerics.Equals(row[i], row[j])))
			}
		}
	}
	if len(exprs) == 0 {
		return Cons(true)
	}
	return And(exprs...)
}

func makeSudoku(hint *SudokuHint) (*SudokuBoard, Expression) {
	var board SudokuBoard
	exprs := make([]Expression, 0, N*N*N)

	for i := range hint {
		for j := range hint[i] {
			if hint[i][j] > 0 {
				board[i][j] = numerics.Constant(hint[i][j], bits)
			} else {
				board[i][j] = numerics.Variable(bits)
				exprs = append(exprs, makeSingleNumber(board[i][j]))
			}
		}
	}

	for i := range board {
		row, col, tile := board.Row(i), board.Col(i), board.Tile(i)
		rowExpr, colExpr, tileExpr := makeUniqueRow(&row), makeUniqueRow(&col), makeUniqueRow(&tile)
		exprs = append(exprs, And(rowExpr, colExpr, tileExpr))
	}

	return &board, And(exprs...)
}

func solveBDD(expr Expression) (Model, bool) {
	expr = algorithm.PruneUnary(expr)

	log.Println("Size of expression:", Size(expr))

	tree := algorithm.FromExpression(expr)

	log.Println("Size of tree:", Size(tree))

	if bdd.Sat(tree) {
		if model, ok := bdd.FindModel(tree); ok {
			return model, true
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}

func solveCDCL(expr Expression) (Model, bool) {
	nnf := algorithm.NNF(expr)

	cnf := algorithm.TransformTseitin(nnf)

	log.Printf("Solving %d clauses", len(cnf))

	varCount := len(Variables(cnf))

	log.Printf("%d variables", varCount)

	if sat, m := algorithm.CDCL(cnf); sat {
		if model, ok := bdd.FindModel(m); ok {
			return model, true
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}

func main() {
	flag.Parse()

	if *cpuProfile != "" {
		if fProfile, err := os.Create(*cpuProfile); err != nil {
			log.Fatal(err)
		} else {
			defer fProfile.Close()
			if err = pprof.StartCPUProfile(fProfile); err != nil {
				log.Fatal(err)
			} else {
				defer pprof.StopCPUProfile()
			}
		}
	}

	if *heapProfile != "" {
		if fProfile, err := os.Create(*heapProfile); err != nil {
			log.Fatal(err)
		} else {
			defer fProfile.Close()
			timer := time.NewTimer(1 * time.Second)
			go func() {
				for _ = range timer.C {
					pprof.WriteHeapProfile(fProfile)
				}
			}()
		}
	}

	log.Printf("Using %d bits", bits)

	//defaultHint := &HardSudokuExample
	defaultHint := &TwoTwoExample
	board, expr := makeSudoku(defaultHint)

	log.Printf(gobdd.PrintExpressiontree(expr))

	var model Model
	var ok bool

	if *useBdd {
		// BDD solver requires a lot of memory for a 9x9 sudoku, 4x4 works
		model, ok = solveBDD(expr)
	} else if *useCdcl {
		// CDCL requires less memory, but is less efficient in solving problems
		model, ok = solveCDCL(expr)
	} else {
		log.Fatal("Enable at least one solving strategy")
	}

	if !ok {
		log.Fatal("Failed to solve or find solution")
	}

	for i := range defaultHint {
		for j := range defaultHint[i] {
			if defaultHint[i][j] > 0 {
				fmt.Printf("%d\t", defaultHint[i][j])
			} else {
				hintResolv, _ := board[i][j].Resolve(model)
				fmt.Printf("%d\t", hintResolv)
			}
		}
		fmt.Println()
	}
}
