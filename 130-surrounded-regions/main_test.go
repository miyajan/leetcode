package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Simple X",
			args: args{
				board: [][]byte{
					[]byte("X"),
				},
			},
			want: [][]byte{
				[]byte("X"),
			},
		},
		{
			name: "Simple O",
			args: args{
				board: [][]byte{
					[]byte("O"),
				},
			},
			want: [][]byte{
				[]byte("O"),
			},
		},
		{
			name: "Flip surrounded resion",
			args: args{
				board: [][]byte{
					[]byte("XXX"),
					[]byte("XOX"),
					[]byte("XXX"),
				},
			},
			want: [][]byte{
				[]byte("XXX"),
				[]byte("XXX"),
				[]byte("XXX"),
			},
		},
		{
			name: "Example",
			args: args{
				board: [][]byte{
					[]byte("XXXX"),
					[]byte("XOOX"),
					[]byte("XXOX"),
					[]byte("XOXX"),
				},
			},
			want: [][]byte{
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XOXX"),
			},
		},
		{
			name: "Empty",
			args: args{
				board: [][]byte{},
			},
			want: [][]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			for y, line := range tt.args.board {
				for x, actual := range line {
					if actual != tt.want[y][x] {
						t.Errorf("expected board[%d][%d] = %v, actual board[%d][%d] = %v",
							y, x, tt.want[y][x], y, x, actual)
					}
				}
			}
		})
	}
}
