package main

import (
	"reflect"
	"testing"
)

func Test_isRobotBounded(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				instructions: "GGLLGG",
			},
			want: true,
		},
		{
			args: args{
				instructions: "GG",
			},
			want: false,
		},
		{
			args: args{
				instructions: "GL",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRobotBounded(tt.args.instructions); got != tt.want {
				t.Errorf("isRobotBounded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gardenNoAdj(t *testing.T) {
	type args struct {
		N     int
		paths [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				N: 3,
				paths: [][]int{
					{1, 2},
					{2, 3},
					{3, 1},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			args: args{
				N: 4,
				paths: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			want: []int{1, 2, 1, 2},
		},
		{
			args: args{
				N: 4,
				paths: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 1},
					{1, 3},
					{2, 4},
				},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gardenNoAdj(tt.args.N, tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gardenNoAdj() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSumAfterPartitioning(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				A: []int{1, 15, 7, 9, 2, 5, 10},
				K: 3,
			},
			want: 84,
		},
		{
			args: args{
				A: []int{1},
				K: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumAfterPartitioning(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("maxSumAfterPartitioning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestDupSubstring(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				S: "banana",
			},
			want: "ana",
		},
		{
			args: args{
				S: "abcd",
			},
			want: "",
		},
		{
			args: args{
				S: "akyjaakyj",
			},
			want: "akyj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDupSubstring(tt.args.S); got != tt.want {
				t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
