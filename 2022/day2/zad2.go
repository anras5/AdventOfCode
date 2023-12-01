package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Zad2() {
	file, err := os.Open("day2/data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	rules := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	for scanner.Scan() {

		score += rules[scanner.Text()]

	}

	fmt.Println(score)
}