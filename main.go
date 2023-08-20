package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type SalaryType struct {
	experience string
	jobTitle string
	salary string
	location string
	companySize string
}

func main()  {
	file, err := os.Open("./salary.csv");
	data := []SalaryType{}

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	decodedFile := csv.NewReader(file)

	for {
		records, err := decodedFile.Read()

		if err == io.EOF {
			break	
		}

		if err != nil {
			log.Fatal(err)
		}		

		filteredRecord := handleFilterRecord(records)
		
		
		data = append(data, filteredRecord)

	}
	
	for _, eachData := range data {
		fmt.Printf("%v\n", eachData)
	}

	fmt.Println(data[0])
}

func handleFilterRecord(records []string) SalaryType  {
	filteredRecord := SalaryType {}

	for i, record := range records {
			if i == 1 {
				filteredRecord.experience = record
			}
			if i == 3 {
				filteredRecord.jobTitle = record
			}
			if i == 6 {
				filteredRecord.salary = record
			}
			if i == 7 {
				filteredRecord.location = record
			}
			if i == 8 {
				filteredRecord.companySize = record
			}
		}

	return filteredRecord
}