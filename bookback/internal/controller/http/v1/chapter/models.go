package chapter

import "github.com/SShlykov/zeitment/bookback/internal/models"

type Options struct {
}

type requestModel struct {
	Options Options         `json:"options"`
	Chapter *models.Chapter `json:"chapter"`
}

type responseSingleModel struct {
	Chapter *models.Chapter `json:"chapter"`
	Status  string          `json:"status"`
}

type responseListModel struct {
	Chapters []models.Chapter `json:"chapters"`
	Status   string           `json:"status"`
}
