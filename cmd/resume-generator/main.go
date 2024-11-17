package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ceeyahya/resume_generator/internal/models"
	"github.com/ceeyahya/resume_generator/internal/pdf"
)

func main() {
	resumeData, err := LoadResume("data/resume.json")
	if err != nil {
		log.Fatal(err)
	}

	marotoCore := pdf.GetMaroto()

	resume, err := pdf.GenerateResume(marotoCore, resumeData)
	if err != nil {
		log.Fatal(err)
	}

	err = resume.Save("docs/YahyaChahineResume.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func LoadResume(filepath string) (*models.Resume, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("couldn't read the data file", err)
	}

	var resume models.Resume

	err = json.Unmarshal(file, &resume)
	if err != nil {
		return nil, fmt.Errorf("error while parsing the data file: %s", err)
	}

	return &resume, nil
}
