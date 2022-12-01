package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func part1(elves [][]int) int {
	var ans, total int

	for _, calories := range elves {
		total = 0
		for _, n := range calories {
			total += n
		}
		if total > ans {
			ans = total
		}
	}
	return ans
}

func part2(elves [][]int) int {
	var ans, total int
	totals := make([]int, len(elves))

	for _, calories := range elves {
		total = 0
		for _, n := range calories {
			total += n
		}
		totals = append(totals, total)
	}

	sort.Ints(totals)
	leaders := totals[len(totals)-3:]
	for _, leader := range leaders {
		ans += leader
	}
	return ans
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elves := make([][]int, 0)
	elf := make([]int, 0)

	var line string
	var n int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if line != "" {
			n, _ = strconv.Atoi(line)
			elf = append(elf, n)
		} else {
			elves = append(elves, elf)
			elf = make([]int, 0)
		}

	}
	elves = append(elves, elf)
	fmt.Println("Part1:", part1(elves))
	fmt.Println("Part2:", part2(elves))
}
