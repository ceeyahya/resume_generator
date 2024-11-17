package main

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const (
	// Base Spacing
	baseSpacing = 2.0

	// Font Sizes
	textXL   = 18.0
	textL    = 14.0
	textM    = 12.0
	textBase = 11.0
	textS    = 10.0

	// Items Per Row
	languagesPerRow       = 2
	toolsPerRow           = 4
	softSkillsPerRow      = 4
	spokenLanguagesPerRow = 3

	// Vertical spacing between elements
	spacingXS = baseSpacing     // 2mm  - Minimal spacing (within same section)
	spacingS  = baseSpacing * 2 // 4mm  - Small spacing (between related elements)
	spacingM  = baseSpacing * 3 // 6mm  - Medium spacing (between sections)
	spacingL  = baseSpacing * 4 // 8mm  - Large spacing (major sections)
	spacingXL = baseSpacing * 6 // 12mm - Extra large spacing (top/bottom margins)
)

func RenderBioSection(m core.Maroto, bio Bio) {
	m.AddAutoRow(
		text.NewCol(12, bio.Name, props.Text{
			Size:  textXL,
			Style: fontstyle.Bold,
		}),
	)

	m.AddAutoRow(
		text.NewCol(12, bio.Position, props.Text{
			Size:  textBase,
			Style: fontstyle.Normal,
			Top:   baseSpacing,
		}),
	)

	m.AddAutoRow(
		text.NewCol(2, bio.Contact.PhoneNumber, props.Text{
			Size:  textBase,
			Style: fontstyle.Normal,
			Top:   baseSpacing,
		}),

		text.NewCol(3, bio.Contact.Email, props.Text{
			Size:  textBase,
			Style: fontstyle.Normal,
			Top:   baseSpacing,
		}),

		text.NewCol(1, "Github", props.Text{
			Size:      textBase,
			Style:     fontstyle.Normal,
			Top:       baseSpacing,
			Hyperlink: &bio.Social.Github,
		}),

		text.NewCol(1, "LinkedIn", props.Text{
			Size:      textBase,
			Style:     fontstyle.Normal,
			Top:       baseSpacing,
			Hyperlink: &bio.Social.Linkedin,
		}),

		text.NewCol(2, "Personal Website", props.Text{
			Size:      textBase,
			Style:     fontstyle.Normal,
			Top:       baseSpacing,
			Hyperlink: &bio.Social.Portfolio,
		}),
	)

	m.AddRow(spacingL)
}

func RenderEducationSection(m core.Maroto, education []Education) {
	m.AddAutoRow(
		text.NewCol(12, "Education", props.Text{
			Size:  textL,
			Style: fontstyle.Bold,
		}),
	)

	m.AddRow(spacingS)

	for i, e := range education {
		m.AddAutoRow(
			text.NewCol(12, e.Institution, props.Text{
				Size:  textM,
				Style: fontstyle.Bold,
			}),
		)

		m.AddAutoRow(
			text.NewCol(
				6,
				fmt.Sprintf("%s in %s — %s", e.Credential, e.Field, e.GraduationYear),
				props.Text{
					Size: textS,
					Top:  1,
				},
			),

			text.NewCol(
				6,
				fmt.Sprintf("%s, %s", e.Location.City, e.Location.Country),
				props.Text{
					Size:  textS,
					Align: align.Right,
					Top:   1,
				},
			),
		)

		if e.Thesis != "" {
			m.AddAutoRow(text.NewCol(12, e.Thesis, props.Text{
				Size: textBase,
				Top:  baseSpacing,
			}))
		}

		if i < len(education)-1 {
			m.AddRow(spacingS)
		}

	}

	m.AddRow(spacingL)
}

func RenderExperienceSection(m core.Maroto, experience []WorkExperience) {
	m.AddAutoRow(text.NewCol(12, "Work Experience", props.Text{
		Size:  textL,
		Style: fontstyle.Bold,
	}))

	m.AddRow(spacingS)

	for i, e := range experience {
		m.AddAutoRow(
			text.NewCol(6, e.Company, props.Text{
				Size:  textM,
				Style: fontstyle.Bold,
			}),

			text.NewCol(
				6,
				fmt.Sprintf("%s, %s", e.Location.City, e.Location.Country),
				props.Text{
					Size:  textS,
					Align: align.Right,
				},
			),
		)
		m.AddAutoRow(
			text.NewCol(
				12,
				fmt.Sprintf("%s, %s — %s", e.Position, e.Period.Start, e.Period.End),
				props.Text{
					Size: textS,
					Top:  1,
				},
			),
		)

		for _, highlight := range e.Highlights {
			m.AddAutoRow(
				text.NewCol(12, fmt.Sprintf("• %s", highlight), props.Text{
					Size: textBase,
					Top:  baseSpacing,
				}),
			)
		}

		if i < len(experience)-1 {
			m.AddRow(spacingS)
		}
	}

	m.AddRow(spacingL)
}

