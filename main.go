package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of the exercise: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}

	switch number {
	case 1:
		Advent1()
	default:
		fmt.Println("Exercise not found.")
	}
}
