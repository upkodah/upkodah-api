package room

import (
	"github.com/jinzhu/gorm"
	"github.com/upkodah/upkodah-api/pkg/facility"
)

// stringify Image Array
type Images string

type RoomShort struct {
	gorm.Model
	RoomID    uint
	RegionID  uint
	Lat       float32
	Lon       float32
	Name      string
	RmType    string
	TrdType   string
	MonPrice  uint
	CharPrice uint
	RmSize    float32
}

type Room struct {
	gorm.Model
	Room       RoomShort
	Desc       string
	MngFee     uint
	Images     Images
	FlrNum     uint
	RmCount    uint
	Facilities facility.Facilities
	Addr       string
	PhNum      string
}
