package main

import "math"

func main() {
	solve(1, 0, 1, 0.001)  // D < 0
	solve(1, 0, -1, 0.001) // D > 0
	solve(-0.2, -0.05, 0.1, 0.08250000000000005)
	//solve(0, 1, 2, 0.001)
	solve(math.NaN(), math.NaN(), math.NaN(), 0.001)
	solve(math.Inf(1), math.Inf(1), math.Inf(1), 0.001)
	solve(math.Inf(-1), math.Inf(-1), math.Inf(-1), 0.001)
}
