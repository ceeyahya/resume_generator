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

func WorkExperienceSection(m core.Maroto, experience []models.WorkExperience) {
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
