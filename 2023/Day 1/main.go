package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numberWords = map[string]int{
	"zero":  0,
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

func main() {

	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a slice to hold the lines
	var lines []string

	// Read and store each line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var digits []string
	var outputTotal int

	// Print each line
	for i, line := range lines {
		fmt.Printf("Line %d: %s\n", i+1, line)

		characters := strings.Split(line, "")

		for j, c := range characters {

			if _, err := strconv.Atoi(c); err == nil {
				digits = append(digits, c)
			} else if c == "o" {
				numResult := checkNumWord(j, line, "one")
				if numResult != "" {
					digits = append(digits, numResult)
				}
			} else if c == "t" {
				numResult := checkNumWord(j, line, "two")
				if numResult != "" {
					digits = append(digits, numResult)
				}

				numResult = checkNumWord(j, line, "three")
				if numResult != "" {
					digits = append(digits, numResult)
				}

			} else if c == "f" {
				numResult := checkNumWord(j, line, "four")
				if numResult != "" {
					digits = append(digits, numResult)
				}

				numResult = checkNumWord(j, line, "five")
				if numResult != "" {
					digits = append(digits, numResult)
				}

			} else if c == "s" {
				numResult := checkNumWord(j, line, "six")
				if numResult != "" {
					digits = append(digits, numResult)
				}

				numResult = checkNumWord(j, line, "seven")
				if numResult != "" {
					digits = append(digits, numResult)
				}

			} else if c == "e" {
				numResult := checkNumWord(j, line, "eight")
				if numResult != "" {
					digits = append(digits, numResult)
				}

			} else if c == "n" {
				numResult := checkNumWord(j, line, "nine")
				if numResult != "" {
					digits = append(digits, numResult)
				}
			}
		}

		fmt.Println("digits: ", digits)

		if len(digits) == 2 {
			outputTotal += createNumber(digits[0], digits[1])
		} else if len(digits) > 2 {
			outputTotal += createNumber(digits[0], digits[len(digits)-1])
		} else if len(digits) == 1 {
			outputTotal += createNumber(digits[0], digits[0])
		}

		digits = nil
		fmt.Println("outputTotal: ", outputTotal)

	}
}

func createNumber(digit1 string, digit2 string) int {
	str1 := digit1
	str2 := digit2

	num := str1 + str2

	// Convert str1 to an integer
	output, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("output: ", output)
	return output
}

func checkNumWord(index int, line string, numWord string) string {
	if len(line) >= index+len(numWord) {
		if line[index:index+len(numWord)] == numWord {
			fmt.Println("adding ", numWord)
			return strconv.Itoa(numberWords[numWord])
		}
	}

	return ""
}
