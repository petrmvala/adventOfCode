package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := GetSets(string(f))
	fmt.Println("sets:", len(s))
}

func GetSets(in string) []string {
	sets := map[string]bool{}
	paths := map[string][]string{}
	addPath := func(paths map[string][]string, start, end string) []string {
		if _, ok := paths[start]; !ok {
			return []string{end}
		}
		return append(paths[start], end)
	}
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nodes := strings.Split(line, "-")
		paths[nodes[0]] = addPath(paths, nodes[0], nodes[1])
		paths[nodes[1]] = addPath(paths, nodes[1], nodes[0])
	}
	for first, seconds := range paths {
		for _, second := range seconds {
			thirds, ok := paths[second]
			if !ok {
				continue
			}
			for _, third := range thirds {
				firsts, ok := paths[third]
				if !ok {
					continue
				}
				for _, firstEnd := range firsts {
					if first == firstEnd {
						foundT := false
						ss := []string{first, second, third}
						for _, val := range ss {
							if strings.HasPrefix(val, "t") {
								foundT = true
							}
						}
						if !foundT {
							continue
						}
						sort.Strings(ss)
						s := fmt.Sprintf("%s,%s,%s", ss[0], ss[1], ss[2])
						sets[s] = true
					}
				}
			}
		}
	}
	l := []string{}
	for s := range sets {
		l = append(l, s)
	}
	sort.Strings(l)
	return l
}
