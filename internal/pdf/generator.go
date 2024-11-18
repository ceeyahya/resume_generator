package pdf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"

	"github.com/ceeyahya/resume_generator/internal/models"
	"github.com/ceeyahya/resume_generator/internal/renderer"
)

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

func GetMaroto() core.Maroto {
	fontFamily := "Heliotrope"

	customFonts, err := repository.New().
		AddUTF8Font(fontFamily, fontstyle.Normal, "assets/fonts/HeliotropeBook.ttf").
		AddUTF8Font(fontFamily, fontstyle.Italic, "assets/fonts/HeliotropeBookItalic.ttf").
		AddUTF8Font(fontFamily, fontstyle.Bold, "assets/fonts/HeliotropeSemibold.ttf").
		AddUTF8Font(fontFamily, fontstyle.BoldItalic, "assets/fonts/HeliotropeSemiboldItalic.ttf").
		Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithTopMargin(10).
		WithBottomMargin(10).
		WithLeftMargin(10).
		WithRightMargin(10).
		WithCustomFonts(customFonts).
		WithDefaultFont(&props.Font{Family: fontFamily}).
		Build()

	m := maroto.New(cfg)

	return m
}

func GenerateResume(m core.Maroto, resume *models.Resume) (core.Document, error) {
	renderer.BioSection(m, resume.Bio)
	renderer.EducationSection(m, resume.Education)
	renderer.WorkExperienceSection(m, resume.WorkExperience)
	renderer.SkillsSection(m, resume.Skills)
	renderer.LanguagesSection(m, resume.Languages)

	return m.Generate()
}
