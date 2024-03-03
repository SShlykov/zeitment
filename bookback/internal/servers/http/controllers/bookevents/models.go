package bookevents

import "github.com/SShlykov/zeitment/bookback/internal/models"

type Options struct{}

type requestModel struct {
	Options    Options           `json:"options"`
	BookEvents *models.BookEvent `json:"book_events"`
}

type responseSingleModel struct {
	BookEvent *models.BookEvent `json:"book_event"`
	Status    string            `json:"status"`
}

type responseListModel struct {
	BookEvents []models.BookEvent `json:"book_events"`
	Status     string             `json:"status"`
}
