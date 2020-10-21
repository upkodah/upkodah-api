package room

import (
	"github.com/jinzhu/gorm"
	"github.com/upkodah/upkodah-api/pkg/facility"
)

// stringify Image Array
type Images string

type Room struct {
	gorm.Model
	RegionID   uint
	Lat        float32
	Lon        float32
	Name       string
	RmType     string
	TrdType    string
	MonPrice   uint
	DpsPrice   uint
	RmSize     float32
	Desc       string
	MngFee     uint
	Images     Images
	FlrNum     uint
	RmCnt      uint
	Facilities []facility.Facility `gorm:"many2many:room_facilities;"`
	Addr       string
	PhnNum     string
}
