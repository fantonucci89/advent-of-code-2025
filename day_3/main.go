package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/fantonucci89/advent_of_code_2025/internal/utils"
)

func findHigherNumber(line string) (int, error) {
	tens, _ := strconv.Atoi(line[0:1])
	unit, _ := strconv.Atoi(line[1:2])

	for i := 2; i < len(line)-1; i++ {
		charInt, _ := strconv.Atoi(string(line[i]))
		if charInt > tens {
			tens = charInt
			nextCharInt, _ := strconv.Atoi(string(line[i+1]))
			unit = nextCharInt
			fmt.Println(tens, unit)
			continue
		}
		if charInt > unit {
			unit = charInt
			fmt.Println(tens, unit)
			continue
		}
	}

	result, err := fmt.Printf("%d%d", tens, unit)
	if err != nil {
		return 0, err
	}
	fmt.Println("End iteration")
	return result, nil
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
