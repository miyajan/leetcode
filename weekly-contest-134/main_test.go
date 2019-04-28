package main

import (
	"reflect"
	"testing"
)

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A = [1,4,2], B = [1,2,4] -> 2",
			args: args{
				A: []int{1, 4, 2},
				B: []int{1, 2, 4},
			},
			want: 2,
		},
		{
			name: "A = [2,5,1,2,5], B = [10,5,2,1,5,2] -> 3",
			args: args{
				A: []int{2, 5, 1, 2, 5},
				B: []int{10, 5, 2, 1, 5, 2},
			},
			want: 3,
		},
		{
			name: "A = [1,3,7,1,7,5], B = [1,9,2,5,1] -> 2",
			args: args{
				A: []int{1, 3, 7, 1, 7, 5},
				B: []int{1, 9, 2, 5, 1},
			},
			want: 2,
		},
		{
			name: "A = [1,4,2,5], B = [1,2,5,4] -> 3",
			args: args{
				A: []int{1, 4, 2, 5},
				B: []int{1, 2, 5, 4},
			},
			want: 3,
		},
		{
			name: "Long Arguments",
			args: args{
				A: []int{15, 14, 1, 7, 15, 1, 12, 18, 9, 15, 1, 20, 18, 15, 16, 18, 11, 8, 11, 18, 11, 11, 17, 20, 16, 20, 15, 15, 9, 18, 16, 4, 16, 1, 13, 10, 10, 20, 4, 18, 17, 3, 8, 1, 8, 19, 14, 10, 10, 12},
				B: []int{12, 8, 17, 4, 2, 18, 16, 10, 11, 12, 7, 1, 8, 16, 4, 14, 12, 18, 18, 19, 19, 1, 11, 18, 1, 6, 12, 17, 6, 19, 10, 5, 11, 16, 6, 17, 12, 1, 9, 3, 19, 2, 18, 18, 2, 4, 11, 11, 14, 9, 20, 19, 2, 20, 9, 15, 8, 7, 8, 6, 19, 12, 4, 11, 18, 18, 1, 6, 9, 17, 13, 19, 5, 4, 14, 9, 11, 15, 2, 5, 4, 1, 10, 11, 6, 4, 9, 7, 11, 7, 3, 8, 11, 12, 4, 19, 12, 17, 14, 18},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_colorBorder(t *testing.T) {
	type args struct {
		grid  [][]int
		r0    int
		c0    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "grid = [[1,1],[1,2]], r0 = 0, c0 = 0, color = 3 -> [[3, 3], [3, 2]]",
			args: args{
				grid: [][]int{
					{1, 1},
					{1, 2},
				},
				r0:    0,
				c0:    0,
				color: 3,
			},
			want: [][]int{
				{3, 3},
				{3, 2},
			},
		},
		{
			name: "grid = [[1,2,2],[2,3,2]], r0 = 0, c0 = 1, color = 3 -> [[1, 3, 3], [2, 3, 3]]",
			args: args{
				grid: [][]int{
					{1, 2, 2},
					{2, 3, 2},
				},
				r0:    0,
				c0:    1,
				color: 3,
			},
			want: [][]int{
				{1, 3, 3},
				{2, 3, 3},
			},
		},
		{
			name: "grid = [[1,1,1],[1,1,1],[1,1,1]], r0 = 1, c0 = 1, color = 2 -> [[2, 2, 2], [2, 1, 2], [2, 2, 2]]",
			args: args{
				grid: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				r0:    1,
				c0:    1,
				color: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 1, 2},
				{2, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorBorder(tt.args.grid, tt.args.r0, tt.args.c0, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("colorBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEscapePossible(t *testing.T) {
	type args struct {
		blocked [][]int
		source  []int
		target  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "blocked = [[0,1],[1,0]], source = [0,0], target = [0,2] -> false",
			args: args{
				blocked: [][]int{
					{0, 1},
					{1, 0},
				},
				source: []int{0, 0},
				target: []int{0, 2},
			},
			want: false,
		},
		{
			name: "blocked = [], source = [0,0], target = [999999,999999] -> true",
			args: args{
				blocked: [][]int{},
				source:  []int{0, 0},
				target:  []int{999999, 999999},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEscapePossible(tt.args.blocked, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("isEscapePossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numMovesStones(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "a = 1, b = 2, c = 5 -> [1, 2]",
			args: args{
				a: 1,
				b: 2,
				c: 5,
			},
			want: []int{1, 2},
		},
		{
			name: "a = 4, b = 3, c = 2 -> [0, 0]",
			args: args{
				a: 4,
				b: 3,
				c: 2,
			},
			want: []int{0, 0},
		},
		{
			name: "a = 3, b = 5, c = 1 -> [1, 2]",
			args: args{
				a: 3,
				b: 5,
				c: 1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMovesStones(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
