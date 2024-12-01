package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	left := []int{}
	right := []int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		r := regexp.MustCompile(`\d+`)
		matches := r.FindAllString(line, -1)
		leftDigit, err := strconv.Atoi(matches[0])
		if err != nil {
			fmt.Println(err)
		}
		left = append(left, leftDigit)
		rightDigit, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Println(err)
		}
		right = append(right, rightDigit)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += diff(left[i], right[i])
	}

	fmt.Println("Sum", sum)

}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
