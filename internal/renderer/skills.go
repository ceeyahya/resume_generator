package renderer

import (
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/ceeyahya/resume_generator/internal/helpers"
	"github.com/ceeyahya/resume_generator/internal/models"
)

func SkillsSection(m core.Maroto, skills models.Skills) {
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
	numRows := helpers.Max(numLangRows, numToolRows)

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
