package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Record struct {
	question string
	answer   string
}

func main() {
	fmt.Println("Starting quiz...")

	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	r := csv.NewReader(bufio.NewReader(csvFile))
	r.FieldsPerRecord = 2

	records := []Record{}
	for {
		cur_record, read_err := r.Read()
		if read_err == io.EOF {
			break
		}
		if read_err != nil {
			log.Fatal(read_err)
		}

		rec := Record{question: cur_record[0], answer: cur_record[1]}
		records = append(records, rec)
	}

	//play game with the records now loaded 
	player_score := 0
	max_score := len(records)
	for _, rec := range records {
		fmt.Print("Question: ", rec.question, ": ")
		var input string
		fmt.Scanln(&input)
		input_ans, _ := strconv.Atoi(input)

		//check to see if it matches the answer
		correct_ans, _ := strconv.Atoi(rec.answer)
		if input_ans == correct_ans {
			player_score++
		}
	}

	fmt.Print("You scored ", player_score, " out of ", max_score, " points.\n")
}
