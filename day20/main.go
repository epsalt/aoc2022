package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

type encryptedFile []*int

func mod(a, b int) int {
	return (a%b + b) % b
}

func (f encryptedFile) move(i, j int) {
	if j == 0 {
		return
	}

	k := mod(i+j, len(f)-1)
	n := f[i]
	f = slices.Delete(f, i, i+1)
	f = slices.Insert(f, k, n)
}

func solve(nums encryptedFile, key int, rounds int) (ans int) {
	for i := range nums {
		p := nums[i]
		*p *= key
	}

	new := make(encryptedFile, len(nums))
	copy(new, nums)

	for round := 0; round < rounds; round++ {
		for _, n := range nums {
			for j, nn := range new {
				if n == nn {
					new.move(j, *nn)
					break
				}
			}
		}
	}
	for i, n := range new {
		if *n == 0 {
			for _, x := range []int{1000, 2000, 3000} {
				ans += *new[(i+x)%len(new)]
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

	scanner := bufio.NewScanner(file)
	nums := make(encryptedFile, 0)

	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		nums = append(nums, &val)
	}
	fmt.Println("Part1:", solve(nums, 1, 1))
	fmt.Println("Part2:", solve(nums, 811589153, 10))
}
