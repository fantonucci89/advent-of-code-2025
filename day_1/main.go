package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const StartingPoint int = 50

func ReadData() (*os.File, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

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
	var counter int = 0
	currentValue := StartingPoint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentValue = RotateDial(currentValue, scanner.Text(), &counter)
	}
	return counter
}

func main() {
	file, err := ReadData()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	result := DecryptPassword(file)
	fmt.Print(result)
}
