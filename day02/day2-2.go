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
		for _, match := range matches {
			digit, _ := strconv.Atoi(match)
			digits = append(digits, digit)
		}

		if isSafeReport(digits) {
			safeReportCount++
		} else {
			// Check with Problem Dampener: remove each level and recheck
			for i := 0; i < len(digits); i++ {
				modifiedDigits := append([]int{}, digits[:i]...)
				modifiedDigits = append(modifiedDigits, digits[i+1:]...)
				if isSafeReport(modifiedDigits) {
					safeReportCount++
					break
				}
			}
		}
	}
	fmt.Println("Safe report count with Problem Dampener:", safeReportCount)
}

func isSafeReport(digits []int) bool {
	if len(digits) < 2 {
		return true // A single level is always safe
	}

	increasing := digits[1] >= digits[0]
	for i := 1; i < len(digits); i++ {
		absoluteDiff := calculateDifference3(digits[i-1], digits[i])
		diff := digits[i-1] - digits[i]
		if (absoluteDiff < 1 || absoluteDiff > 3) || (increasing && diff > 0) || (!increasing && diff < 0) {
			return false
		}
	}
	return true
}

func calculateDifference3(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
