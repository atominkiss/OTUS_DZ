package main

import (
	"fmt"
	"math"
	"math/big"
)

func solve(a, b, c float64) ([]float64, error) {
	disc := b*b - 4*a*c
	switch big.NewFloat(disc).Cmp(big.NewFloat(0)) {
	//case -1:
	//	fmt.Println([]float64{})
	//	return []float64{}, fmt.Errorf("Дискриминант меньше 0!")
	case 0:
		x := (-b + math.Sqrt(disc)) / (2 * a)
		fmt.Println([]float64{x, x})
		return []float64{x, x}, nil
	case 1:
		x1 := (-b + math.Sqrt(disc)) / (2 * a)
		x2 := (-b - math.Sqrt(disc)) / (2 * a)
		fmt.Println([]float64{x1, x2})
		return []float64{x1, x2}, nil
	}

	return []float64{}, fmt.Errorf("Дискриминант меньше 0!")
}
