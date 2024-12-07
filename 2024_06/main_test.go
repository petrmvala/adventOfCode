package main

import (
	"reflect"
	"testing"
)

const sample = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

var trace = map[Position]bool{
	{x: 1, y: 4}: true,
	{x: 1, y: 5}: true,
	{x: 1, y: 6}: true,
	{x: 1, y: 7}: true,
	{x: 1, y: 8}: true,
	{x: 2, y: 4}: true,
	{x: 2, y: 8}: true,
	{x: 3, y: 4}: true,
	{x: 3, y: 8}: true,
	{x: 4, y: 2}: true,
	{x: 4, y: 3}: true,
	{x: 4, y: 4}: true,
	{x: 4, y: 5}: true,
	{x: 4, y: 6}: true,
	{x: 4, y: 8}: true,
	{x: 5, y: 2}: true,
	{x: 5, y: 4}: true,
	{x: 5, y: 6}: true,
	{x: 5, y: 8}: true,
	{x: 6, y: 2}: true,
	{x: 6, y: 3}: true,
	{x: 6, y: 4}: true,
	{x: 6, y: 5}: true,
	{x: 6, y: 6}: true,
	{x: 6, y: 7}: true,
	{x: 6, y: 8}: true,
	{x: 7, y: 1}: true,
	{x: 7, y: 2}: true,
	{x: 7, y: 3}: true,
	{x: 7, y: 4}: true,
	{x: 7, y: 5}: true,
	{x: 7, y: 6}: true,
	{x: 7, y: 7}: true,
	{x: 8, y: 1}: true,
	{x: 8, y: 2}: true,
	{x: 8, y: 3}: true,
	{x: 8, y: 4}: true,
	{x: 8, y: 5}: true,
	{x: 8, y: 6}: true,
	{x: 8, y: 7}: true,
	{x: 9, y: 7}: true,
}

func TestFirstStar(t *testing.T) {
	mat := toMatrix(sample)
	got := FirstStar(mat)
	want := 41
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestTraceGuard(t *testing.T) {
	mat := toMatrix(sample)
	got := traceGuard(mat)
	want := trace
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetGuard(t *testing.T) {
	mat := toMatrix(sample)
	curr, next, err := getGuard(mat)
	if err != nil {
		t.Errorf("no error expected but got one: %q", err)
	}
	wantCurr := Position{x: 6, y: 4}
	wantNext := Position{x: 5, y: 4}
	if !reflect.DeepEqual(curr, wantCurr) {
		t.Errorf("got current position at %+v, want %+v", curr, wantCurr)
	}
	if !reflect.DeepEqual(next, wantNext) {
		t.Errorf("got next position at %+v, want %+v", next, wantNext)
	}
}
