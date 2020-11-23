package region

import (
	"github.com/upkodah/upkodah-api/pkg/room"
	"github.com/upkodah/upkodah-api/pkg/util"
	"time"
)

type Search struct {
	util.DBModel `gorm:"embedded"`
	Lat          float64 `json:"latitude" gorm:"not null"`
	Lon          float64 `json:"longitude" gorm:"not null"`
	Goos         []Goo   `json:"goos"`
	Facilities   string  `json:"facilities"`
}

type Goo struct {
	util.DBModel `gorm:"embedded"`
	Name         string  `json:"name"`
	Lat          float64 `json:"latitude" gorm:"not null"`
	Lon          float64 `json:"longitude" gorm:"not null"`
	Dongs        []Dong  `json:"dongs"`
}

type Dong struct {
	util.DBModel `gorm:"embedded"`
	Name         string  `json:"name"`
	Lat          float64 `json:"latitude" gorm:"not null"`
	Lon          float64 `json:"longitude" gorm:"not null"`
	GooID        uint    `json:"gooId"`
	Grids        []Grid  `json:"grids"`
}

type Grid struct {
	ID        string      `json:"id" gorm:"unique_index; not null"`
	Lat       float64     `json:"latitude" gorm:"not null"`
	Lon       float64     `json:"longitude" gorm:"not null"`
	CreatedAt time.Time   `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP;autoCreateTime;not null"`
	UpdatedAt time.Time   `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;not null"`
	DeletedAt *time.Time  `json:"deletedAt" sql:"index"`
	DongID    uint        `json:"dongId"`
	Rooms     []room.Room `json:"rooms"`
}
