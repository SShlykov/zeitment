package page

import "github.com/SShlykov/zeitment/bookback/internal/models"

type Options struct {
}

type requestModel struct {
	Options Options      `json:"options"`
	Page    *models.Page `json:"page"`
}

type responseSingleModel struct {
	Page   *models.Page `json:"page"`
	Status string       `json:"status"`
}

type responseListModel struct {
	Pages  []models.Page `json:"pages"`
	Status string        `json:"status"`
}
