package page

import "github.com/SShlykov/zeitment/bookback/internal/domain/entity"

type Options struct {
	Limit  uint64 `json:"limit,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
}

type requestModel struct {
	Options Options      `json:"options,omitempty"`
	Page    *entity.Page `json:"page,omitempty"`
}

type responseSingleModel struct {
	Page   *entity.Page `json:"page"`
	Status string       `json:"status"`
}

type responseListModel struct {
	Pages  []*entity.Page `json:"pages"`
	Status string         `json:"status"`
}
