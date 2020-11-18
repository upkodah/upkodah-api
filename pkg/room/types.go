package room

import (
	"github.com/jinzhu/gorm"
	"github.com/upkodah/upkodah-api/pkg/util"
)

type Room struct {
	gorm.Model
	GridID     uint             `json:"gridId" gorm:"index"`
	Lct        string           `json:"location" gorm:"index"`
	EstType    uint             `json:"estateType" gorm:"not null;index"`
	TrdType    string           `json:"tradeType" gorm:"not null;index"`
	Title      string           `json:"title"`
	Price      uint             `json:"price"`
	Deposit    uint             `json:"deposit"`
	FloorStr   string           `json:"floorStr"`
	RealSize   float64          `json:"realSize"`
	RoughSize  float64          `json:"roughSize"`
	Facilities util.ArrayStr    `json:"facilities"`
	ImgUrls    util.UrlArrayStr `json:"imgUrls"`
}

type Info struct {
	gorm.Model
	RoomID     uint             `json:"roomId" gorm:"index"`
	GridID     uint             `json:"gridId" gorm:"index"`
	Lct        string           `json:"location"`
	EstType    uint             `json:"estateType"`
	TrdType    string           `json:"tradeType"`
	Title      string           `json:"title"`
	Price      uint             `json:"price"`
	Deposit    uint             `json:"deposit"`
	FloorStr   string           `json:"floorStr"`
	RealSize   float64          `json:"realSize"`
	RoughSize  float64          `json:"roughSize"`
	Facilities util.ArrayStr    `json:"facilities"`
	ImgUrls    util.UrlArrayStr `json:"imgUrls"`
	Addr       string           `json:"address"`
	RoadAddr   string           `json:"roadAddress"`
	Describe   string           `json:"describe"`
	IsAnimal   bool             `json:"isAnimal"`
	IsBalcony  bool             `json:"isBalcony"`
	IsElevator bool             `json:"isElevator"`
	BathNum    uint             `json:"bathNum"`
	BedNum     uint             `json:"bedNum"`
	Direct     string           `json:"direct"`
	HeatType   string           `json:"heatType"`
	TotalCost  string           `json:"totalCost"`
	PhoneNum   string           `json:"phoneNum"`
}
