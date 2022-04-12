package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {

	type args struct {
		a float64
		b float64
		c float64
	}

	tests := []struct {
		name string
		args args
		a    float64
		b    float64
		c    float64
		want []float64
	}{
		{
			name: "no root",
			args: args{
				a: 1,
				b: 0,
				c: 1,
			},
			want: []float64{},
		},
		{
			name: "2 root ",
			args: args{
				a: 1,
				b: 0,
				c: -1,
			},
			want: []float64{1, -1},
		},
		{
			name: "1 root ",
			args: args{
				a: 1,
				b: 2,
				c: 1,
			},
			want: []float64{-1, -1},
		},
		{
			name: "exeption",
			args: args{
				a: 0,
				b: 2,
				c: 1,
			},
			want: []float64{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if big.NewFloat(tt.args.a).Cmp(big.NewFloat(0)) != 0 {
				result, err := solve(tt.args.a, tt.args.b, tt.args.c)
				if err != nil {

				}
				if !reflect.DeepEqual(result, tt.want) {
					t.Errorf("have = %v, want %v", result, tt.want)
				}
			}
		})
	}
}
