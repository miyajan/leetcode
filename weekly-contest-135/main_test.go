package main

import (
	"reflect"
	"testing"
)

func Test_isBoomerang(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				points: [][]int{
					{1, 1},
					{2, 3},
					{3, 2},
				},
			},
			want: true,
		},
		{
			args: args{
				points: [][]int{
					{1, 1},
					{2, 2},
					{3, 3},
				},
			},
			want: false,
		},
		{
			args: args{
				points: [][]int{
					{1, 1},
					{1, 1},
					{3, 3},
				},
			},
			want: false,
		},
		{
			args: args{
				points: [][]int{
					{1, 1},
					{1, 2},
					{1, 3},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBoomerang(tt.args.points); got != tt.want {
				t.Errorf("isBoomerang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bstToGst(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 2,
							Right: &TreeNode{
								Val: 3,
							},
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 7,
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			want: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 36,
					Left: &TreeNode{
						Val: 36,
					},
					Right: &TreeNode{
						Val: 35,
						Right: &TreeNode{
							Val: 33,
						},
					},
				},
				Right: &TreeNode{
					Val: 21,
					Left: &TreeNode{
						Val: 26,
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstToGst(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstToGst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minScoreTriangulation(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				A: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			args: args{
				A: []int{3, 7, 4, 5},
			},
			want: 144,
		},
		{
			args: args{
				A: []int{1, 3, 1, 4, 1, 5},
			},
			want: 13,
		},
		{
			args: args{
				A: []int{4, 3, 4, 3, 5},
			},
			want: 132,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minScoreTriangulation(tt.args.A); got != tt.want {
				t.Errorf("minScoreTriangulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numMovesStonesII(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				stones: []int{7, 4, 9},
			},
			want: []int{1, 2},
		},
		{
			args: args{
				stones: []int{6, 5, 4, 3, 10},
			},
			want: []int{2, 3},
		},
		{
			args: args{
				stones: []int{100, 101, 104, 102, 103},
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMovesStonesII(tt.args.stones); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStonesII() = %v, want %v", got, tt.want)
			}
		})
	}
}
