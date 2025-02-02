package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func showScore(questions uint8, answers uint8) {
	fmt.Printf("\nTotal Questions: %d\n", questions)
	fmt.Printf("Correct Answers: %d\n", answers)
	fmt.Println("Thanks for playing! Goodbye!")
}



func main() {
	quiz := flag.String("quiz", "problems.csv", "The file for the quiz you wish to take.")

	timeLimit := flag.Int("time-limit", 30, "The amount of time to take the quiz.")

	flag.Parse()
	
	for {
		fmt.Print("Press 'Enter' to begin:")
		fmt.Scanln()
		break
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	f, err := os.ReadFile(*quiz)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(f)))

	var totalQuestions uint8

	var totalCorrect uint8

	go func() {
		<-timer.C
		fmt.Println("\nTime's up!")
		showScore(totalQuestions, totalCorrect)
		os.Exit(0)
	} ()


	for {
		row, err := r.Read()

		if err == io.EOF {
			showScore(totalQuestions, totalCorrect)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(row[0])

		var res int16
		fmt.Scanln(&res)

		ans64, err := strconv.ParseInt(row[1], 10, 16)

		if err != nil {
			log.Fatal(err)
			fmt.Println("Invalid response. Please enter a number.")
			continue
		}

		ans := int16(ans64)

		if res == ans {
			totalCorrect ++
		}

		totalQuestions++
	}

	timer.Stop()
}