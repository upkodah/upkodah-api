package room

import (
	"github.com/upkodah/upkodah-api/pkg/util"
)

type Room struct {
	util.DBModel `gorm:"embedded"`
	GridID       uint             `json:"gridId" gorm:"not null;index"`
	Lat          float64          `json:"latitude" gorm:"not null"`
	Lon          float64          `json:"longitude" gorm:"not null"`
	EstType      uint             `json:"estateType" gorm:"not null;index"`
	TrdType      uint             `json:"tradeType" gorm:"not null;index"`
	Title        string           `json:"title" gorm:"not null; size:255"`
	Price        uint             `json:"price" gorm:"not null"`
	Deposit      uint             `json:"deposit" gorm:"not null"`
	FloorStr     string           `json:"floorStr" gorm:"size:255"`
	RealSize     float64          `json:"realSize" gorm:"not null"`
	RoughSize    float64          `json:"roughSize" gorm:"not null"`
	Facilities   util.ArrayStr    `json:"facilities" gorm:"not null; size:255"`
	ImgUrls      util.UrlArrayStr `json:"imgUrls" gorm:"size:65535"`
	DID          string           `json:"d_id" gorm:"unique_index; not null; size:255"`
}

type Info struct {
	util.DBModel `gorm:"embedded"`
	RoomID       uint             `json:"roomId" gorm:"not null;index"`
	GridID       uint             `json:"gridId" gorm:"not null;index"`
	Lat          float64          `json:"latitude" gorm:"not null"`
	Lon          float64          `json:"longitude" gorm:"not null"`
	EstType      uint             `json:"estateType" gorm:"not null"`
	TrdType      uint             `json:"tradeType" gorm:"not null"`
	Title        string           `json:"title" gorm:"not null; size:1000"`
	Price        uint             `json:"price" gorm:"not null"`
	Deposit      uint             `json:"deposit" gorm:"not null"`
	FloorStr     string           `json:"floorStr" gorm:"size:255"`
	RealSize     float64          `json:"realSize" gorm:"not null"`
	RoughSize    float64          `json:"roughSize" gorm:"not null"`
	Facilities   util.ArrayStr    `json:"facilities" gorm:"not null; size:255"`
	ImgUrls      util.UrlArrayStr `json:"imgUrls" gorm:"size:65535"`
	Addr         string           `json:"address" gorm:"size:255"`
	RoadAddr     string           `json:"roadAddress" gorm:"size:255"`
	Descrip      string           `json:"describe" gorm:"size:65535"`
	IsAnimal     bool             `json:"isAnimal"`
	IsBalcony    bool             `json:"isBalcony"`
	IsElevator   bool             `json:"isElevator"`
	BathNum      uint             `json:"bathNum"`
	BedNum       uint             `json:"bedNum"`
	Direct       string           `json:"direct" gorm:"size:255"`
	HeatType     string           `json:"heatType" gorm:"size:255"`
	TotalCostStr string           `json:"totalCost" gorm:"size:255"`
	PhoneNumStr  string           `json:"phoneNum" gorm:"size:100"`
}
