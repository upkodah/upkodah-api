package room

import (
	"github.com/upkodah/upkodah-api/pkg/util"
)

/*
부동산 매물 정보를 가지고 있는 struct
Room은 지도상에 보여지는 정보를 담고있고,
Info는 실제 상세 페이지에 있는 정보를 담고 있다.

DID는 다방의 매물 번호로 이를 통해 기존에 이미 등록한 매물인지 확인 후 변경, 삭제를 처리한다.
Facilities는 facility code 들을 string 형태로 가지고 있다.
*/

type Room struct {
	util.DBModel `gorm:"embedded"`
	GridID       int              `json:"gridId" gorm:"not null;index"`
	Lat          float64          `json:"latitude" gorm:"not null"`
	Lon          float64          `json:"longitude" gorm:"not null"`
	EstType      int              `json:"estateType" gorm:"not null;index"`
	TrdType      int              `json:"tradeType" gorm:"not null;index"`
	Title        string           `json:"title" gorm:"not null; size:255"`
	Price        int              `json:"price" gorm:"not null"`
	Deposit      int              `json:"deposit" gorm:"not null"`
	FloorStr     string           `json:"floorStr" gorm:"size:255"`
	RealSize     float64          `json:"realSize" gorm:"not null"`
	RoughSize    float64          `json:"roughSize" gorm:"not null"`
	Facilities   string           `json:"facilities" gorm:"not null; size:255"`
	ImgUrls      util.UrlArrayStr `json:"imgUrls" gorm:"size:65535"`
	DID          string           `json:"d_id" gorm:"unique_index; not null; size:255"`
}

type Info struct {
	util.DBModel `gorm:"embedded"`
	RoomID       int              `json:"roomId" gorm:"not null;index"`
	GridID       uint             `json:"gridId" gorm:"not null;index"`
	Lat          float64          `json:"latitude" gorm:"not null"`
	Lon          float64          `json:"longitude" gorm:"not null"`
	EstType      int              `json:"estateType" gorm:"not null"`
	TrdType      int              `json:"tradeType" gorm:"not null"`
	Title        string           `json:"title" gorm:"not null; size:1000"`
	Price        int              `json:"price" gorm:"not null"`
	Deposit      int              `json:"deposit" gorm:"not null"`
	FloorStr     string           `json:"floorStr" gorm:"size:255"`
	RealSize     float64          `json:"realSize" gorm:"not null"`
	RoughSize    float64          `json:"roughSize" gorm:"not null"`
	Facilities   string           `json:"facilities" gorm:"not null; size:255"`
	ImgUrls      util.UrlArrayStr `json:"imgUrls" gorm:"size:65535"`
	Addr         string           `json:"address" gorm:"size:255"`
	RoadAddr     string           `json:"roadAddress" gorm:"size:255"`
	Descrip      string           `json:"describe" gorm:"size:65535"`
	IsAnimal     bool             `json:"isAnimal"`
	IsBalcony    bool             `json:"isBalcony"`
	IsElevator   bool             `json:"isElevator"`
	BathNum      int              `json:"bathNum"`
	BedNum       int              `json:"bedNum"`
	Direct       string           `json:"direct" gorm:"size:255"`
	HeatType     string           `json:"heatType" gorm:"size:255"`
	TotalCostStr string           `json:"totalCost" gorm:"size:255"`
	PhoneNumStr  string           `json:"phoneNum" gorm:"size:100"`
}
