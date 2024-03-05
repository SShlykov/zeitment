package paragraph

import "github.com/SShlykov/zeitment/bookback/internal/models"

type Options struct {
}

type requestModel struct {
	Options   Options           `json:"options"`
	Paragraph *models.Paragraph `json:"paragraph"`
}

type responseSingleModel struct {
	Paragraph *models.Paragraph `json:"paragraph"`
	Status    string            `json:"status"`
}

type responseListModel struct {
	Paragraphs []models.Paragraph `json:"paragraphs"`
	Status     string             `json:"status"`
}
