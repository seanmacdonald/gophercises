package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Record struct {
	question string
	answer   string
}

func main() {
	//first check flags to determine which game file to use
	var fname = flag.String("csv", "problems.csv", "The file that we get the questions and answers from")
	var time_limit = flag.Int("timer", 30, "Time limit for the quiz.")
	flag.Parse()
	fmt.Println(*fname)
	fmt.Println(*time_limit)

	csvFile, err := os.Open(*fname)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	r := csv.NewReader(bufio.NewReader(csvFile))
	r.FieldsPerRecord = 2

	//records := []Record{}
	records := make([]Record, 0, 100)
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

	//play game with the records now loaded and start timer
	end_time := make(chan string)
	go timer(*time_limit, end_time)

	player_score := 0
	max_score := len(records)

GameLoop:
	for num, rec := range records {
		fmt.Print("Question ", num + 1, ": ", rec.question, ": ")
		c := make(chan string)
		var input_ans int
		go input(c)

		//wait for time to end or user input
	IdleLoop:
		for {
			select {
			case x := <-c:
				input_ans, _ = strconv.Atoi(x)
				break IdleLoop

			case y := <-end_time:
				if y == "END" {
					fmt.Println("\nTime is up.")
					break GameLoop
				}
			}

		}

		//check to see if it matches the answer
		correct_ans, _ := strconv.Atoi(rec.answer)
		if input_ans == correct_ans {
			player_score++
		}
	}

	fmt.Print("You scored ", player_score, " out of ", max_score, " points.\n")
}

func input(c chan string) {
	var in string
	fmt.Scanln(&in)
	c <- in
	close(c)
}

func timer(seconds int, c chan string) {
	var val time.Duration = time.Duration(seconds * 1000)
	time.Sleep(val * time.Millisecond)
	c <- "END"
}
