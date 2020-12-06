package search

import (
	"github.com/upkodah/upkodah-api/pkg/util"
)

type Search struct {
	util.DBModel
	Keyword    string  `json:"keyword" gorm:"index"`
	Lat        float64 `json:"latitude" gorm:"not null; index"`
	Lon        float64 `json:"longitude" gorm:"not null; index"`
	GridID     int     `json:"gridId" gorm:"not null"`
	EstType    int     `json:"estateType" gorm:"default:0;not null;index"`
	TrdType    int     `json:"tradeType" gorm:"default:0;not null;index"`
	Price      int     `json:"price" gorm:"not null; default:10000000"`
	Deposit    int     `json:"deposit" gorm:"not null; default:10000000"`
	Time       int     `json:"time" gorm:"not null; index; default:60"`
	Facilities string  `json:"facilities"`
	Items      []Item  `json:"items"`
}

type Item struct {
	util.DBModel
	Keyword   string `json:"keyword"`
	TransType int    `json:"transType"`
	Name      string `json:"name"` // metro : 호선, bus : 노선
	Station   string `json:"station"`
	GridID    int    `json:"gridId" gorm:"not null"`
	SearchID  int    `json:"searchId" gorm:"not null; index"`
	Time      int    `json:"time" gorm:"not null"`
}
