package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

type Candidate struct {
	value   int
	x_start int
	x_end   int
}

func main() {

	readFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	s := bufio.NewScanner(readFile)
	s.Split(bufio.ScanLines)

	var schema string

	for s.Scan() {
		schema += fmt.Sprintln(s.Text()) // Println will add back the final '\n'
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	start := time.Now()
	sum := sum(schema)
	timeSum := time.Since(start)

	// start = time.Now()
	// gears := gears(schema)
	// timeGears := time.Since(start)
	// fmt.Println("sum: ", sum, " gears: ", gears)
	fmt.Println("sum time elapsed: ", timeSum)
	// fmt.Println("gears time elapsed: ", timeGears)
	if sum != 556367 {
		fmt.Println("the sum is wrong")
	}

}

// func gears(schematic string) int {
// 	return 0
// }

func sum(schematic string) int {
	var (
		valid               bool
		pos_x, sum          int
		buffer              string
		abovePool, thisPool []Candidate
	)
	thisSymbol := make(map[int]bool)
	aboveSymbol := thisSymbol

	pos_x = -1
	for _, s := range schematic {
		pos_x++
		switch {
		case unicode.IsDigit(s):
			buffer += string(s)
			if aboveSymbol[pos_x] {
				valid = true
			}
		case s == '.':
			if aboveSymbol[pos_x] {
				valid = true
			}
			writeValue(&buffer, &thisPool, &sum, pos_x-1, valid)
			valid = aboveSymbol[pos_x]
		case s == '\n':
			writeValue(&buffer, &thisPool, &sum, pos_x-1, valid)
			abovePool, thisPool = thisPool, []Candidate{}
			aboveSymbol, thisSymbol = thisSymbol, map[int]bool{}
			valid = false
			pos_x = -1
		default:
			thisSymbol[pos_x] = true
			valid = true
			writeValue(&buffer, &thisPool, &sum, pos_x-1, valid)
			for _, c := range abovePool {
				if c.x_start-1 > pos_x {
					continue
				}
				if c.x_end+1 < pos_x {
					continue
				}
				sum += c.value
			}
		}
	}

	return sum
}

func writeValue(buffer *string, thisPool *[]Candidate, sum *int, pos_x int, valid bool) {
	buflen := len(*buffer)
	if buflen == 0 {
		return
	}

	value, err := strconv.Atoi(*buffer)
	if err != nil {
		log.Fatal(err)
	}
	*buffer = ""

	if !valid {
		*thisPool = append(*thisPool, Candidate{
			value:   value,
			x_start: pos_x - (buflen - 1),
			x_end:   pos_x,
		})
		return
	}

	*sum += value
}
