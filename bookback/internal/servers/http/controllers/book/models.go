package book

import "github.com/SShlykov/zeitment/bookback/internal/models"

type Options struct {
}

type requestModel struct {
	Options Options      `json:"options"`
	Book    *models.Book `json:"book"`
}

type responseSingleModel struct {
	Book   *models.Book `json:"book"`
	Status string       `json:"status"`
}

type responseListModel struct {
	Books  []models.Book `json:"books"`
	Status string        `json:"status"`
}
