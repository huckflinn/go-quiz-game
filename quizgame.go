package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	f, err := os.ReadFile("./QuizInput.csv")

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
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(row[0])
	}
}