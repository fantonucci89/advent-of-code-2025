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

func fullCirclesDone(value int) int {
	return value / 100
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
		if num > current {
			*counter += fullCirclesDone(current + num)
		}
		return (current - num + 100) % 100
	case 'R':
		if (num + current) >= 100 {
			*counter += fullCirclesDone(num + current)
		}
		return (current + num) % 100
	}
	return current
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
