package main

import (
	"reflect"
	"testing"
)

func Test_countLength(t *testing.T) {
	type args struct {
		x1 []float64
		x2 []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				x1: []float64{1, 1},
				x2: []float64{1, 1},
			},
			want: 0,
		}, {
			args: args{
				x1: []float64{0, 0},
				x2: []float64{0, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLength(tt.args.x1, tt.args.x2); got != tt.want {
				t.Errorf("countLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clonePoints(t *testing.T) {
	type args struct {
		points [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			args: args{
				points: [][]float64{{0, 0}, {1, 1}, {-1, -1}},
			},
			want: [][]float64{{0, 0}, {1, 1}, {-1, -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeCopy(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clonePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
