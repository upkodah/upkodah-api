package region

import (
	"github.com/jinzhu/gorm"
	"github.com/upkodah/upkodah-api/pkg/room"
)

type Search struct {
	gorm.Model
	Lct        string `json:"location"`
	Cnt        uint   `json:"count"`
	Goos       []Goo  `json:"goos"`
	Facilities string `json:"facilities"`
}

type Goo struct {
	gorm.Model
	Name  string `json:"name"`
	Cnt   uint   `json:"count"`
	Dongs []Dong `json:"dongs"`
}

type Dong struct {
	gorm.Model
	Name  string `json:"name"`
	Cnt   uint   `json:"count"`
	GooID uint   `json:"gooId"`
	Grids []Grid `json:"grids"`
}

type Grid struct {
	gorm.Model
	Lct    string      `json:"location"`
	Cnt    uint        `json:"count"`
	DongID uint        `json:"dongId"`
	Rooms  []room.Room `json:"rooms"`
}
