package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var task int = 1

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

func CalculateInvalidIdInRange(r *int, idRange string) {
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
		if task == 1 && hasRepeatedDigits(id) {
			*r += id
		}
		if task == 2 && hasRepeatedDigitsImproved(id) {
			*r += id
		}
	}
}

func hasRepeatedDigits(id int) bool {
	strId := strconv.Itoa(id)
	if len(strId)%2 != 0 {
		return false
	}
	mid := len(strId) / 2
	return strId[:mid] == strId[mid:]
}

func hasRepeatedDigitsImproved(id int) bool {
	strId := strconv.Itoa(id)

	// Try all possible pattern lengths from 1 to half the string length
	for patternLen := 1; patternLen <= len(strId)/2; patternLen++ {
		// Check if the string length is divisible by the pattern length
		if len(strId)%patternLen != 0 {
			continue
		}

		// Extract the pattern
		pattern := strId[:patternLen]

		// Check if the entire string is made up of this pattern repeated
		matched := true
		for i := patternLen; i < len(strId); i += patternLen {
			if strId[i:i+patternLen] != pattern {
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
	useTask := flag.Int("task", 1, "Task number (1 or 2)")
	flag.Parse()

	task = *useTask
	var listOfRanges []string

	ListRanges(&listOfRanges)
	var result int
	for _, r := range listOfRanges {
		CalculateInvalidIdInRange(&result, r)
	}
	fmt.Println("Sum of invalid IDs:", result)
}
