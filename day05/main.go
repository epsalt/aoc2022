package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type op struct {
	amount, from, to int
}

type stack []rune

func (s stack) pop() (rune, stack) {
	var element rune
	element, s = s[len(s)-1], s[:len(s)-1]
	return element, s
}

func (s stack) lift(n int) (stack, stack) {
	var s2 stack
	s2, s = s[len(s)-n:], s[:len(s)-n]
	return s2, s
}

func part1(lines []string) (ans string) {
	stacks, ops := process(lines)
	for _, op := range ops {
		for i := 0; i < op.amount; i++ {
			var moving rune
			moving, stacks[op.from] = stacks[op.from].pop()
			stacks[op.to] = append(stacks[op.to], moving)
		}
	}
	for _, s := range stacks {
		ans += string(s[len(s)-1])
	}
	return ans
}

func part2(lines []string) (ans string) {
	stacks, ops := process(lines)
	for _, op := range ops {
		var moving stack
		moving, stacks[op.from] = stacks[op.from].lift(op.amount)
		stacks[op.to] = append(stacks[op.to], moving...)
	}
	for _, s := range stacks {
		ans += string(s[len(s)-1])
	}
	return ans
}

func process(lines []string) ([]stack, []op) {
	stacks := make([]stack, 0)
	for _, line := range lines {
		if line == "" {
			break
		}
		for i, s := range line {
			if unicode.IsLetter(s) {
				n := (i - 1) / 4
				if n >= len(stacks) {
					for j := len(stacks); j < n+1; j++ {
						s := make(stack, 0)
						stacks = append(stacks, s)
					}
				}
				stacks[n] = append([]rune{s}, stacks[n]...)
			}
		}
	}
	ops := make([]op, 0)
	for _, line := range lines {
		var amount, start, end int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &start, &end)
		if err == nil {
			ops = append(ops, op{amount, start - 1, end - 1})
		}
	}
	return stacks, ops
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
	fmt.Println("Part1:", part1(lines))
	fmt.Println("Part2:", part2(lines))
}
