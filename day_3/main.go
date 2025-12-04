package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/fantonucci89/advent_of_code_2025/internal/utils"
)

func findHigherNumber(bank string) (int, error) {
	maxJoltage := 0

	// Try all possible combinations of two digits
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			// Form a 2-digit number using digits at positions i and j
			twoDigitStr := string(bank[i]) + string(bank[j])
			twoDigitNum, err := strconv.Atoi(twoDigitStr)
			if err != nil {
				return 0, err
			}

			// Keep track of the maximum
			if twoDigitNum > maxJoltage {
				maxJoltage = twoDigitNum
			}
		}
	}

	return maxJoltage, nil
}

func main() {
	file, err := utils.ReadData("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res, err := findHigherNumber(scanner.Text())
		if err != nil {
			panic(err)
		}
		counter += res
	}
	fmt.Println(counter)
}
