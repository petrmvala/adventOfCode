package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var fw = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
var bw = regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)

// First star:
// var re = regexp.MustCompile(`([123456789])`)

func main() {
	readFile, err := os.Open("day01/data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	s := bufio.NewScanner(readFile)
	s.Split(bufio.ScanLines)

	sum := 0
	for s.Scan() {
		sum += lineMatch(s.Text())
	}

	fmt.Println("result: ", sum)
}

func lineMatch(line string) int {

	reversed := Reverse(line)

	firstMatch := fw.FindStringIndex(line)
	lastMatch := bw.FindStringIndex(reversed)

	first := spellTostring(line[firstMatch[0]:firstMatch[1]])
	last := spellTostring(reversed[lastMatch[0]:lastMatch[1]])

	// fmt.Printf("line: %s, first: %s, last: %s\n", line, first, last)

	res, err := strconv.Atoi(first + last)
	if err != nil {
		log.Fatalln("failed to convert: ", err)
	}

	return res
}

func spellTostring(spell string) (r string) {
	switch spell {
	// case "0":
	// 	r = "0"
	case "one", "eno", "1":
		r = "1"
	case "two", "owt", "2":
		r = "2"
	case "three", "eerht", "3":
		r = "3"
	case "four", "ruof", "4":
		r = "4"
	case "five", "evif", "5":
		r = "5"
	case "six", "xis", "6":
		r = "6"
	case "seven", "neves", "7":
		r = "7"
	case "eight", "thgie", "8":
		r = "8"
	case "nine", "enin", "9":
		r = "9"
	}
	return
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
