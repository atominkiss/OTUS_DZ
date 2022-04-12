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
			name: "NaN",
			args: args{
				a: math.NaN(),
				b: math.NaN(),
				c: math.NaN(),
			},
			wantErr: "Дискриминант меньше 0!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if big.NewFloat(tt.args.a).Cmp(big.NewFloat(0)) != 0 {
				result, err := solve(tt.args.a, tt.args.b, tt.args.c)
				if err.Error() != tt.wantErr {
					t.Errorf("have = %v, want %v", err, tt.wantErr)
				} else if !reflect.DeepEqual(result, tt.want) {
					t.Errorf("have = %v, want %v", result, tt.want)
				}
			}
		})
	}
}
