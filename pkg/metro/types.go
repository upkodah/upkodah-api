package metro

import "github.com/upkodah/upkodah-api/pkg/util"

type Metro struct {
	util.DBModel
	Line string `json:"line" gorm:"size:255; not null"`
	Name string `json:"name" gorm:"size:255; not null"`
	Time int    `json:"time"`
}

type ResBody struct {
	Documents []Station `json:"documents"`
}

type Station struct {
	ID        string `json:"id"`
	PlaceName string `json:"place_name"`
	Distance  string `json:"distance"`
	X         string `json:"x"`
	Y         string `json:"y"`
}
