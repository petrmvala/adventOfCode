package main

import (
	"reflect"
	"testing"
)

const sample = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestFirstStar(t *testing.T) {
	mat := toMatrix(sample)
	got := FirstStar(mat)
	want := 14
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGetAntennas(t *testing.T) {
	mat := toMatrix(sample)
	got := getAntennas(mat)
	want := Antennas{
		"0": []Position{
			{x: 1, y: 8},
			{x: 2, y: 5},
			{x: 3, y: 7},
			{x: 4, y: 4},
		},
		"A": []Position{
			{x: 5, y: 6},
			{x: 8, y: 8},
			{x: 9, y: 9},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
