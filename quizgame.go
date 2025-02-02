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

func main() {
	quiz := flag.String("quiz", "problems.csv", "The file for the quiz you wish to take.")

	timeLimit := flag.Int("time-limit", 30, "The amount of time to take the quiz.")

	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	f, err := os.ReadFile(*quiz)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(f)))

	var solutions uint8

	var numRight uint8

	go func() {
		<-timer.C
		fmt.Println("Time's up! Thanks for playing!")
		os.Exit(0)
	} ()

	for {
		row, err := r.Read()

		if err == io.EOF {
			fmt.Printf("Total Questions: %d\n", solutions)
			fmt.Printf("Correct Answers: %d\n", numRight)
			fmt.Println("Thanks for playing! Goodbye!")
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
			numRight ++
		}

		solutions++
	}

	timer.Stop()
}