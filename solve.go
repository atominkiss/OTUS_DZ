package main

import (
	"fmt"
	"math"
)

func solve(a, b, c, eps float64) ([]float64, error) {

	D := math.Pow(b, 2) - 4*a*c

	if !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c) && (!math.IsInf(a, 1) && !math.IsInf(b, 1) && !math.IsInf(c, 1)) && (!math.IsInf(a, -1) && !math.IsInf(b, -1) && !math.IsInf(c, -1)) {

		if math.Abs(a-0.0) < eps {
			return []float64{}, fmt.Errorf("a == 0!")
		}

		switch {
		case D-0.0 < -eps:
			fmt.Println([]float64{})
			return []float64{}, fmt.Errorf("Дискриминант меньше 0!")

		case math.Abs(D-0.0) < eps:
			x := (-b + math.Sqrt(D)) / (2 * a)
			fmt.Println([]float64{x, x})
			return []float64{x, x}, fmt.Errorf("")

		default:
			x1 := (-b + math.Sqrt(D)) / (2 * a)
			x2 := (-b - math.Sqrt(D)) / (2 * a)
			fmt.Println([]float64{x1, x2})
			return []float64{x1, x2}, fmt.Errorf("")
		}
	} else {
		return []float64{}, fmt.Errorf("Коэффициенты не числа!")
	}
}
