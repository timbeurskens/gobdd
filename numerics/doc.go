// Package numerics contains arithmetic operations on numeric types, suitable for solving in the SAT-solvers implemented in the algorithms package.
// This version currently only supports numerics of the "Naturals" class: whole numbers greater than, and including zero.
// Integers, fractionals and fixed-point classes could be added in later versions.
// Solving numeric equations requires a significant amount of computing power.
// The current CDCL implementation is not suitable for computing multiplications within reasonable time windows.
package numerics
