package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fantonucci89/advent_of_code_2025/internal/utils"
)

const StartingPoint int = 50

func RotateDial(current int, move string, counter *int) int {
	// Read each line of input file
	if len(move) < 2 {
		return current
	}
	dir := move[0]
	num, err := strconv.Atoi(move[1:])
	if err != nil {
		return current
	}
	switch dir {
	case 'L':
		if current == 0 {
			*counter -= 1
		}
		current = current - num
		if current <= 0 {
			*counter += ((-current) / 100) + 1
		}
	case 'R':
		current = current + num
		*counter += current / 100
	}
	return current % 100
}

func DecryptPassword(file *os.File) int {
	counter := 0
	currentValue := StartingPoint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentValue = RotateDial(currentValue, scanner.Text(), &counter)
	}
	return counter
}

func main() {
	file, err := utils.ReadData("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	result := DecryptPassword(file)
	fmt.Print(result)
}
