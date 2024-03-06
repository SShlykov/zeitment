package bookevents

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Options struct{}

type requestModel struct {
	Options    Options           `json:"options,omitempty"`
	BookEvents *entity.BookEvent `json:"book_events,omitempty"`
}

type responseSingleModel struct {
	BookEvent *entity.BookEvent `json:"book_event"`
	Status    string            `json:"status"`
}

type responseListModel struct {
	BookEvents []*entity.BookEvent `json:"book_events"`
	Status     string              `json:"status"`
}
