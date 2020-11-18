package facility

import (
	"github.com/jinzhu/gorm"
	"github.com/upkodah/upkodah-api/pkg/util"
)

type Facility struct {
	gorm.Model
	Name string   `json:"name"`
	Icon util.Url `json:"icon"`
}
