package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// create parameters
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.Parse()
	// open the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintln("Failed to open csv file:", *csvFilename))
		os.Exit(1)
	}
	// csv read
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse csv file.")
	}
	// get problems
	problems := parseLines(lines)
	// start timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// count correct answers
	correct := 0
	// check every problem
problemloop:
	for i, problem := range problems {
		// ask
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			break problemloop
		case answer := <-answerCh:
			// check answer is correct
			if answer == problem.a {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Not correct :(")
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d! Good job!\n", correct, len(problems))

}

// parse lines
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

// one struct for every problem wherever we take it from
type problem struct {
	q, a string
}

// exit gracefully
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
