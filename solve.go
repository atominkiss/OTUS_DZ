package main

import (
	"fmt"
	"math"
	"math/big"
)

func solve(a, b, c float64) ([]float64, error) {

	aa := big.NewFloat(a).Cmp(big.NewFloat(0))

	if aa != 0 {

		if !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c) {

			disc := b*b - 4*a*c

			switch big.NewFloat(disc).Cmp(big.NewFloat(0)) {

			case -1:
				fmt.Println([]float64{})
				return []float64{}, fmt.Errorf("Дискриминант меньше 0!")
			case 0:
				x := (-b + math.Sqrt(disc)) / (2 * a)
				fmt.Println([]float64{x, x})
				return []float64{x, x}, fmt.Errorf("")
			case 1:
				x1 := (-b + math.Sqrt(disc)) / (2 * a)
				x2 := (-b - math.Sqrt(disc)) / (2 * a)
				fmt.Println([]float64{x1, x2})
				return []float64{x1, x2}, fmt.Errorf("")
			}
		} else {
			return []float64{}, fmt.Errorf("Коэффициенты не числа!")
		}
	} else {
		return nil, fmt.Errorf("a == 0!")
	}

	return []float64{}, fmt.Errorf("Неожиданное исключение!")
}
