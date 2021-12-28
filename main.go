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

	readQuestions(questions)
}

func readQuestions(questions []problem) {
	correct := 0
	index := 1
	for _, v := range questions {
		var ans string
		fmt.Printf("#%d - Qual a resposta para: %s ?", index, v.q)
		fmt.Scanf("%s", &ans)
		if ans != v.a {
			fmt.Println("Que pena, você errou!")
			break
		} else {
			fmt.Println("Parabéns, resposta correta.")
			correct++
		}
		index++
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
