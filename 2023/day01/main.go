package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {

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

func partTwo() {

	fileName := "calibrationValues.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	validWords := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {

		str := scanner.Text()

		numbers := findNumberWords(str, validWords)

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

func findNumberWords(text string, validWords map[string]int) []int {
	var numbers []int
	currentWord := ""

	for i := 0; i < len(text); i++ {
		char := string(text[i])
		if num, err := strconv.Atoi(char); err == nil {
			numbers = append(numbers, num)
			currentWord = char

			continue
		}

		currentWord += char

		for key, value := range validWords {
			if strings.Contains(currentWord, key) {
				numbers = append(numbers, value)
				currentWord = char
				break
			}
		}
	}

	return numbers
}
