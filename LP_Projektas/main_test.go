package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_costFunction(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := costFunction(tt.args.points); got != tt.want {
				t.Errorf("costFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countLength(t *testing.T) {
	type args struct {
		x1 []int
		x2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				x1: []int{1, 1},
				x2: []int{1, 1},
			},
			want: 0,
		}, {
			args: args{
				x1: []int{0, 0},
				x2: []int{0, 5},
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

func Test_fillWithRandomPoints(t *testing.T) {
	tests := []struct {
		name string
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillWithRandomPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillWithRandomPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
