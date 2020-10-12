//Data IO

package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Entry struct {
	IdTimeEntry        string `json:"idTimeEntry"`
	Notes     string `json:"notes"`
	CreateDate string `json:"createDate"`
	DueDate string `json:"dueDate"`
}

var entries []Entry

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func readData(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	entries = []Entry{}

	for _, record := range records {
		entry := Entry{
			IdTimeEntry:        record[0],
			Notes:     record[1],
			CreateDate:   record[2],
			DueDate: record[3]}
		entries = append(entries, entry)
	}
	file.Close()
}

func writeData(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, entry := range entries {
		record := []string{entry.IdTimeEntry, entry.Notes, entry.CreateDate,
			entry.DueDate}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
