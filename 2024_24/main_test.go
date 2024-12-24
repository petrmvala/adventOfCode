package main

import (
	"testing"
)

const inputSmall = `x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02
`

func TestGetZet(t *testing.T) {
	got := GetZet(inputSmall)
	want := 4
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
