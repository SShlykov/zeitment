package paragraph

import "github.com/SShlykov/zeitment/bookback/internal/domain/entity"

type Options struct {
	Limit  uint64 `json:"limit,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
}

type requestModel struct {
	Options   Options           `json:"options,omitempty"`
	Paragraph *entity.Paragraph `json:"paragraph,omitempty"`
}

type responseSingleModel struct {
	Paragraph *entity.Paragraph `json:"paragraph"`
	Status    string            `json:"status"`
}

type responseListModel struct {
	Paragraphs []*entity.Paragraph `json:"paragraphs"`
	Status     string              `json:"status"`
}
