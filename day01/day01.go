package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part string) {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var answer []int

	replace := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := make([]string, 2)

		if part == "part2" {
			for k, v := range replace {
				line = strings.Replace(line, k, v, -1)
			}
		}

		for _, char := range line {
			val := string(char)
			_, err := strconv.Atoi(val)
			if err != nil {
				continue
			}

			numbers[1] = val
			if numbers[0] == "" {
				numbers[0] = val
			}
		}

		number, _ := strconv.Atoi(strings.Join(numbers, ""))
		answer = append(answer, number)
	}

	var sum int
	for _, a := range answer {
		sum += a
	}
	log.Println(sum)
}
