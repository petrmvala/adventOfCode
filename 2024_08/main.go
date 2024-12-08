package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Position struct {
	x, y int
}

type Antennas map[string][]Position

type Antinodes map[Position]bool

func main() {
	start := time.Now()
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	mat := toMatrix(string(f))
	first := FirstStar(mat)
	elapsed := time.Since(start)
	fmt.Println("first:", first, "elapsed:", elapsed)
}

func FirstStar(data [][]string) int {
	ant := getAntennas(data)
	nod := getAntinodes(ant, len(data), len(data[0]))
	return len(nod)
}

func getAntinodes(antennas Antennas, lenX, lenY int) Antinodes {
	n := Antinodes{}
	for _, freq := range antennas {
		for ant := 0; ant < len(freq)-1; ant++ {
			for next := ant + 1; next < len(freq); next++ {
				vec := Position{x: freq[next].x - freq[ant].x, y: freq[next].y - freq[ant].y}
				p := Position{
					x: freq[ant].x + 2*vec.x,
					y: freq[ant].y + 2*vec.y,
				}
				if p.x >= 0 && p.x < lenX && p.y >= 0 && p.y < lenY {
					n[p] = true
				}
				p = Position{
					x: freq[ant].x - vec.x,
					y: freq[ant].y - vec.y,
				}
				if p.x >= 0 && p.x < lenX && p.y >= 0 && p.y < lenY {
					n[p] = true
				}
			}
		}
	}
	return n
}

func getAntennas(data [][]string) Antennas {
	a := Antennas{}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == "." {
				continue
			}
			if _, ok := a[data[i][j]]; !ok {
				a[data[i][j]] = []Position{}
			}
			a[data[i][j]] = append(a[data[i][j]], Position{x: i, y: j})
		}
	}
	return a
}

func toMatrix(data string) [][]string {
	mat := make([][]string, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		chars := strings.Split(line, "")
		if len(line) == 0 {
			continue
		}
		mat = append(mat, chars)
	}
	return mat
}
