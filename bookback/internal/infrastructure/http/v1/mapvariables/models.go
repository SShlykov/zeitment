package mapvariables

import "github.com/SShlykov/zeitment/bookback/internal/domain/entity"

type Options struct{}

type requestModel struct {
	Options      Options             `json:"options,omitempty"`
	MapVariables *entity.MapVariable `json:"map_variables,omitempty"`
}

type responseSingleModel struct {
	MapVariable *entity.MapVariable `json:"map_variable"`
	Status      string              `json:"status"`
}

type responseListModel struct {
	MapVariables []*entity.MapVariable `json:"map_variables"`
	Status       string                `json:"status"`
}
