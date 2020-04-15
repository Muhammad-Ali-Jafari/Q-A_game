package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "some csv file that contains Q&As")
	tie := flag.Int("time", 30, "time needed to answer the questions")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Couldn't parse csv file! :(")
	}

	problems := parseLines(lines)
	correct := 0
	for i, problem := range problems{
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
