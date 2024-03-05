package mapvariables

import "github.com/SShlykov/zeitment/bookback/internal/models"

type Options struct{}

type requestModel struct {
	Options      Options             `json:"options"`
	MapVariables *models.MapVariable `json:"map_variables"`
}

type responseSingleModel struct {
	MapVariable *models.MapVariable `json:"map_variable"`
	Status      string              `json:"status"`
}

type responseListModel struct {
	MapVariables []models.MapVariable `json:"map_variables"`
	Status       string               `json:"status"`
}
