package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Position struct {
	x, y int
}

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	mat := toMatrix(string(f))
	first := FirstStar(mat)
	tFirst := time.Since(start)
	fmt.Println("firstStar:", first, "elapsed:", tFirst)
}

func FirstStar(data [][]string) int {
	trace := traceGuard(data)
	return len(trace)
}

func traceGuard(data [][]string) map[Position]bool {
	current, next, err := getGuard(data)
	if err != nil {
		log.Fatal(err)
	}
	trace := make(map[Position]bool)
	trace[current] = true
	for {
		if !positionInMap(data, next) {
			break
		}
		if positionIsObstacle(data, next) { // turn 90 degrees right
			next = turnRight(current, next)
		}
		if positionIsObstacle(data, next) { // turn 180 degrees
			next = turnRight(current, next)
		}
		direction := Position{
			x: next.x - current.x,
			y: next.y - current.y,
		}
		current = next
		trace[current] = true
		next = Position{
			x: next.x + direction.x,
			y: next.y + direction.y,
		}
	}
	return trace
}

func turnRight(current, next Position) Position {
	direction := Position{
		x: next.x - current.x,
		y: next.y - current.y,
	}
	var c Position
	switch {
	case direction.x == -1 && direction.y == 0: // ^
		c = Position{x: current.x, y: current.y + 1} // >
	case direction.x == 0 && direction.y == 1: // >
		c = Position{x: current.x + 1, y: current.y} // v
	case direction.x == +1 && direction.y == 0: // v
		c = Position{x: current.x, y: current.y - 1} // <
	case direction.x == 0 && direction.y == -1: // <
		c = Position{x: current.x - 1, y: current.y} // ^
	}
	return c
}

func positionIsObstacle(data [][]string, p Position) bool {
	if data[p.x][p.y] == "#" {
		return true
	}
	return false
}

func positionInMap(data [][]string, p Position) bool {
	if p.x < 0 || p.x >= len(data) || p.y < 0 || p.y >= len(data[p.x]) {
		return false
	}
	return true
}

func getGuard(data [][]string) (current, next Position, err error) {
	for r := 0; r < len(data); r++ {
		for c := 0; c < len(data[r]); c++ {
			p := Position{x: r, y: c}
			switch data[r][c] {
			case "^":
				return p, Position{x: r - 1, y: c}, nil
			case ">":
				return p, Position{x: r, y: c + 1}, nil
			case "v":
				return p, Position{x: r + 1, y: c}, nil
			case "<":
				return p, Position{x: r, y: c - 1}, nil
			default:
				continue
			}
		}
	}
	return Position{}, Position{}, errors.New("initial position not found")
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
