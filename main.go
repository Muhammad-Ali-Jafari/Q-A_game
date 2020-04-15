package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	fmt.Printf("Please enter questions file's name: ")

	var questionFile string
	_, err := fmt.Scanf("%s\n", &questionFile)

	file, err := os.Open(questionFile)
	if err != nil {
		panic(err.Error())
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Couldn't parse csv file! :(")
	}

	fmt.Printf("Please enter the timer(in seconds): ")

	var enteredTime int
	_, err = fmt.Scanf("%d", &enteredTime)
	if err != nil {
		panic(err.Error())
	}

	problems := parseLines(lines)
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)

		var answer string
		_, _ = fmt.Scanf("%s\n", &answer)
		if problem.a == answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
