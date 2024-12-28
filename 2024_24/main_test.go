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
	inputs, ops := LoadInput(inputSmall)
	got := GetZet(inputs, ops)
	want := 4
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGetOpByInput(t *testing.T) {
	_, ops := LoadInput(inputSmall)
	got, err := getOpByInputs(ops, "x00", "y00", "AND")
	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}
	gotRev, err := getOpByInputs(ops, "y00", "x00", "AND")
	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}
	want := "z00"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	if gotRev != want {
		t.Errorf("got %s, want %s", gotRev, want)
	}
}
