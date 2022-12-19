package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type loc struct{ x, y int }
type check struct{ left, rock, move int }

func solve(moves []int, rocks [][]string, limit1 int, limit2 int) (ans1 int, ans2 int) {
	height := -1
	jet := true
	njet := 0
	state := make(map[loc]bool)
	cycle := make(map[check][]int)

	for r := 0; r < 10000; r++ {
		rock := rocks[r%len(rocks)]
		left := 2
		bottom := height + 3

		falling := make([]loc, 0)
		for i := range rock {
			for j := range rock[0] {
				if rock[i][j] == '#' {
					falling = append(falling, loc{j + left, len(rock) - i + bottom})
				}
			}
		}
		for true {
			if jet {
				collision := false
				next := make([]loc, 0)

				dx := moves[njet%len(moves)]

				for _, l := range falling {
					new := loc{l.x + dx, l.y}
					if state[new] == true || new.x < 0 || new.x > 6 {
						collision = true
						break
					} else {
						next = append(next, new)
					}
				}
				jet = false
				njet += 1
				if !collision {
					falling = next
				}
			} else {
				landed := false
				next := make([]loc, 0)
				dy := -1
				for _, l := range falling {
					new := loc{l.x, l.y + dy}
					if state[new] == true || new.y == -1 {
						landed = true
						break
					} else {
						next = append(next, new)
					}
				}
				jet = true
				if landed {
					left := 7
					for _, l := range falling {
						state[l] = true
						if l.y > height {
							height = l.y
						}
						if l.x < left {
							left = l.x
						}
					}
					if r == limit1 {
						ans1 = height + 1
					}

					c := check{left, r % len(rocks), njet % len(moves)}
					if v, ok := cycle[c]; ok {
						if (limit2-r)%(r-v[0]) == 0 {
							ans2 = height + (limit2-r)/(r-v[0])*(height-v[1])
							return
						}
					}
					cycle[c] = []int{r, height}
					break
				} else {
					falling = next
				}

			}
		}
	}
	return ans1, ans2
}

func main() {
	s, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	moves := make([]int, len(s)-1)
	for i, r := range strings.TrimSpace(string(s)) {
		if r == '<' {
			moves[i] = -1
		} else {
			moves[i] = 1
		}
	}
	rocks := [][]string{{"####"}, {".#.", "###", ".#."}, {"..#", "..#", "###"}, {"#", "#", "#", "#"}, {"##", "##"}}

	p1, p2 := solve(moves, rocks, 2022, 1000000000000)
	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
}
