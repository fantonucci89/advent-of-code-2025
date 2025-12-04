package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	TaskStr string
	Task    int = 1
)

func ReadData() (*os.File, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ListRanges(listOfRanges *[]string) {
	file, err := ReadData()
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Add each line to the list of ranges
		*listOfRanges = strings.Split(line, ",")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func CalculateInvalidIDInRange(r *int, idRange string) {
	// Split the range into start and end
	parts := strings.Split(idRange, "-")
	if len(parts) != 2 {
		fmt.Println("Invalid range format:", idRange)
		return
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return
	}

	// Check for invalid IDs in the range
	for id := start; id <= end; id++ {
		if Task == 1 && hasRepeatedDigits(id) {
			*r += id
		}
		if Task == 2 && hasRepeatedDigitsImproved(id) {
			*r += id
		}
	}
}

func hasRepeatedDigits(id int) bool {
	strID := strconv.Itoa(id)
	if len(strID)%2 != 0 {
		return false
	}
	mid := len(strID) / 2
	return strID[:mid] == strID[mid:]
}

func hasRepeatedDigitsImproved(id int) bool {
	strID := strconv.Itoa(id)

	// Try all possible pattern lengths from 1 to half the string length
	for patternLen := 1; patternLen <= len(strID)/2; patternLen++ {
		// Check if the string length is divisible by the pattern length
		if len(strID)%patternLen != 0 {
			continue
		}

		// Extract the pattern
		pattern := strID[:patternLen]

		// Check if the entire string is made up of this pattern repeated
		matched := true
		for i := patternLen; i < len(strID); i += patternLen {
			if strID[i:i+patternLen] != pattern {
				matched = false
				break
			}
		}

		// If we found a repeating pattern, this ID is invalid
		if matched {
			return true
		}
	}
	return false
}

func main() {
	if TaskStr != "" {
		if t, err := strconv.Atoi(TaskStr); err == nil {
			Task = t
		}
	}
	var listOfRanges []string

	ListRanges(&listOfRanges)
	var result int
	for _, r := range listOfRanges {
		CalculateInvalidIDInRange(&result, r)
	}
	fmt.Println("Sum of invalid IDs:", result)
}
