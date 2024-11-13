package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

func main() {
	resumeData, err := LoadResume("data/resume.json")
	if err != nil {
		log.Fatal(err)
	}

	marotoCore := GetMaroto()

	resume, err := GenerateResume(marotoCore, resumeData)
	if err != nil {
		log.Fatal(err)
	}

	err = resume.Save("docs/YahyaChahineResume.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithTopMargin(10).
		WithBottomMargin(10).
		WithLeftMargin(10).
		WithRightMargin(10).
		Build()

	m := maroto.New(cfg)

	return m
}

func GenerateResume(m core.Maroto, resume *Resume) (core.Document, error) {
	return m.Generate()
}

func LoadResume(filepath string) (*Resume, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("couldn't read the data file", err)
	}

	var resume Resume

	err = json.Unmarshal(file, &resume)
	if err != nil {
		return nil, fmt.Errorf("error while parsing the data file: %s", err)
	}

	return &resume, nil
}
