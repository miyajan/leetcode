package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "('PAYPALISHIRING', 3) => 'PAHNAPLSIIGYIR'",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "('PAYPALISHIRING', 4) => 'PINALSIGYAHRPI'",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "('', 0) => ''",
			args: args{
				s:       "",
				numRows: 0,
			},
			want: "",
		},
		{
			name: "('A', 1) => 'A'",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "('A', 2) => 'A'",
			args: args{
				s:       "A",
				numRows: 2,
			},
			want: "A",
		},
		{
			name: "('AB', 1) => 'AB'",
			args: args{
				s:       "AB",
				numRows: 1,
			},
			want: "AB",
		},
		{
			name: "('AB', 2) => 'AB'",
			args: args{
				s:       "AB",
				numRows: 2,
			},
			want: "AB",
		},
		{
			name: "('ABC', 2) => 'ACB'",
			args: args{
				s:       "ABC",
				numRows: 2,
			},
			want: "ACB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
