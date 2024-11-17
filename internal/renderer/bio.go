package renderer

import (
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/ceeyahya/resume_generator/internal/models"
)

func BioSection(m core.Maroto, bio models.Bio) {
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
