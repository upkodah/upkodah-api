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
	TrdType      string           `json:"tradeType" gorm:"not null;index"`
	Title        string           `json:"title" gorm:"not null"`
	Price        uint             `json:"price" gorm:"not null"`
	Deposit      uint             `json:"deposit" gorm:"not null"`
	FloorStr     string           `json:"floorStr"`
	RealSize     float64          `json:"realSize" gorm:"not null"`
	RoughSize    float64          `json:"roughSize" gorm:"not null"`
	Facilities   util.ArrayStr    `json:"facilities" gorm:"not null"`
	ImgUrls      util.UrlArrayStr `json:"imgUrls"`
	DID          string           `json:"d_id" gorm:"unique_index; not null"`
}

type Info struct {
	util.DBModel `gorm:"embedded"`
	RoomID       uint             `json:"roomId" gorm:"not null;index"`
	GridID       uint             `json:"gridId" gorm:"not null;index"`
	Lat          float64          `json:"latitude" gorm:"not null"`
	Lon          float64          `json:"longitude" gorm:"not null"`
	EstType      uint             `json:"estateType" gorm:"not null"`
	TrdType      string           `json:"tradeType" gorm:"not null"`
	Title        string           `json:"title" gorm:"not null"`
	Price        uint             `json:"price" gorm:"not null"`
	Deposit      uint             `json:"deposit" gorm:"not null"`
	FloorStr     string           `json:"floorStr"`
	RealSize     float64          `json:"realSize" gorm:"not null"`
	RoughSize    float64          `json:"roughSize" gorm:"not null"`
	Facilities   util.ArrayStr    `json:"facilities" gorm:"not null"`
	ImgUrls      util.UrlArrayStr `json:"imgUrls"`
	Addr         string           `json:"address"`
	RoadAddr     string           `json:"roadAddress"`
	Descrip      string           `json:"describe"`
	IsAnimal     bool             `json:"isAnimal"`
	IsBalcony    bool             `json:"isBalcony"`
	IsElevator   bool             `json:"isElevator"`
	BathNum      uint             `json:"bathNum"`
	BedNum       uint             `json:"bedNum"`
	Direct       string           `json:"direct"`
	HeatType     string           `json:"heatType"`
	TotalCost    string           `json:"totalCost"`
	PhoneNum     string           `json:"phoneNum"`
}
