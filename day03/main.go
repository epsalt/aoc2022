package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func score(r rune) int {
	switch {
	case 97 <= r && r <= 122:
		return int(r - 96)
	case 65 <= r && r <= 90:
		return int(r - 38)
	}
	panic("Bad Rune")
}

func part1(sacks []string) (ans int) {
	for _, sack := range sacks {
		compartments := [2]string{sack[:len(sack)/2], sack[len(sack)/2:]}

		m := make(map[rune]bool)
		for _, item := range compartments[0] {
			m[item] = true
		}

		found := make(map[rune]bool)
		for _, item := range compartments[1] {
			if m[item] == true {
				found[item] = true
			}
		}

		for k := range found {
			ans += score(k)
		}

	}
	return ans
}

func part2(sacks []string) (ans int) {
	for i := 0; i < len(sacks); i = i + 3 {
		found := make(map[rune]int)
		group := [3]string{sacks[i], sacks[i+1], sacks[i+2]}

		for _, sack := range group {
			added := make(map[rune]bool)
			for _, item := range sack {
				if !added[item] {
					found[item] += 1
					added[item] = true
				}
			}
		}

		for k, v := range found {
			if v == 3 {
				ans += score(k)
			}
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

	sacks := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sacks = append(sacks, line)
	}

	fmt.Println("Part1:", part1(sacks))
	fmt.Println("Part2:", part2(sacks))
}
