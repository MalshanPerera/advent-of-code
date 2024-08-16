package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fileName := "calibrationValues.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {

		str := scanner.Text()
		numbers := []int{}

		for i := 0; i < len(str); i++ {
			if num, err := strconv.Atoi(string(str[i])); err == nil {
				numbers = append(numbers, num)
			}
		}

		if len(numbers) > 0 {
			firstNumber, lastNumber := numbers[0], numbers[len(numbers)-1]

			finalStr := fmt.Sprintf("%d%d", firstNumber, lastNumber)
			if finalNumber, err := strconv.Atoi(finalStr); err == nil {
				sum += finalNumber
			}
		}
	}

	fmt.Println("Sum of the calibration values:", sum)
}
