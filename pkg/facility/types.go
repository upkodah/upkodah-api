package facility

import (
	"github.com/upkodah/upkodah-api/pkg/util"
)

type Facility struct {
	util.DBModel `gorm:"embedded"`
	Name         string   `json:"name" gorm:"unique"`
	Code         string   `json:"code" gorm:"unique_index"`
	Type         uint     `json:"type" gorm:"not null"`
	Icon         util.Url `json:"icon"`
}
