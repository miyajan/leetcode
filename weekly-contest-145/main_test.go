package main

import (
	"reflect"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
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
				arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
				arr2: []int{2, 1, 4, 3, 9, 6},
			},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			args: args{
				arr1: []int{28, 6, 22, 8, 44, 17},
				arr2: []int{22, 28, 8, 6},
			},
			want: []int{22, 28, 8, 6, 17, 44},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcaDeepestLeaves(t *testing.T) {
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
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: &TreeNode{
				Val: 4,
			},
		},
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcaDeepestLeaves(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lcaDeepestLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestWPI(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				hours: []int{9, 9, 6, 0, 6, 6, 9},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWPI(tt.args.hours); got != tt.want {
				t.Errorf("longestWPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestSufficientTeam(t *testing.T) {
	type args struct {
		req_skills []string
		people     [][]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				req_skills: []string{"java", "nodejs", "reactjs"},
				people:     [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
			},
			want: []int{2, 0},
		},
		{
			args: args{
				req_skills: []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
				people:     [][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}},
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSufficientTeam(tt.args.req_skills, tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestSufficientTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
