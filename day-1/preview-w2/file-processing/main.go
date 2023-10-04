package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[ERROR] :", r)
		}
	}()

	args := os.Args
	if len(args) != 3 {
		log.Fatal("Masukkan argumen <input.csv> <output.csv>")
	}

	// Regexp - check ekstensi file dalam argumen
	regx := regexp.MustCompile(`^\w+.csv$`)
	if !regx.MatchString(args[1]) && !regx.MatchString(args[2]) {
		log.Fatal("Ekstensi file harus .csv")
	}

	// Reader file input
	inputFile, err := os.Open(args[1])
	if err != nil {
		panic("File not exist")
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	failOnError(err, "Failed to read all records")

	wg := sync.WaitGroup{}
	updatedRecords := make(chan []string)
	// defer close(updatedRecords)

	for _, record := range records[1:] {
		wg.Add(1)
		go func(record []string) {
			defer wg.Done()
			updatedRecords <- []string{
				strings.ToUpper(record[0]),
				record[1],
				"Mr." + record[2],
			}
		}(record)
	}

	go func() {
		wg.Wait()
		close(updatedRecords)
	}()

	// wg.Wait()

	outputFilePath := os.Args[2]
	generateOutput(outputFilePath, updatedRecords)

}

func generateOutput(outputPath string, records chan []string) {
	outputFile, err := os.Create(outputPath)
	failOnError(err, "Failed to create output file")

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	writer.Write([]string{"Name", "Age", "Occupation"}) // CSV Header

	// proses chan : updated record
	for record := range records {
		writer.Write(record)
		failOnError(err, fmt.Sprintf("Failed to write record: %v", record))
	}
}
