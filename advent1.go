package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func Advent1() {
	f, err := os.Open("advent_1.txt")

	if err != nil{
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := []int{}

	for scanner.Scan(){
		line := scanner.Text()
		
		firstDigitIndex, lastDigitIndex := -1, -1

		for i := range line{
			if line[i] >= '0' && line[i] <= '9'{
				firstDigitIndex = i
				break
			}
		}

		for i := len(line) - 1; i > firstDigitIndex; i--{
			if line[i] >= '0' && line[i] <= '9'{
				lastDigitIndex = i
				break
			}
		}

		if lastDigitIndex != -1 {
			number, err := strconv.Atoi((string(line[firstDigitIndex])+string(line[lastDigitIndex])))
			if err != nil{
				fmt.Println(err)
			} else {
				numbers = append(numbers, number)
			}
		}else {
			number, err := strconv.Atoi(string(line[firstDigitIndex]))
			if err != nil{
				fmt.Println(err)
			} else {
				numbers = append(numbers, number)
			}
		}
	}

	fmt.Println(sumSlice(numbers))
}