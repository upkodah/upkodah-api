package region

import (
	"github.com/jinzhu/gorm"
	"github.com/upkodah/upkodah-api/pkg/facility"
	"github.com/upkodah/upkodah-api/pkg/room"
)

type Search struct {
	gorm.Model
	Lat        float32
	Lon        float32
	time       uint
	cnt        uint
	Goos       []Goo               `gorm:"many2many:search_regions;"`
	Facilities []facility.Facility `gorm:"many2many:search_facilities;"`
}

type Goo struct {
	gorm.Model
	Name  string
	cnt   uint
	Dongs []Dong
}

type Dong struct {
	gorm.Model
	Name  string
	cnt   uint
	GooID uint
	Grids []Grid
}

type Grid struct {
	gorm.Model
	Lat    float32
	Lon    float32
	cnt    uint
	DongID uint
	Rooms  []room.Room
}
