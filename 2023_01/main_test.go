package main

import "testing"

func TestLineMatch(t *testing.T) {

	cases := []struct {
		got  string
		want int
	}{
		{"asldjkf12asldfj", 12},
		{"12asldfj", 12},
		{"asldjkf1sldf2asldfj", 12},
		{"asldjkf12", 12},
		{"12", 12},
		{"1lskdjf2", 12},
		{"1", 11},
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"7pqrstsizeron", 77},
		{"one", 11},
		{"xhzone6", 16},
		{"gkhpvkeightsixronethreeone2", 82},
		{"5eight34sckhhxrtwonem", 51},
	}

	for _, tt := range cases {
		got := lineMatch(tt.got)
		if got != tt.want {
			t.Errorf("got %d, want %d", got, tt.want)
		}
	}
}
