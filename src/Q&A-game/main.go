package main

import (
	"encoding/csv"
	"fmt"
	"github.com/AlecAivazis/survey"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

var qs = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "What is your name?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "game",
		Prompt: &survey.Select{
			Message: "Choose a game:",
			Options: []string{"Math(easy)", "Math(medium)"},
			Default: "Math(easy)",
		},
	},
}

func main() {
	answers := struct {
		Name string
		Game string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var questionFile string
	switch answers.Game {
	case "Math(easy)":
		questionFile = "/home/nvsh116/projects/back/Q-A_game/src/Q&A-game/easy_problems.csv"
	case "Math(medium)":
		questionFile = "/home/nvsh116/projects/back/Q-A_game/src/Q&A-game/medium_problems.csv"
	}

	file, err := os.Open(questionFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Println("Couldn't parse csv file! :(")
	}

	fmt.Printf("Please enter the timer(in seconds): ")

	var timeLimit int
	_, err = fmt.Scanf("%d", &timeLimit)
	if err != nil {
		fmt.Println(err)
		return
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correct := 0
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == p.a {
				correct++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
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
