package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	quiz := flag.String("quiz", "problems.csv", "The file for the quiz you wish to take.")

	flag.Parse()

	f, err := os.ReadFile(*quiz)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(f)))

	var solutions uint8

	var numRight uint8

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
	}
}