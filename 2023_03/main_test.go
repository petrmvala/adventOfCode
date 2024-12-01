package main

import (
	"fmt"
	"testing"
)

var schematic = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestSum(t *testing.T) {
	got := sum(schematic)
	want := 4361
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAngles(t *testing.T) {
	s := "...\n.1.\n...\n"
	for i := 0; i < 12; i++ {
		if i == 3 || i == 5 || i == 7 || i == 11 {
			continue
		}
		ss := s[:i] + "*" + s[i+1:]
		if sum(ss) != 1 {
			fmt.Println(ss)
			t.Error("value not counted")
		}
	}
}
