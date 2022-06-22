package main

import (
	"math"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {

	type args struct {
		a   float64
		b   float64
		c   float64
		eps float64
	}

	tests := []struct {
		name    string
		args    args
		a       float64
		b       float64
		c       float64
		eps     float64
		want    []float64
		wantErr string
	}{
		{
			name: "no root",
			args: args{
				a:   1,
				b:   0,
				c:   1,
				eps: 0.01,
			},
			want:    []float64{},
			wantErr: "Дискриминант меньше 0!",
		},
		{
			name: "2 root ",
			args: args{
				a:   1,
				b:   0,
				c:   -1,
				eps: 0.01,
			},
			want:    []float64{1, -1},
			wantErr: "",
		},
		{
			name: "1 root ",
			args: args{
				a:   -0.2,                //1,
				b:   -0.05,               //2,
				c:   0.1,                 //1,
				eps: 0.08250000000000005, //0.01,
			},
			want:    []float64{-0.8430703308172536, -0.8430703308172536},
			wantErr: "",
		},
		{
			name: "a == 0",
			args: args{
				a:   0,
				b:   1,
				c:   1,
				eps: 0.01,
			},
			want:    []float64{},
			wantErr: "a == 0!",
		},
		{
			name: "NaN",
			args: args{
				a:   math.NaN(),
				b:   math.NaN(),
				c:   math.NaN(),
				eps: 0.01,
			},
			wantErr: "Коэффициенты не числа!",
		},
		{
			name: "+Inf",
			args: args{
				a:   math.Inf(1),
				b:   math.Inf(1),
				c:   math.Inf(1),
				eps: 0.01,
			},
			wantErr: "Коэффициенты положительные бесконечности!",
		},
		{
			name: "-Inf",
			args: args{
				a:   math.Inf(-1),
				b:   math.Inf(-1),
				c:   math.Inf(-1),
				eps: 0.01,
			},
			wantErr: "Коэффициенты отрицательные бесконечности!",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if (!math.IsNaN(tt.args.a) && !math.IsNaN(tt.args.b) && !math.IsNaN(tt.args.c)) && (!math.IsInf(tt.args.a, 1) && !math.IsInf(tt.args.b, 1) && !math.IsInf(tt.args.c, 1)) && (!math.IsInf(tt.args.a, -1) && !math.IsInf(tt.args.b, -1) && !math.IsInf(tt.args.c, -1)) {

				result, err := solve(tt.args.a, tt.args.b, tt.args.c, tt.args.eps)

				if err.Error() != tt.wantErr {
					t.Errorf("have = %v, want %v", err, tt.wantErr)
				} else if !reflect.DeepEqual(result, tt.want) {
					t.Errorf("have = %v, want %v", result, tt.want)
				}
			}
		})
	}
}
