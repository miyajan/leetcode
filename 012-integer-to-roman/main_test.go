package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "3 -> III",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "4 -> IV",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "9 -> IX",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "58 -> LVIII",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "1994 -> MCMXCIV",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
