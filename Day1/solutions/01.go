package solutions

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type intTuple struct {
	a, b int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1() int {
	leftList := []int{}
	rightList := []int{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		halves := strings.Split(line, "   ")

		left, err := strconv.Atoi(halves[0])
		check(err)
		right, err := strconv.Atoi(halves[1])
		check(err)

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	pairs, err := utils.Zip(leftList, rightList)
	check(err)

	distances := []int{}

	for _, pair := range pairs {
		distance := utils.Abs(pair.A - pair.B)
		distances = append(distances, distance)
	}

	sum := 0
	for _, dist := range distances {
		sum += dist
	}

	return sum
}
