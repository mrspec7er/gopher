package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/leekchan/accounting"
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

	decodedFile := csv.NewReader(file);
	totalSalary := 0;
	datacount := 0

	for {
		records, err := decodedFile.Read()

		if err == io.EOF {
			break	
		}

		if err != nil {
			log.Fatal(err)
		}		

		filteredRecord := handleFilterRecord(records);
		salary, err := strconv.Atoi(records[6]);
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(salary)
		totalSalary += salary;
		datacount += 1
		
		data = append(data, filteredRecord)

	}
	
	// for _, eachData := range data {
	// 	fmt.Printf("%v\n", eachData)
	// }

	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2}

	fmt.Println("AVERAGE: ", totalSalary/datacount)
	fmt.Println("AVARAGE in IDR: ", ac.FormatMoney(totalSalary/datacount * 15000))
	fmt.Println("TOTAL: ", totalSalary);
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