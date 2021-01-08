package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	csvFile, err := os.Open("oscar_age_male.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(csvFile)

	const nameColumn = 3
	actorMap := make(map[string]int)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		actorName := record[nameColumn]
		actorMap[actorName]++
	}

	for name, count := range actorMap {
		if count > 1 {
			fmt.Println(name)
		}
	}
}
