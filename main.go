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
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	questions := fillQuiz(csvLines)

	readQuestions(questions, *timer)
}

func readQuestions(questions []problem, timer time.Timer) {
	correct := 0

	for i, v := range questions {
		fmt.Printf("#%d - Qual a resposta para: %s ?\n", i+1, v.q)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nVocê acertou %d de %d.\n", correct, len(questions))
			return
		case ans := <-answerCh:
			if ans == v.a {
				correct++
			}
		}

	}
	fmt.Printf("Você acertou %d de %d.\n", correct, len(questions))
}

func fillQuiz(lines [][]string) []problem {
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
