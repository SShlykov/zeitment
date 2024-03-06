package book

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Options struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type requestModel struct {
	Options Options      `json:"options"`
	Book    *entity.Book `json:"book"`
}

type responseSingleModel struct {
	Book   *entity.Book `json:"book"`
	Status string       `json:"status"`
}

type responseListModel struct {
	Books  []entity.Book `json:"books"`
	Status string        `json:"status"`
}
