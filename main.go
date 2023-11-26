package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("timeLimit", 60, "the amount of time of the test to happen")

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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	<-timer.C
	problems := parseLines(lines)
	numCorrect := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var userinput string
			fmt.Scanf("%s\n", &userinput)
		}()
		select {
		case <-timer.C:
			pct := float64(numCorrect) / float64(len(problems)) * 100
			fmt.Printf("You managed to answer %d correct problems out of %d, your percentage is %f \n", numCorrect, len(problems), pct)
			return
		case answer := <-answerCh:
			go func() {
				if answer == p.a {
					fmt.Printf("Correct!\n")
					numCorrect++
				} else {
					fmt.Printf("Incorrect!\n")
				}
			}()
		}
	}
	pct := float64(numCorrect) / float64(len(problems)) * 100
	fmt.Printf("You managed to answer %d correct problems out of %d, your percentage is %f\n", numCorrect, len(problems), pct)

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
