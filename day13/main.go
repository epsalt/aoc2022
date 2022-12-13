package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type Packets []any

func (a Packets) Len() int           { return len(a) }
func (a Packets) Less(i, j int) bool { return compare(a[i], a[j]) == Less }
func (a Packets) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type comp int

const (
	Less = -1
	Same = 0
	More = 1
)

func compare(a any, b any) comp {
	switch a.(type) {
	case float64:
		switch b.(type) {
		case float64:
			if a.(float64) < b.(float64) {
				return Less
			} else if a.(float64) == b.(float64) {
				return Same
			} else {
				return More
			}
		case []any:
			return compare([]any{a.(float64)}, b)
		}
	case []any:
		switch b.(type) {
		case float64:
			return compare(a, []any{b.(float64)})
		case []any:
			var l int
			if len(a.([]any)) > len(b.([]any)) {
				l = len(a.([]any))
			} else {
				l = len(b.([]any))
			}
			for i := 0; i < l; i++ {
				if i == len(a.([]any)) || i == len(b.([]any)) {
					if len(a.([]any)) == len(b.([]any)) {
						return Same
					} else if len(a.([]any)) < len(b.([]any)) {
						return Less
					} else {
						return More
					}
				}
				if compare(a.([]any)[i], b.([]any)[i]) == More {
					return More
				} else if compare(a.([]any)[i], b.([]any)[i]) == Less {
					return Less
				}
			}
			return Same
		}
	}
	panic("bad type")
}

func part1(pairs []Packets) (ans int) {
	for i, p := range pairs {
		if compare(p[0], p[1]) != More {
			ans += i + 1
		}
	}
	return ans
}

func part2(p Packets) (ans int) {
	ans = 1
	dividers := make([][]any, 0)
	for _, d := range []string{"[[2]]", "[[6]]"} {
		var divider []any
		json.Unmarshal([]byte(d), &divider)
		dividers = append(dividers, divider)
		p = append(p, divider)
	}
	sort.Sort(p)
	for i, x := range p {
		if compare(x, dividers[0]) == Same || compare(x, dividers[1]) == Same {
			ans *= i + 1
		}
	}
	return ans
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pairs := make([]Packets, 0)
	packets := make([]any, 0)
	var p []any

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			var result []any
			json.Unmarshal([]byte(line), &result)
			p = append(p, result)
			packets = append(packets, result)
		}

		if len(p) == 2 {
			pairs = append(pairs, p)
			p = make([]any, 0)
		}
	}
	fmt.Println("Part1:", part1(pairs))
	fmt.Println("Part2:", part2(packets))
}
