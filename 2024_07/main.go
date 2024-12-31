package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	d := Data(string(f))
	s := Sum(d)
	fmt.Println("sum:", s)
}

func Data(input string) map[int][]int {
	data := map[int][]int{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		l := strings.Split(line, ":")
		total, err := strconv.Atoi(l[0])
		if err != nil {
			log.Fatal(err)
		}
		l = strings.Fields(l[1])
		elem := []int{}
		for _, e := range l {
			ee, err := strconv.Atoi(e)
			if err != nil {
				log.Fatal(err)
			}
			elem = append(elem, ee)
		}
		data[total] = elem
	}
	return data
}

func Sum(input map[int][]int) int {
	sum := 0
	for t, elems := range input {
		// fmt.Println("examining", elems, "=", t)
		sumAll := 0
		mulAll := 1
		for _, i := range elems {
			sumAll += i
			mulAll *= i
		}
		if sumAll > t || mulAll < t {
			// fmt.Println("out of bounds")
			continue
		}
		if sumAll == t || mulAll == t {
			// fmt.Println("short eval")
			sum += t
			continue
		}
		binLen := strings.Repeat("1", len(elems)-1)    // "1111"
		maxSign, err := strconv.ParseInt(binLen, 2, 0) // 15
		if err != nil {
			log.Fatal(err)
		}
		for i := int64(0); i <= maxSign; i++ {
			// 0 -> "0000", 1 -> "0001", ...
			ops := strings.Split(fmt.Sprintf("%0*b", len(binLen), i), "") // ["0", "0", "0", "1"]
			// fmt.Println("ops:", ops)
			currSum := elems[0]
			for idx := 1; idx < len(elems); idx++ {
				if ops[idx-1] == "1" { // 1 == multiply
					// fmt.Println("running", currSum, "*=", elems[idx])
					currSum *= elems[idx]
				} else {
					// fmt.Println("running", currSum, "+=", elems[idx])
					currSum += elems[idx]
				}
			}
			// fmt.Println("currSum =", currSum)
			if currSum == t {
				// fmt.Println("found number")
				// fmt.Println()
				sum += t
				break
			}
		}

	}
	return sum
}
