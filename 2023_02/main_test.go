package main

import (
	"reflect"
	"testing"
)

var gamesA = []struct {
	line     string
	valid    bool
	requires requires
}{
	{"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true, requires{red: 4, green: 2, blue: 6}},
	{"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true, requires{red: 1, green: 3, blue: 4}},
	{"8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false, requires{red: 20, green: 13, blue: 6}},
	{"1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false, requires{red: 14, green: 3, blue: 15}},
	{"6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true, requires{red: 6, green: 3, blue: 2}},
}

func TestGameValid(t *testing.T) {
	for _, tt := range gamesA {
		got := gameValid(tt.line)
		if got != tt.valid {
			t.Errorf("got %v, want %v", got, tt.valid)
		}
	}
}

func TestGameRequires(t *testing.T) {
	for _, tt := range gamesA {
		got := gameRequires(tt.line)
		if !reflect.DeepEqual(got, tt.requires) {
			t.Errorf("line: %q, got %+v, want %+v", tt.line, got, tt.requires)
		}
	}
}
