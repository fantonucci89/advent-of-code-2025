package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		if hasRepeatedDigitsImproved(id) {
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
	var listDividers []int
	// Find minimum common multiple
	for i := 2; i <= (len(strId) / 2); i++ {
		if len(strId)%i == 0 {
			listDividers = append(listDividers, i)
		}
	}

	if len(listDividers) == 0 {
		listDividers = append(listDividers, 1)
	}

	for _, divider := range listDividers {
		// Split the string into parts of length divider
		substr := strId[:divider]
		matched := true
		for j := divider; j < len(strId); j += divider {
			if strId[j:j+divider] != substr {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}
	return false
}

func main() {
	var listOfRanges []string

	ListRanges(&listOfRanges)
	var result int
	for _, r := range listOfRanges {
		CalculateInvalidIdInRange(&result, r)
	}
	fmt.Println("Sum of invalid IDs:", result)
}
