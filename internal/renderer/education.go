package renderer

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/ceeyahya/resume_generator/internal/models"
)

func EducationSection(m core.Maroto, education []models.Education) {
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
