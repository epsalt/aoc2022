package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type loc struct{ x, y int }
type interval struct{ open, close int }

func (a interval) less(b interval) bool {
	if a.open < b.open {
		return true
	} else if a.open == b.open {
		return a.close < b.close
	}
	return false
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func dist(a loc, b loc) int {
	return int(math.Abs(float64(a.x)-float64(b.x)) +
		math.Abs(float64(a.y)-float64(b.y)))
}

func part1(pairs [][]loc, r int) (ans int) {
	intervals := make([]interval, 0)
	beaconsInRow := make(map[loc]bool)

	for _, pair := range pairs {
		sensor, beacon := pair[0], pair[1]
		d := dist(sensor, beacon)
		dy := dist(sensor, loc{sensor.x, r})

		if beacon.y == r {
			beaconsInRow[beacon] = true
		}

		if dy <= d {
			interval := interval{sensor.x - (d - dy), sensor.x + (d - dy)}
			intervals = append(intervals, interval)
		}
	}

	sort.Slice(intervals, func(i int, j int) bool { return intervals[i].less(intervals[j]) })
	merged := make([]interval, 0)
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1].close < interval.open {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1].close = max(merged[len(merged)-1].close, interval.close)
		}
	}
	for _, interval := range merged {
		ans += interval.close - interval.open + 1
	}
	return ans - len(beaconsInRow)
}

func part2(pairs [][]loc, limit int) (ans int) {
	for r := 0; r < limit+1; r++ {
		intervals := make([]interval, 0)
		for _, pair := range pairs {
			sensor, beacon := pair[0], pair[1]
			d := dist(sensor, beacon)
			dy := dist(sensor, loc{sensor.x, r})

			if dy <= d {
				interval := interval{max(sensor.x-(d-dy), 0), min(sensor.x+(d-dy), limit)}
				intervals = append(intervals, interval)
			}
		}

		sort.Slice(intervals, func(i int, j int) bool { return intervals[i].less(intervals[j]) })
		merged := make([]interval, 0)
		for _, interval := range intervals {
			if len(merged) == 0 || (merged[len(merged)-1].close+1) < interval.open {
				merged = append(merged, interval)
			} else {
				merged[len(merged)-1].close = max(merged[len(merged)-1].close, interval.close)
			}
		}
		if len(merged) > 1 {
			return (merged[0].close+1)*limit + r
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pairs := make([][]loc, 0)

	for scanner.Scan() {
		var a, b, c, d int
		line := scanner.Text()
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &a, &b, &c, &d)
		pairs = append(pairs, []loc{{a, b}, {c, d}})
	}
	fmt.Println("Part1:", part1(pairs, 2000000))
	fmt.Println("Part2:", part2(pairs, 4000000))
}
