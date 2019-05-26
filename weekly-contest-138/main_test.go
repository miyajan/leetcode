package main

import (
	"reflect"
	"testing"
)

func Test_heightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				heights: []int{1, 1, 4, 2, 1, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightChecker(tt.args.heights); got != tt.want {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		X         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
				grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
				X:         3,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.X); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rearrangeBarcodes(t *testing.T) {
	type args struct {
		barcodes []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				barcodes: []int{1, 1, 1, 2, 2, 2},
			},
			want: []int{1, 2, 1, 2, 1, 2},
		},
		{
			args: args{
				barcodes: []int{1, 1, 1, 1, 2, 2, 3, 3},
			},
			want: []int{1, 3, 1, 2, 1, 2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeBarcodes(tt.args.barcodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeBarcodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prevPermOpt1(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				A: []int{3, 2, 1},
			},
			want: []int{3, 1, 2},
		},
		{
			args: args{
				A: []int{1, 1, 5},
			},
			want: []int{1, 1, 5},
		},
		{
			args: args{
				A: []int{1, 9, 4, 6, 7},
			},
			want: []int{1, 7, 4, 6, 9},
		},
		{
			args: args{
				A: []int{3, 1, 1, 3},
			},
			want: []int{1, 1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prevPermOpt1(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prevPermOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
