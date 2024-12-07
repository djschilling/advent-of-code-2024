package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	safeReportCount := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		r := regexp.MustCompile("[0-9]+")
		matches := r.FindAllString(line, -1)
		digits := []int{}
		increasing := true
		safeReport := true
		for index, match := range matches {
			digit, _ := strconv.Atoi(match)
			digits = append(digits, digit)
			if index == 1 {
				increasing = digit >= digits[0]
			}
			if index >= 1 {
				absoluteDiff := calculateDifference2(digits[index-1], digit)
				diff := digits[index-1] - digit
				if (absoluteDiff < 1 || absoluteDiff > 3) || (increasing && diff > 0) || (!increasing && diff < 0) {
					safeReport = false
				}
			}
		}
		if safeReport {
			safeReportCount++
		}
	}
	fmt.Println("Safe report count", safeReportCount)
}

func calculateDifference2(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
