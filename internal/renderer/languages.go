package renderer

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/ceeyahya/resume_generator/internal/models"
)

func LanguagesSection(m core.Maroto, languages []models.Language) {
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
