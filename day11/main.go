package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items []int
	op    func(w int) int
	test  func(w int) bool
	t     int
	f     int
}

func solve(arr []monkey, rounds int, worry func(w int) int) (ans int) {
	count := make(map[int]int)
	for i := 0; i < rounds; i++ {
		for j, m := range arr {
			for _, item := range m.items {
				count[j] += 1
				item = m.op(item)
				item = worry(item)
				if m.test(item) {
					arr[m.t].items = append(arr[m.t].items, item)
				} else {
					arr[m.f].items = append(arr[m.f].items, item)
				}
			}
			arr[j].items = arr[j].items[:0]
		}
	}
	inspects := make([]int, 0)
	for k, v := range count {
		fmt.Println(k, v)
		inspects = append(inspects, v)
	}
	sort.Ints(inspects)
	return inspects[len(inspects)-1] * inspects[len(inspects)-2]
}

func process(lines []string) (arr []monkey, factor int) {
	monkeys := make([]monkey, 0)
	factor = 1
	var curr monkey
	for _, line := range lines {
		switch {
		case strings.Contains(line, "Starting"):
			re := regexp.MustCompile("[0-9]+")
			for _, match := range re.FindAllString(line, -1) {
				val, _ := strconv.Atoi(match)
				curr.items = append(curr.items, val)
			}
		case strings.Contains(line, "Operation"):
			var op, n string
			fmt.Sscanf(strings.TrimSpace(line), "Operation: new = old %s %s", &op, &n)
			if val, err := strconv.Atoi(n); err == nil {
				if op == "*" {
					curr.op = func(w int) int { return w * val }
				} else {
					curr.op = func(w int) int { return w + val }
				}
			} else {
				if op == "*" {
					curr.op = func(w int) int { return w * w }
				} else {
					curr.op = func(w int) int { return w + w }
				}
			}
		case strings.Contains(line, "Test"):
			var n int
			fmt.Sscanf(strings.TrimSpace(line), "Test: divisible by %d", &n)
			factor *= n
			curr.test = func(w int) bool { return w%n == 0 }
		case strings.Contains(line, "If true"):
			var n int
			fmt.Sscanf(strings.TrimSpace(line), "If true: throw to monkey %d", &n)
			curr.t = n
		case strings.Contains(line, "If false"):
			var n int
			fmt.Sscanf(strings.TrimSpace(line), "If false: throw to monkey %d", &n)
			curr.f = n
		case len(line) == 0:
			monkeys = append(monkeys, curr)
			curr = monkey{}
		}
	}
	monkeys = append(monkeys, curr)
	return monkeys, factor
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	monkeys, factor := process(lines)
	fmt.Println("Part1:", solve(monkeys, 20, func(w int) int { return w / 3 }))

	monkeys, factor = process(lines)
	fmt.Println("Part2:", solve(monkeys, 10000, func(w int) int { return w % factor }))
}
