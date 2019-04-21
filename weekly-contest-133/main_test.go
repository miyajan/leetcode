package main

import (
	"reflect"
	"testing"
)

func Test_allCellsDistOrder(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "R = 1, C = 2, r0 = 0, c0 = 0 -> [[0,0],[0,1]]",
			args: args{
				R:  1,
				C:  2,
				r0: 0,
				c0: 0,
			},
			want: [][]int{
				{0, 0},
				{0, 1},
			},
		},
		{
			name: "R = 2, C = 2, r0 = 0, c0 = 1 -> [[0,1],[0,0],[1,1],[1,0]]",
			args: args{
				R:  2,
				C:  2,
				r0: 0,
				c0: 1,
			},
			want: [][]int{
				{0, 1},
				{0, 0},
				{1, 1},
				{1, 0},
			},
		},
		{
			name: "R = 2, C = 3, r0 = 1, c0 = 2 -> [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]",
			args: args{
				R:  2,
				C:  3,
				r0: 1,
				c0: 2,
			},
			want: [][]int{
				{1, 2},
				{0, 2},
				{1, 1},
				{0, 1},
				{1, 0},
				{0, 0},
			},
		},
		{
			name: "R = 0, C = 0, r0 = 0, c0 = 0 -> []",
			args: args{
				R:  0,
				C:  0,
				r0: 0,
				c0: 0,
			},
			want: [][]int{},
		},
		{
			name: "R = 1, C = 0, r0 = 0, c0 = 0 -> []",
			args: args{
				R:  1,
				C:  0,
				r0: 0,
				c0: 0,
			},
			want: [][]int{},
		},
		{
			name: "R = 0, C = 1, r0 = 0, c0 = 0 -> []",
			args: args{
				R:  0,
				C:  1,
				r0: 0,
				c0: 0,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoCitySchedCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[10,20],[30,200],[400,50],[30,20]] -> 110",
			args: args{
				costs: [][]int{
					{10, 20},
					{30, 200},
					{400, 50},
					{30, 20},
				},
			},
			want: 110,
		},
		{
			name: "[[10,20],[30,200]] -> 50",
			args: args{
				costs: [][]int{
					{10, 20},
					{30, 200},
				},
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoCitySchedCost(tt.args.costs); got != tt.want {
				t.Errorf("twoCitySchedCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSumTwoNoOverlap(t *testing.T) {
	type args struct {
		A []int
		L int
		M int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A = [0,6,5,2,2,5,1,9,4], L = 1, M = 2 -> 20",
			args: args{
				A: []int{0, 6, 5, 2, 2, 5, 1, 9, 4},
				L: 1,
				M: 2,
			},
			want: 20,
		},
		{
			name: "A = [3,8,1,3,2,1,8,9,0], L = 3, M = 2 -> 29",
			args: args{
				A: []int{3, 8, 1, 3, 2, 1, 8, 9, 0},
				L: 3,
				M: 2,
			},
			want: 29,
		},
		{
			name: "A = [2,1,5,6,0,9,5,0,3,8], L = 4, M = 3 -> 31",
			args: args{
				A: []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
				L: 4,
				M: 3,
			},
			want: 31,
		},
		{
			name: "A = [2,1], L = 1, M = 1 -> 3",
			args: args{
				A: []int{2, 1},
				L: 1,
				M: 1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumTwoNoOverlap(tt.args.A, tt.args.L, tt.args.M); got != tt.want {
				t.Errorf("maxSumTwoNoOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamChecker_Query(t *testing.T) {
	type fields struct {
		words []string
	}
	type io struct {
		letter byte
		want   bool
	}
	type args struct {
		ios []io
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "words = ['cd','f','kl'], ['a' -> false, 'b' -> false, 'c' -> false, 'd' -> true, 'e' -> false, 'f' -> true, 'g' -> false, 'h' -> false, 'i' -> false, 'j' -> false, 'k' -> false, 'l' -> true]",
			fields: fields{
				words: []string{"cd", "f", "kl"},
			},
			args: args{
				ios: []io{
					{
						letter: 'a',
						want:   false,
					},
					{
						letter: 'b',
						want:   false,
					},
					{
						letter: 'c',
						want:   false,
					},
					{
						letter: 'd',
						want:   true,
					},
					{
						letter: 'e',
						want:   false,
					},
					{
						letter: 'f',
						want:   true,
					},
					{
						letter: 'g',
						want:   false,
					},
					{
						letter: 'h',
						want:   false,
					},
					{
						letter: 'i',
						want:   false,
					},
					{
						letter: 'j',
						want:   false,
					},
					{
						letter: 'k',
						want:   false,
					},
					{
						letter: 'l',
						want:   true,
					},
				},
			},
		},
		{
			name: "words = ['abaa','abaab','aabbb','bab','ab', ['a' -> false, 'a' -> false, 'b' -> true, 'b' -> false, 'b' -> true]",
			fields: fields{
				words: []string{"abaa", "abaab", "aabbb", "bab", "ab"},
			},
			args: args{
				ios: []io{
					{
						letter: 'a',
						want:   false,
					},
					{
						letter: 'a',
						want:   false,
					},
					{
						letter: 'b',
						want:   true,
					},
					{
						letter: 'b',
						want:   false,
					},
					{
						letter: 'b',
						want:   true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.words)
			for i := 0; i < len(tt.args.ios); i++ {
				if got := this.Query(tt.args.ios[i].letter); got != tt.args.ios[i].want {
					t.Errorf("StreamChecker.Query(%c) = %v, want %v", tt.args.ios[i].letter, got, tt.args.ios[i].want)
				}
			}
		})
	}
}
