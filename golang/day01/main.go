package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile() []int {
	file, _ := os.Open("../../inputs/day01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []int{}
	for scanner.Scan() {
		l := scanner.Text()
		// fmt.Println(l)
		value, _ := strconv.Atoi(l)
		res = append(res, value)
	}
	return res
}

func part1(list []int) int {
	total := 0
	for _, v := range list {
		total += v
	}
	return total
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
	fmt.Println("part2: ", part2(list))
}
