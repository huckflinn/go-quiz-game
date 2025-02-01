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
}