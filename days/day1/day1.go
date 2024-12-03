package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func GetDistance(left, right int) int {
  return int(math.Abs(float64(left) - float64(right)))
}

func part1() int {
  file, _ := os.Open("input.txt")
  defer file.Close()

  left := []int{}
  right := []int{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    nums := strings.Split(line, " ")

    numLeft, _ := strconv.Atoi(nums[0])
    numRight, _ := strconv.Atoi(nums[1])

    left = append(left, numLeft)
    right = append(right, numRight)
  }

  slices.Sort(left)
  slices.Sort(right)

  totalDistances := 0
  for idx := 0; idx < len(left); idx++ {
    totalDistances += GetDistance(left[idx], right[idx])
  }

  return totalDistances
}

func part2() int {
  file, _ := os.Open("input.txt")
  defer file.Close()

  left := []int{}
  right := map[int]int{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    nums := strings.Split(line, " ")

    numLeft, _ := strconv.Atoi(nums[0])
    numRight, _ := strconv.Atoi(nums[1])

    left = append(left, numLeft)
    right[numRight] = right[numRight] + 1
  }

  slices.Sort(left)

  similarityScore := 0
  for idx := 0; idx < len(left); idx++ {
    leftNum := left[idx]
    count := right[leftNum]

    similarityScore += leftNum * count
  }

  return similarityScore
}

func main() {
  fmt.Println(part1())
  fmt.Println(part2())
}
