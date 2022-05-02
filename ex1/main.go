package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type QuestionRecord struct {
	Question string
	Answer   int
	Correct  bool
}

func main() {

	filePtr := flag.String("file", "problems.csv", "Provide a csv file")
	timePtr := flag.Int("time", 30, "Please provide a Quiz time in seconds")

	flag.Parse()

	QuestionSheet := createSheet(*filePtr)

	fmt.Println("I'm going to ask you some questions. Once you press 'Enter' the quiz will start.")
	fmt.Scan()
	go func() {
		<-time.After(time.Duration(*timePtr * int(time.Second)))
		fmt.Println("Timer expired")
		checkResult(QuestionSheet)
		os.Exit(0)
	}()

	startQuiz(QuestionSheet)
	checkResult(QuestionSheet)
}

func createSheet(file string) []QuestionRecord {
	var QuestionSheet []QuestionRecord

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range data {
		var rec QuestionRecord
		for i, field := range line {
			if i == 0 {
				rec.Question = field
			} else if i == 1 {
				rec.Answer, err = strconv.Atoi(field)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		QuestionSheet = append(QuestionSheet, rec)
	}
	return QuestionSheet
}

func startQuiz(QuestionSheet []QuestionRecord) []QuestionRecord {
	var answer int
	for i, rec := range QuestionSheet {
		fmt.Printf("What is %q?\n", rec.Question)
		fmt.Scan(&answer)
		(&QuestionSheet[i]).Correct = answer == rec.Answer
	}
	return QuestionSheet
}

func checkResult(QuestionSheet []QuestionRecord) {
	var questionNo int
	var sum int
	for _, rec := range QuestionSheet {
		questionNo += 1
		if rec.Correct {
			sum += 1
		}
	}
	fmt.Printf("Answered %v correct out of %v questions.\n", sum, questionNo)
}
