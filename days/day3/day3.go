package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	instructionRe := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	stock := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions := instructionRe.FindAllString(line, -1)
		for _, instruction := range instructions {
			var num1, num2 int
			fmt.Sscanf(instruction, "mul(%d,%d)", &num1, &num2)
			stock += num1 * num2
		}
	}

	return stock
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructionRe := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	stock := 0
	nextDisabled := false
	for scanner.Scan() {
		line := scanner.Text()
		instructions := instructionRe.FindAllString(line, -1)
		for _, instruction := range instructions {
			if instruction == "don't()" {
				nextDisabled = true
				continue
			} else if instruction == "do()" {
				nextDisabled = false
				continue
			} else if nextDisabled == false {
				var num1, num2 int
				fmt.Sscanf(instruction, "mul(%d,%d)", &num1, &num2)
				stock += num1 * num2
			}
		}
	}

	return stock
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
