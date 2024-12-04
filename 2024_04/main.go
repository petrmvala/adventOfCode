package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(f)
	start := time.Now()
	xmas := XMAS(data)
	xtime := time.Since(start)
	start = time.Now()
	xMask := XMask(data)
	xMtime := time.Since(start)
	fmt.Println("XMAS:", xmas, "elapsed:", xtime)
	fmt.Println("X-MAS:", xMask, "elapsed:", xMtime)
}

func XMAS(data string) int {
	lines := getLines(data)
	sum := 0
	for _, line := range lines {
		idx := 0
		for {
			l := line[idx:]
			xmas := strings.Index(l, "XMAS")
			samx := strings.Index(l, "SAMX")
			first := 0
			if xmas+samx == -2 {
				// nothing found
				break
			}
			if xmas > -1 && (samx == -1 || xmas < samx) {
				first = xmas
			} else {
				first = samx
			}
			sum++
			idx += first + 1
		}
	}
	return sum
}

func XMask(data string) int {
	mat := toMatrix(data)
	sum := 0
	for row := 1; row < len(mat)-1; row++ {
		for col := 1; col < len(mat[row])-1; col++ {
			if mat[row][col] != "A" ||
				(mat[row-1][col-1] != "M" && mat[row-1][col-1] != "S") ||
				(mat[row+1][col+1] != "M" && mat[row+1][col+1] != "S") ||
				(mat[row-1][col+1] != "M" && mat[row-1][col+1] != "S") ||
				(mat[row+1][col-1] != "M" && mat[row+1][col-1] != "S") {
				continue
			}
			if (mat[row-1][col-1] == "M" && mat[row+1][col+1] != "S") ||
				(mat[row-1][col-1] == "S" && mat[row+1][col+1] != "M") ||
				(mat[row-1][col+1] == "M" && mat[row+1][col-1] != "S") ||
				(mat[row-1][col+1] == "S" && mat[row+1][col-1] != "M") {
				continue
			}
			sum++
		}
	}
	return sum
}

func getLines(data string) []string {
	mat := toMatrix(data)
	lines := []string{}
	lines = append(lines, getLinesByRows(mat)...)
	lines = append(lines, getLinesByColumns(mat)...)
	lines = append(lines, getLinesDiagFalling(mat)...)
	lines = append(lines, getLinesDiagRising(mat)...)
	return lines
}

func getLinesByRows(mat [][]string) []string {
	lines := []string{}
	tmp := make([]string, len(mat[0]))
	for row := 0; row < len(mat); row++ {
		copy(tmp, mat[row])
		lines = append(lines, strings.Join(tmp, ""))
	}
	return lines
}

func getLinesByColumns(mat [][]string) []string {
	lines := []string{}
	for col := 0; col < len(mat[0]); col++ {
		c := []string{}
		for row := 0; row < len(mat); row++ {
			c = append(c, mat[row][col])
		}
		lines = append(lines, strings.Join(c, ""))
	}
	return lines
}

func getLinesDiagFalling(mat [][]string) []string {
	lines := []string{}
	startRow := 0
	startCol := len(mat[startRow]) - 1
	for {
		diag := []string{}
		row := startRow
		col := startCol
		for row < len(mat) && col < len(mat[row]) {
			diag = append(diag, mat[row][col])
			row++
			col++
		}
		if startCol > 0 {
			startCol--
		} else if startRow < len(mat)-1 {
			startRow++
		} else {
			break
		}
		if len(diag) < 4 {
			continue
		}
		lines = append(lines, strings.Join(diag, ""))
	}
	return lines
}

func getLinesDiagRising(mat [][]string) []string {
	lines := []string{}
	startRow := 0
	startCol := 0
	for {
		diag := []string{}
		row := startRow
		col := startCol
		for row >= 0 && col < len(mat[row]) {
			diag = append(diag, mat[row][col])
			row--
			col++
		}
		if startRow < len(mat)-1 {
			startRow++
		} else if startCol < len(mat[startRow])-1 {
			startCol++
		} else {
			break
		}
		if len(diag) < 4 {
			continue
		}
		lines = append(lines, strings.Join(diag, ""))
	}
	return lines
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
