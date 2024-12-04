package main

import (
	"reflect"
	"testing"
)

const sample = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
const small = `..X...
.SAMX.
.A..A.
XMAS.S
.X....
`

var (
	smallMat = [][]string{
		{".", ".", "X", ".", ".", "."},
		{".", "S", "A", "M", "X", "."},
		{".", "A", ".", ".", "A", "."},
		{"X", "M", "A", "S", ".", "S"},
		{".", "X", ".", ".", ".", "."},
	}
	smallLinesByRows = []string{
		"..X...",
		".SAMX.",
		".A..A.",
		"XMAS.S",
		".X....",
	}
	smallLinesByColumns = []string{
		"...X.",
		".SAMX",
		"XA.A.",
		".M.S.",
		".XA..",
		"...S.",
	}
	smallLinesFalling = []string{
		"XMAS",
		".A...",
		".S.S.",
		".AA.",
	}
	smallLinesRising = []string{
		"XAA.",
		".M.M.",
		"XA.X.",
		".SA.",
	}
)

func TestXMAS(t *testing.T) {
	got := XMAS(sample)
	want := 18
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestXMask(t *testing.T) {
	got := XMask(sample)
	want := 9
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGetLines(t *testing.T) {
	got := getLines(small)
	want := []string{}
	want = append(want, smallLinesByRows...)
	want = append(want, smallLinesByColumns...)
	want = append(want, smallLinesFalling...)
	want = append(want, smallLinesRising...)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLinesByRows(t *testing.T) {
	got := getLinesByRows(smallMat)
	want := smallLinesByRows
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLinesByColumns(t *testing.T) {
	got := getLinesByColumns(smallMat)
	want := smallLinesByColumns
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLinesDiagFalling(t *testing.T) {
	got := getLinesDiagFalling(smallMat)
	want := smallLinesFalling
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLinesDiagRising(t *testing.T) {
	got := getLinesDiagRising(smallMat)
	want := smallLinesRising
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestToMatrix(t *testing.T) {
	got := toMatrix(small)
	want := smallMat
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
