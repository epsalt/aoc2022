package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type resources struct {
	ore,
	clay,
	obsidian,
	geode int
}

type blueprint struct {
	oreCost,
	clayCost,
	obsidianCost,
	geodeCost resources
}

type state struct {
	bag,
	robots resources
	t int
}

func solve(bps []blueprint, limit int) (counts []int) {
	robots := resources{1, 0, 0, 0}
	bag := resources{0, 0, 0, 0}
	var dfs func(bp blueprint, robots resources, bag resources, t int) int

	for _, bp := range bps {
		cache := make(map[state]int)
		max := 0

		dfs = func(bp blueprint, robots resources, bag resources, t int) (ans int) {
			if t == 0 {
				if bag.geode > max {
					max = bag.geode
				}
				return bag.geode
			}

			opts := make([]int, 0)
			s := state{bag, robots, t}
			remaining := (robots.geode+t)*(robots.geode+t)/2 - ((robots.geode - 1) * robots.geode / 2)

			if bag.geode+remaining <= max {
				return remaining
			}

			if val, ok := cache[s]; ok {
				return val
			}

			var maxes resources
			for _, r := range []resources{bp.oreCost, bp.clayCost, bp.obsidianCost, bp.geodeCost} {
				if r.ore > maxes.ore {
					maxes.ore = r.ore
				}
				if r.clay > maxes.clay {
					maxes.clay = r.clay
				}
				if r.obsidian > maxes.obsidian {
					maxes.obsidian = r.obsidian
				}
			}

			g := resources{bag.ore + robots.ore, bag.clay + robots.clay, bag.obsidian + robots.obsidian, bag.geode + robots.geode}

			if bag.ore >= bp.oreCost.ore && robots.ore < maxes.ore {
				newRobots := resources{robots.ore + 1, robots.clay, robots.obsidian, robots.geode}
				newBag := resources{g.ore - bp.oreCost.ore, g.clay, g.obsidian, g.geode}
				opts = append(opts, dfs(bp, newRobots, newBag, t-1))
			}
			if bag.ore >= bp.clayCost.ore && robots.clay < maxes.clay {
				newRobots := resources{robots.ore, robots.clay + 1, robots.obsidian, robots.geode}
				newBag := resources{g.ore - bp.clayCost.ore, g.clay, g.obsidian, g.geode}
				opts = append(opts, dfs(bp, newRobots, newBag, t-1))
			}
			if bag.ore >= bp.obsidianCost.ore && bag.clay >= bp.obsidianCost.clay && robots.obsidian < maxes.obsidian {
				newRobots := resources{robots.ore, robots.clay, robots.obsidian + 1, robots.geode}
				newBag := resources{g.ore - bp.obsidianCost.ore, g.clay - bp.obsidianCost.clay, g.obsidian, g.geode}
				opts = append(opts, dfs(bp, newRobots, newBag, t-1))
			}
			if bag.ore >= bp.geodeCost.ore && bag.obsidian >= bp.geodeCost.obsidian {
				newRobots := resources{robots.ore, robots.clay, robots.obsidian, robots.geode + 1}
				newBag := resources{g.ore - bp.geodeCost.ore, g.clay, g.obsidian - bp.geodeCost.obsidian, g.geode}

				opts = append(opts, dfs(bp, newRobots, newBag, t-1))
			}

			opts = append(opts, dfs(bp, robots, g, t-1))

			for _, opt := range opts {
				if opt > ans {
					ans = opt
				}
			}

			cache[s] = ans
			return ans
		}
		counts = append(counts, dfs(bp, robots, bag, limit))
	}
	return counts
}

func part1(bps []blueprint, limit int) (ans int) {
	for i, c := range solve(bps, limit) {
		ans += (i + 1) * c
	}
	return ans
}

func part2(bps []blueprint, limit int) (ans int) {
	ans = 1
	for _, c := range solve(bps, limit) {
		ans *= c
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
	blueprints := make([]blueprint, 0)

	for scanner.Scan() {
		line := scanner.Text()
		var n, a, b, c, d, e, f int

		fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &n, &a, &b, &c, &d, &e, &f)
		bp := blueprint{
			oreCost:      resources{a, 0, 0, 0},
			clayCost:     resources{b, 0, 0, 0},
			obsidianCost: resources{c, d, 0, 0},
			geodeCost:    resources{e, 0, f, 0},
		}

		blueprints = append(blueprints, bp)
	}
	fmt.Println("Part1:", part1(blueprints, 24))
	fmt.Println("Part2:", part2(blueprints[:3], 32))
}
