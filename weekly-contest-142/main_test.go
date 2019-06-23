package main

import (
	"reflect"
	"testing"
)

func Test_sampleStats(t *testing.T) {
	type args struct {
		count []int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			args: args{
				count: []int{0, 1, 3, 4},
			},
			want: []float64{1, 3, 2.375, 2.5, 3},
		},
		{
			args: args{
				count: []int{0, 4, 3, 2, 2},
			},
			want: []float64{1, 4, 2.1818181818181817, 2, 1},
		},
		{
			args: args{
				count: []int{0},
			},
			want: []float64{0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleStats(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sampleStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				trips: [][]int{
					{2, 1, 5},
					{3, 3, 7},
				},
				capacity: 4,
			},
			want: false,
		},
		{
			args: args{
				trips: [][]int{
					{2, 1, 5},
					{3, 3, 7},
				},
				capacity: 5,
			},
			want: true,
		},
		{
			args: args{
				trips: [][]int{
					{2, 1, 5},
					{3, 5, 7},
				},
				capacity: 3,
			},
			want: true,
		},
		{
			args: args{
				trips: [][]int{
					{3, 2, 7},
					{3, 7, 9},
					{8, 3, 9},
				},
				capacity: 11,
			},
			want: true,
		},
		{
			args: args{
				trips: [][]int{
					{5, 1, 3},
					{5, 3, 6},
					{1, 1, 2},
					{9, 2, 5},
					{6, 5, 7},
					{2, 3, 6},
					{9, 3, 5},
				},
				capacity: 26,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findInMountainArray(t *testing.T) {
	type args struct {
		target      int
		mountainArr *MountainArray
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				target: 3,
				mountainArr: &MountainArray{
					arr: []int{1, 2, 3, 4, 5, 3, 1},
				},
			},
			want: 2,
		},
		{
			args: args{
				target: 3,
				mountainArr: &MountainArray{
					arr: []int{0, 1, 2, 4, 2, 1},
				},
			},
			want: -1,
		},
		{
			args: args{
				target: 5,
				mountainArr: &MountainArray{
					arr: []int{3, 5, 3, 2, 0},
				},
			},
			want: 1,
		},
		{
			args: args{
				target: 5,
				mountainArr: &MountainArray{
					arr: []int{1, 5, 3},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInMountainArray(tt.args.target, tt.args.mountainArr); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_braceExpansionII(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				expression: "{a,b}{c{d,e}}",
			},
			want: []string{"acd", "ace", "bcd", "bce"},
		},
		{
			args: args{
				expression: "{{a,z},a{b,c},{ab,z}}",
			},
			want: []string{"a", "ab", "ac", "z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := braceExpansionII(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceExpansionII() = %v, want %v", got, tt.want)
			}
		})
	}
}
