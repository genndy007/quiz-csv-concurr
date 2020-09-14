package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// create parameters
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
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
	//fmt.Println(problems)
	correct := 0
	// check every problem
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		// check answer is correct
		if answer == problem.a {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Not correct :(")
		}
	}
	fmt.Printf("You scored %d out of %d! Good job!\n", correct, len(problems))

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
