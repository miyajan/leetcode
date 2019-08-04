package main

import (
	"reflect"
	"testing"
)

func Test_movesToMakeZigzag(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{9, 6, 1, 6, 2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{2, 7, 10, 9, 8, 9},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{10, 4, 4, 10, 10, 6, 2, 3},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToMakeZigzag(tt.args.nums); got != tt.want {
				t.Errorf("movesToMakeZigzag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_btreeGameWinningMove(t *testing.T) {
	type args struct {
		root *TreeNode
		n    int
		x    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 8,
							},
							Right: &TreeNode{
								Val: 9,
							},
						},
						Right: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 10,
							},
							Right: &TreeNode{
								Val: 11,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				n: 11,
				x: 3,
			},
			want: true,
		},
		{
			args: args{
				root: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 3,
							Right: &TreeNode{
								Val: 2,
							},
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 1,
								},
							},
						},
					},
					Right: &TreeNode{
						Val: 7,
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				n: 9,
				x: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := btreeGameWinningMove(tt.args.root, tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("btreeGameWinningMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	sa := Constructor(3)
	sa.Set(0, 5)
	snapId := sa.Snap()
	if !reflect.DeepEqual(snapId, 0) {
		t.Errorf("sa.Snap() = %v, want %v", snapId, 0)
	}
	sa.Set(0, 6)
	got := sa.Get(0, 0)
	if !reflect.DeepEqual(got, 5) {
		t.Errorf("sa.Get(0, 0) = %v, want %v", got, 5)
	}
}

func Test_longestDecomposition(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				text: "ghiabcdefhelloadamhelloabcdefghi",
			},
			want: 7,
		},
		{
			args: args{
				text: "merchant",
			},
			want: 1,
		},
		{
			args: args{
				text: "antaprezatepzapreanta",
			},
			want: 11,
		},
		{
			args: args{
				text: "aaa",
			},
			want: 3,
		},
		{
			args: args{
				text: "elvtoelvto",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDecomposition(tt.args.text); got != tt.want {
				t.Errorf("longestDecomposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
