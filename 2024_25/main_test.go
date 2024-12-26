package main

import (
	"reflect"
	"testing"
)

const schemata = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####
`

func TestGetSchematics(t *testing.T) {
	gotKeys, gotLocks := GetSchematics(schemata)
	wantKeys := []Schematic{
		{5, 0, 2, 1, 3},
		{4, 3, 4, 0, 2},
		{3, 0, 2, 0, 1},
	}
	wantLocks := []Schematic{
		{0, 5, 3, 4, 3},
		{1, 2, 0, 5, 3},
	}
	if !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Errorf("got %v keys, want %v keys", gotKeys, wantKeys)
	}
	if !reflect.DeepEqual(gotLocks, wantLocks) {
		t.Errorf("got %v locks, want %v locks", gotLocks, wantLocks)
	}
}

func TestGetPairs(t *testing.T) {
	got := GetPairs(schemata)
	want := 3
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
