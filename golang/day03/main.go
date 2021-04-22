package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func parseClaim(line string) Claim {
	id, x, y, w, h := 0, 0, 0, 0, 0
	// #1 @ 1,3: 4x4
	res := new(Claim)
	fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
	res.id = id
	res.x = x
	res.y = y
	res.width = w
	res.height = h
	return *res
}

func totalWidth(claim Claim) int {
	return 2*claim.x + claim.width
}

func totalHeight(claim Claim) int {
	return 2*claim.y + claim.height
}

func maxClaims(claims []Claim) (int, int) {
	maxH := totalHeight(claims[0])
	maxW := totalWidth(claims[0])
	for _, c := range claims {
		if maxW < totalWidth(c) {
			maxW = totalWidth(c)
		}
		if maxH < totalHeight(c) {
			maxH = totalHeight(c)
		}
	}
	return maxH, maxW
}

func readFile() []Claim {
	file, _ := os.Open("../../inputs/day03.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []Claim{}
	for scanner.Scan() {
		l := scanner.Text()
		res = append(res, parseClaim(l))
	}
	return res
}

func fillMatrix(list []Claim) [][]int {
	maxH, maxW := maxClaims(list)
	matrix := make([][]int, maxW)
	for i := 0; i < maxW; i++ {
		matrix[i] = make([]int, maxH)
	}
	for _, c := range list {
		for i := 0; i < c.width; i++ {
			for j := 0; j < c.height; j++ {
				matrix[c.x+i][c.y+j]++
			}
		}
	}
	return matrix
}

func part1(list []Claim) int {
	matrix := fillMatrix(list)
	maxW := len(matrix)
	maxH := len(matrix[0])
	sum := 0
	for i := 0; i < maxW; i++ {
		for j := 0; j < maxH; j++ {
			if matrix[i][j] > 1 {
				sum++
			}
		}
	}

	return sum
}

func overlap(c Claim, matrix *[][]int) bool {
	for i := 0; i < c.width; i++ {
		for j := 0; j < c.height; j++ {
			if (*matrix)[c.x+i][c.y+j] > 1 {
				return true
			}
		}
	}
	return false
}

func part2(list []Claim) int {
	matrix := fillMatrix(list)
	for _, c := range list {
		if !overlap(c, &matrix) {
			return c.id
		}
	}

	return -1
}

func main() {
	file, _ := os.Open("../../inputs/day01.txt")
	defer file.Close()
	list := readFile()
	start := time.Now()
	fmt.Println("part1: ", part1(list))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(list))
	fmt.Println(time.Since(start))
}
