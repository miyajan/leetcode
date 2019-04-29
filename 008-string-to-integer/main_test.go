package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "'42' -> 42",
			args: args{
				str: "42",
			},
			want: 42,
		},
		{
			name: "'   -42' -> -42",
			args: args{
				str: "   -42",
			},
			want: -42,
		},
		{
			name: "'4193 with words' -> 4193",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "'words and 987' -> 0",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "'-91283472332' -> -2147483648",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "'+' -> 0",
			args: args{
				str: "+",
			},
			want: 0,
		},
		{
			name: "'20000000000000000000' -> 2147483648",
			args: args{
				str: "20000000000000000000",
			},
			want: 2147483647,
		},
		{
			name: "'0-1' -> 0",
			args: args{
				str: "0-1",
			},
			want: 0,
		},
		{
			name: "'  0000000000012345678' -> 12345678",
			args: args{
				str: "  0000000000012345678",
			},
			want: 12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
