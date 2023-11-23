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
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Can't parse CSV file")
	}

	problems := parseLines(lines)

	for i, p := range problems {
		isCorrect := false
		var userinput string
		for !isCorrect {
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			fmt.Scanf("%s\n", &userinput)
			if userinput == p.a {
				fmt.Printf("Correct!\n")
				isCorrect = true
			} else {
				fmt.Printf("Incorrect!\n")
			}

		}
	}
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
