package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
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

	similarityScore := 0
	for _, leftDigit := range left {
		foundRight := 0
		for _, rightDigit := range right {
			if leftDigit == rightDigit {
				foundRight += 1
			}
		}
		similarityScore += leftDigit * foundRight
	}

	fmt.Println("Similarity Score", similarityScore)

}
