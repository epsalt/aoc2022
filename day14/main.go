package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type loc struct{ x, y int }

func (a loc) less(b loc) bool {
	if a.x < b.x {
		return true
	} else if a.x == b.x {
		return a.y < b.y
	}
	return false
}

func part1(m map[loc]bool) (ans int) {
	void := 0
	for k := range m {
		if k.y > void {
			void = k.y
		}
	}
	for true {
		curr := loc{500, 0}
		for true {
			if curr.y > void {
				return ans
			}
			if !m[loc{curr.x, curr.y + 1}] {
				curr.y += 1
			} else if !m[loc{curr.x - 1, curr.y + 1}] {
				curr.x -= 1
				curr.y += 1
			} else if !m[loc{curr.x + 1, curr.y + 1}] {
				curr.x += 1
				curr.y += 1
			} else {
				m[curr] = true
				ans += 1
				break
			}
		}
	}
	return ans
}

func part2(m map[loc]bool) (ans int) {
	floor := 0
	for k := range m {
		if k.y > floor {
			floor = k.y
		}
	}
	floor += 1
	for true {
		curr := loc{500, 0}
		for true {
			if !m[loc{curr.x, curr.y + 1}] && curr.y < floor {
				curr.y += 1
			} else if !m[loc{curr.x - 1, curr.y + 1}] && curr.y < floor {
				curr.x -= 1
				curr.y += 1
			} else if !m[loc{curr.x + 1, curr.y + 1}] && curr.y < floor {
				curr.x += 1
				curr.y += 1
			} else {
				m[curr] = true
				ans += 1
				if curr.x == 500 && curr.y == 0 {
					return ans
				}
				break
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

	m1, m2 := make(map[loc]bool), make(map[loc]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		locs := make([]loc, 0)

		for _, s := range strings.Split(line, " -> ") {
			var x, y int
			fmt.Sscanf(s, "%d,%d", &x, &y)
			locs = append(locs, loc{x, y})
		}

		curr := locs[0]
		for _, next := range locs[1:] {
			pair := []loc{curr, next}
			sort.Slice(pair, func(i int, j int) bool { return pair[i].less(pair[j]) })

			if curr.x != next.x {
				for i := pair[0].x; i < pair[1].x+1; i++ {
					m1[loc{i, curr.y}], m2[loc{i, curr.y}] = true, true
				}
			} else {
				for j := pair[0].y; j < pair[1].y+1; j++ {
					m1[loc{curr.x, j}], m2[loc{curr.x, j}] = true, true
				}
			}
			curr = next
		}
	}
	fmt.Println("Part1:", part1(m1))
	fmt.Println("Part2:", part2(m2))
}
