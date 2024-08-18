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
	fileName := "gameData.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxCube := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var sum int
	for scanner.Scan() {

		str := scanner.Text()

		strList := strings.Split(str, ":")
		gameID := strList[0]
		subList := strings.Split(strList[1], ";")

		isValid := true
		for _, value := range subList {

			cubes := strings.Split(value, ",")

			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)

				strList := strings.Split(cube, " ")
				cubeColor := strList[1]
				availableCubes := strList[0]

				if max, exists := maxCube[cubeColor]; exists {
					currentCount, err := strconv.Atoi(availableCubes)
					if err != nil {
						panic(err)
					}

					if max < currentCount {
						isValid = false
						break
					}
				}
			}

			if !isValid {
				break
			}
		}

		if isValid {
			id := strings.Split(gameID, " ")[1]
			num, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}

			sum += num
		}
	}

	fmt.Println("Sum of the power of these sets:", sum)
}
