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
	count      uint
	Regions    []Region `gorm:"many2many:search_regions;"`
	Facilities facility.Facilities
}

type Region struct {
	gorm.Model
	Lft   int
	Rgt   int
	Top   int
	Btm   int
	Rooms []room.RoomShort
}
