package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type interval struct {
	start int
	end   int
}

func part1(pairs [][]interval) (ans int) {
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if (a.start <= b.start) && (a.end >= b.end) ||
			(b.start <= a.start) && (b.end >= a.end) {
			ans += 1
		}
	}
	return ans
}

func part2(pairs [][]interval) (ans int) {
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if (a.start <= b.start) && (a.end >= b.start) ||
			(b.start <= a.start) && (b.end >= a.start) {
			ans += 1
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

	pairs := make([][]interval, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a, b, c, d int
		line := scanner.Text()
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d)
		pairs = append(pairs, []interval{interval{a, b}, interval{c, d}})
	}

	fmt.Println("Part1:", part1(pairs))
	fmt.Println("Part2:", part2(pairs))
}
