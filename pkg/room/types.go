package room

import (
	"github.com/jinzhu/gorm"
	"github.com/upkodah/upkodah-api/pkg/facility"
)

// stringify Image Array
type Images string

type Room struct {
	gorm.Model
	GridID     uint
	InfoID     uint
	RmType     string
	TrdType    string
	Lat        float64
	Lon        float64
	Addr       string
	Name       string
	MonPrice   uint
	DpsPrice   uint
	MngFee     uint
	RmSize     float64
	Facilities []facility.Facility `gorm:"many2many:room_facilities;"`
}

type Info struct {
	gorm.Model
	RoomID uint
	Desc   string
	Images Images
	FlrNum uint
	RmCnt  uint
	PhnNum string
}
