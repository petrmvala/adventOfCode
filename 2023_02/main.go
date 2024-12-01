package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type requires struct {
	red   int
	green int
	blue  int
}

func main() {
	sumGames()
}

func sumGames() {
	readFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	s := bufio.NewScanner(readFile)
	s.Split(bufio.ScanLines)

	sum := 0
	require := 0
	for s.Scan() {
		sp := strings.Split(s.Text(), ":")
		gameId, err := strconv.Atoi(strings.Split(sp[0], " ")[1])
		if err != nil {
			log.Fatalln(err)
		}
		if gameValid(sp[1]) {
			sum += gameId
		}
		r := gameRequires(sp[1])
		require += r.blue * r.green * r.red
	}

	fmt.Println("validSum: ", sum)
	fmt.Println("requireSum: ", require)
}

func gameRequires(line string) requires {
	requires := requires{red: 0, green: 0, blue: 0}
	sets := strings.Split(line, ";")
	ss := []string{}
	for _, set := range sets {
		ss = append(ss, strings.Split(set, ",")...)
	}
	for _, cubes := range ss {
		s := strings.Split(strings.TrimSpace(cubes), " ")
		if len(s) != 2 {
			log.Fatalln("cubes split into invalid length: ", line, " split: ", s)
		}
		qty, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		color := s[1]
		if color == "red" && qty > requires.red {
			requires.red = qty
		}
		if color == "green" && qty > requires.green {
			requires.green = qty
		}
		if color == "blue" && qty > requires.blue {
			requires.blue = qty
		}
	}
	return requires
}

func gameValid(line string) bool {
	sets := strings.Split(line, ";")
	ss := []string{}
	for _, set := range sets {
		ss = append(ss, strings.Split(set, ",")...)
	}
	for _, cubes := range ss {
		s := strings.Split(strings.TrimSpace(cubes), " ")
		if len(s) != 2 {
			log.Fatalln("cubes split into invalid length: ", line, " split: ", s)
		}
		qty, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		color := s[1]
		if qty > limitOf(color) {
			return false
		}
	}
	return true
}

func limitOf(color string) (limit int) {
	switch color {
	case "red":
		limit = 12
	case "green":
		limit = 13
	case "blue":
		limit = 14
	default:
		log.Fatalf("unknown color: %s", color)
	}
	return
}
