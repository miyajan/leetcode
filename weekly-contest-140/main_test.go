package main

import (
	"reflect"
	"testing"
)

func Test_findOcurrences(t *testing.T) {
	type args struct {
		text   string
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				text:   "alice is a good girl she is a good student",
				first:  "a",
				second: "good",
			},
			want: []string{"girl", "student"},
		},
		{
			args: args{
				text:   "we will we will rock you",
				first:  "we",
				second: "will",
			},
			want: []string{"we", "rock"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOcurrences(tt.args.text, tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOcurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sufficientSubset(t *testing.T) {
	type args struct {
		root  *TreeNode
		limit int
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
							Val: -99,
							Left: &TreeNode{
								Val: -99,
							},
							Right: &TreeNode{
								Val: -99,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: -99,
							Left: &TreeNode{
								Val: 12,
							},
							Right: &TreeNode{
								Val: 13,
							},
						},
						Right: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: -99,
							},
							Right: &TreeNode{
								Val: 14,
							},
						},
					},
				},
			},
			want: &TreeNode{
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
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
						Right: &TreeNode{
							Val: 14,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sufficientSubset(tt.args.root, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sufficientSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numTilePossibilities(t *testing.T) {
	type args struct {
		tiles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				tiles: "AAB",
			},
			want: 8,
		},
		{
			args: args{
				tiles: "AAABBC",
			},
			want: 188,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilePossibilities(tt.args.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestSubsequence(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				text: "cdadabcc",
			},
			want: "adbc",
		},
		{
			args: args{
				text: "abcd",
			},
			want: "abcd",
		},
		{
			args: args{
				text: "ecbacba",
			},
			want: "eacb",
		},
		{
			args: args{
				text: "leetcode",
			},
			want: "letcod",
		},
		{
			args: args{
				text: "abdaddcdbdbadcaddbcddaaccdcdca",
			},
			want: "abcd",
		},
		{
			args: args{
				text: "ddeeeccdce",
			},
			want: "cde",
		},
		{
			args: args{
				text: "aaeeeceabbfdaefdeefbabaabbdbaadbebfaabfadcaacddebfdbeaceffaaadcaeddbadebdebccbcbccefeaffbfdcdaefeefeffefcfbbbdabdbddcaaeaacecefbbbeaacdafdfbcdfcbedaff",
			},
			want: "abcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.args.text); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
