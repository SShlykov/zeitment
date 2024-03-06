package chapter

import "github.com/SShlykov/zeitment/bookback/internal/domain/entity"

type Options struct {
	Limit  uint64 `json:"limit,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
}

type requestModel struct {
	Options Options         `json:"options,omitempty"`
	Chapter *entity.Chapter `json:"chapter,omitempty"`
}

type responseSingleModel struct {
	Chapter *entity.Chapter `json:"chapter"`
	Status  string          `json:"status"`
}

type responseListModel struct {
	Chapters []*entity.Chapter `json:"chapters"`
	Status   string            `json:"status"`
}
