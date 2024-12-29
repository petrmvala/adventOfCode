package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	correct, incorrect := Order(string(f))
	fmt.Println("correct order sum:", correct)
	fmt.Println("incorrect order sum:", incorrect)
}

func Order(input string) (int, int) {
	rules := map[string]bool{}
	correct := 0
	incorrect := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "|") {
			rules[line] = true
			continue
		}
		numbers := strings.Split(line, ",")
		if !isValid(numbers, rules) {
			nums := fixOrder(numbers, rules)
			n, err := strconv.Atoi(nums[len(nums)/2])
			if err != nil {
				log.Fatal(err)
			}
			incorrect += n
			continue
		}
		n, err := strconv.Atoi(numbers[len(numbers)/2])
		if err != nil {
			log.Fatal(err)
		}
		correct += n
	}
	return correct, incorrect
}

func isValid(numbers []string, rules map[string]bool) bool {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			n := numbers[j] + "|" + numbers[i]
			if rules[n] {
				return false
			}
		}
	}
	return true
}

func fixOrder(numbers []string, rules map[string]bool) []string {
	sort.Slice(numbers, func(i, j int) bool {
		n := numbers[j] + "|" + numbers[i]
		return !rules[n]
	})
	return numbers
}