func RenderSkillsSection(m core.Maroto, skills Skills) {
	m.AddAutoRow(
		text.NewCol(4, "Languages", props.Text{
			Size:  textL,
			Style: fontstyle.Bold,
		}),
		text.NewCol(8, "Tools", props.Text{
			Size:  textL,
			Style: fontstyle.Bold,
		}),
	)

	m.AddRow(spacingS)

	// Calculate how many rows we need
	numLangRows := (len(skills.Technical.Languages) + languagesPerRow - 1) / languagesPerRow
	numToolRows := (len(skills.Technical.Tools) + toolsPerRow - 1) / toolsPerRow
	numRows := Max(numLangRows, numToolRows)

	// Render each row
	for i := 0; i < numRows; i++ {
		// Calculate indices for this row
		langStart := i * languagesPerRow
		toolStart := i * toolsPerRow

		var rowCols []core.Col

		for j := 0; j < languagesPerRow; j++ {
			idx := langStart + j
			if idx < len(skills.Technical.Languages) {
				rowCols = append(rowCols,
					text.NewCol(2, skills.Technical.Languages[idx], props.Text{
						Size: textBase,
					}),
				)
			} else {
				// Empty column to maintain grid
				rowCols = append(rowCols, text.NewCol(2, "", props.Text{}))
			}
		}

		// Add tool columns (4 per row)
		for j := 0; j < toolsPerRow; j++ {
			idx := toolStart + j
			if idx < len(skills.Technical.Tools) {
				rowCols = append(rowCols,
					text.NewCol(3, skills.Technical.Tools[idx], props.Text{
						Size: textBase,
					}),
				)
			} else {
				// Empty column to maintain grid
				rowCols = append(rowCols, text.NewCol(3, "", props.Text{}))
			}
		}

		m.AddAutoRow(rowCols...)

		if i < numRows-1 {
			m.AddRow(baseSpacing)
		}
	}

	m.AddRow(spacingL)

	m.AddAutoRow(
		text.NewCol(12, "Soft Skills", props.Text{
			Size:  textL,
			Style: fontstyle.Bold,
		}),
	)

	numSoftRows := (len(skills.Soft) + softSkillsPerRow - 1) / softSkillsPerRow

	m.AddRow(spacingS)

	for i := 0; i < numSoftRows; i++ {
		softStart := i * softSkillsPerRow
		var rowCols []core.Col

		for j := 0; j < softSkillsPerRow; j++ {
			idx := softStart + j
			if idx < len(skills.Soft) {
				rowCols = append(rowCols,
					text.NewCol(3, skills.Soft[idx], props.Text{
						Size: textBase,
					}),
				)
			} else {
				rowCols = append(rowCols, text.NewCol(3, "", props.Text{}))
			}
		}

		m.AddAutoRow(rowCols...)

		if i < numSoftRows-1 {
			m.AddRow(baseSpacing)
		}
	}

	m.AddRow(spacingL)
}

func RenderLanguagesSection(m core.Maroto, languages []Language) {
	m.AddAutoRow(
		text.NewCol(12, "Spoken Languages", props.Text{
			Size:  textL,
			Style: fontstyle.Bold,
		}),
	)

	m.AddRow(spacingS)

	numRows := (len(languages) + 2) / spokenLanguagesPerRow

	for i := 0; i < numRows; i++ {
		langStart := i * 3
		var rowCols []core.Col

		for j := 0; j < 3; j++ {
			idx := langStart + j
			if idx < len(languages) {
				rowCols = append(rowCols,
					text.NewCol(4,
						fmt.Sprintf("%s: %s",
							languages[idx].Name,
							languages[idx].Profeciency,
						),
						props.Text{
							Size: textBase,
						},
					),
				)
			} else {
				rowCols = append(rowCols, text.NewCol(4, "", props.Text{}))
			}
		}

		m.AddAutoRow(rowCols...)

		if i < numRows-1 {
			m.AddRow(spacingS)
		}
	}
}
