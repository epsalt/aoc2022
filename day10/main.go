package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	NOOP = iota
	ADDX
)

type op struct {
	name   int
	amount int
}

func run(ops []op) []int {
	x := 1
	cycle := 0
	adds := make([]int, 0)
	xs := make([]int, 0)

	for len(ops) > 0 || len(adds) > 0 {
		cycle += 1

		if len(adds) > 0 {
			x += adds[0]
			adds = adds[1:]
		}

		if len(adds) == 0 && len(ops) > 0 {
			if ops[0].name == ADDX {
				adds = []int{0, ops[0].amount}
			}
			ops = ops[1:]
		}
		xs = append(xs, x)
	}
	return xs
}

func part1(ops []op) (ans int) {
	for i, x := range run(ops) {
		cycle := i + 1
		if (cycle-20)%40 == 0 {
			ans += cycle * x
		}
	}
	return ans
}

func part2(ops []op) (ans []string) {
	var row string
	for i, x := range run(ops) {
		cycle := i + 1
		position := (cycle - 1) % 40
		if x <= position+1 && x >= position-1 {
			row += "#"
		} else {
			row += "."
		}
		if position == 39 {
			fmt.Println(row)
			row = ""
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
	ops := make([]op, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "noop" {
			ops = append(ops, op{NOOP, 0})
		} else {
			var amount int
			fmt.Sscanf(line, "addx %d", &amount)
			ops = append(ops, op{ADDX, amount})
		}
	}
	fmt.Println("Part1:", part1(ops))
	fmt.Println("Part2:")
	part2(ops)
}
