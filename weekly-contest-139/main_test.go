package main

import (
	"reflect"
	"testing"
)

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str1: "ABCABC",
				str2: "ABC",
			},
			want: "ABC",
		},
		{
			args: args{
				str1: "ABABAB",
				str2: "AB",
			},
			want: "AB",
		},
		{
			args: args{
				str1: "LEET",
				str2: "CODE",
			},
			want: "",
		},
		{
			args: args{
				str1: "TAUXXTAUXXTAUXXTAUXXTAUXX",
				str2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			},
			want: "TAUXX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxEqualRowsAfterFlips(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				matrix: [][]int{
					{0, 1},
					{1, 1},
				},
			},
			want: 1,
		},
		{
			args: args{
				matrix: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: 2,
		},
		{
			args: args{
				matrix: [][]int{
					{0, 0, 0},
					{0, 0, 1},
					{1, 1, 0},
				},
			},
			want: 2,
		},
		{
			args: args{
				matrix: [][]int{
					{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1},
					{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
					{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1},
					{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
					{1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEqualRowsAfterFlips(tt.args.matrix); got != tt.want {
				t.Errorf("maxEqualRowsAfterFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addNegabinary(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr1: []int{1, 1, 1, 1, 1},
				arr2: []int{1, 0, 1},
			},
			want: []int{1, 0, 0, 0, 0},
		},
		{
			args: args{
				arr1: []int{1, 0},
				arr2: []int{1, 0},
			},
			want: []int{1, 1, 0, 0},
		},
		{
			args: args{
				arr1: []int{1},
				arr2: []int{1, 1},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addNegabinary(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addNegabinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numSubmatrixSumTarget(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				matrix: [][]int{
					{0,1,0},
					{1,1,1},
					{0,1,0},
				},
				target: 0,
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{
					{1,-1},
					{-1,1},
				},
				target: 0,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubmatrixSumTarget(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("numSubmatrixSumTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
