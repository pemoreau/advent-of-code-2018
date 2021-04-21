package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() []string {
	file, _ := os.Open("../../inputs/day02.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		l := scanner.Text()
		res = append(res, l)
	}
	return res
}

func check2(s string) bool {
	for _, b := range check(s) {
		if b == 2 {
			// fmt.Printf("check2: %#U occurs %d\n", a, b)
			return true
		}
	}
	return false
}
func check3(s string) bool {
	for _, b := range check(s) {
		if b == 3 {
			// fmt.Printf("check3: %#U occurs %d\n", a, b)
			return true
		}
	}
	return false
}

func check(s string) map[rune]int {
	visited := map[rune]int{}
	for _, b := range s {
		visited[b] += 1
	}
	return visited
}

func part1(list []string) int {
	nb2 := 0
	nb3 := 0
	for _, v := range list {
		if check2(v) {
			nb2 += 1
		}
		if check3(v) {
			nb3 += 1
		}
	}
	return nb2 * nb3
}

func part2(list []int) int {
	visited := map[int]bool{0: true}
	index := 0
	total := 0
	for {
		total += list[index]
		index = (index + 1) % len(list)
		if visited[total] {
			return total
		}
		visited[total] = true

	}
}

func main() {
	file, _ := os.Open("../../inputs/day01.txt")
	defer file.Close()
	list := readFile()

	fmt.Println("part1: ", part1(list))
	// fmt.Println("part2: ", part2(list))
}
