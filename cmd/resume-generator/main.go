package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ceeyahya/resume_generator/internal/pdf"
)

func main() {
	var (
		dataFilePath   string
		outputFileName string
	)

	flag.StringVar(&dataFilePath, "filepath", "", "Path to the data file")
	flag.StringVar(&outputFileName, "output", "", "Name of the output file")

	flag.Parse()

	if dataFilePath == "" {
		log.Fatal("the filepath flag is required")
	}

	if _, err := os.Stat(dataFilePath); os.IsNotExist(err) {
		log.Fatalf("datafile does not exist: %s", dataFilePath)
	}

	if err := os.MkdirAll("docs", 0755); err != nil {
		log.Fatalf("could not create output directory: %v", err)
	}

	resumeData, err := pdf.LoadResume(dataFilePath)
	if err != nil {
		log.Fatal(err)
	}

	marotoCore := pdf.GetMaroto()

	resume, err := pdf.GenerateResume(marotoCore, resumeData)
	if err != nil {
		log.Fatal(err)
	}

	err = resume.Save(fmt.Sprintf("docs/%s.pdf", outputFileName))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Resume generated successfully: %s.pdf\n", outputFileName)
}
