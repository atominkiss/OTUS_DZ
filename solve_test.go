package main

import (
	"math"
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
		name    string
		args    args
		a       float64
		b       float64
		c       float64
		want    []float64
		wantErr string
	}{
		{
			name: "no root",
			args: args{
				a: 1,
				b: 0,
				c: 1,
			},
			want:    []float64{},
			wantErr: "Дискриминант меньше 0!",
		},
		{
			name: "2 root ",
			args: args{
				a: 1,
				b: 0,
				c: -1,
			},
			want:    []float64{1, -1},
			wantErr: "",
		},
		{
			name: "1 root ",
			args: args{
				a: 1,
				b: 2,
				c: 1,
			},
			want:    []float64{-1, -1},
			wantErr: "",
		},
		{
			name: "a == 0",
			args: args{
				a: 0,
				b: 1,
				c: 1,
			},
			want:    []float64{},
			wantErr: "a == 0!",
		},
		{
			name: "NaN",
			args: args{
				a: math.NaN(),
				b: math.NaN(),
				c: math.NaN(),
			},
			wantErr: "Коэффициенты не числа!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !math.IsNaN(tt.args.a) && !math.IsNaN(tt.args.b) && !math.IsNaN(tt.args.c) {

				if big.NewFloat(tt.args.a).Cmp(big.NewFloat(0)) != 0 {

					result, err := solve(tt.args.a, tt.args.b, tt.args.c)

					if err.Error() != tt.wantErr {
						t.Errorf("have = %v, want %v", err, tt.wantErr)
					} else if !reflect.DeepEqual(result, tt.want) {
						t.Errorf("have = %v, want %v", result, tt.want)
					}
				}
			}
		})
	}
}
