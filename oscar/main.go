package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFile, err := os.Open("oscar_age_male.csv")
	csvReader := csv.NewReader(csvFile)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	actorMap := make(map[string]int)
	for i := 0; i < len(records); i++ {
		row := records[i]
		actorName := row[3]
		actorMap[actorName]++
	}

	for name, count := range actorMap {
		if count > 1 {
			fmt.Println(name)
		}
	}
}
