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

	fileName := "calibrationValues.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum = 0
	for scanner.Scan() {

		str := scanner.Text()
		splitText := strings.Split(str, "")
		numbers := []int{}

		for i := 0; i < len(splitText); i++ {
			if num, err := strconv.Atoi(splitText[i]); err == nil {
				numbers = append(numbers, num)
			}
		}

		firstNumber, lastNumber := numbers[0], numbers[len(numbers)-1]

		if len(numbers) == 1 {
			lastNumber = firstNumber
		}

		if finalNumber, err := strconv.Atoi(fmt.Sprintf("%d%d", firstNumber, lastNumber)); err == nil {
			sum += finalNumber
		}
	}

	fmt.Println("Sum of the calibration values:", sum)
}
