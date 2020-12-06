package search

import (
	"errors"
	"github.com/upkodah/upkodah-api/pkg/db"
	"strconv"
)

func (s *Search) SetConditions(m map[string]string) error {
	var err error

	if s.TrdType, err = strconv.Atoi(m["trade_type"]); err != nil {
		return errors.New("SetConditions/" + err.Error())
	}
	if s.EstType, err = strconv.Atoi(m["estate_type"]); err != nil {
		return errors.New("SetConditions/" + err.Error())
	}
	if s.Lat, err = strconv.ParseFloat(m["latitude"], 64); err != nil {
		return errors.New("SetConditions/" + err.Error())
	}
	if s.Lon, err = strconv.ParseFloat(m["longitude"], 64); err != nil {
		return errors.New("SetConditions/" + err.Error())
	}
	if s.Time, err = strconv.Atoi(m["time"]); err != nil {
		return errors.New("SetConditions/" + err.Error())
	}
	if s.Price, err = strconv.Atoi(m["price"]); err != nil {
		return errors.New("SetConditions/" + err.Error())
	}
	if s.Deposit, err = strconv.Atoi(m["deposit"]); err != nil {
		return errors.New("SetConditions/" + err.Error())
	}
	s.GridID = calcGridID(s.Lat, s.Lon)
	return nil
}

func (s *Search) SetAll() error {

	var items []Item
	if s.ID != 0 {
		err := db.Conn.First(s).Error
		if err != nil {
			return errors.New("SetAll/" + err.Error())
		}
	} else if s.Lon != 0 && s.Lat != 0 {
		if err := db.Conn.Where("trd_type=? AND est_type=? AND grid_id AND time=?", s.TrdType, s.EstType, s.GridID, s.Time).First(s).Error; err != nil {
			return errors.New("SetAll/" + err.Error())
		}
	} else {
		return errors.New("SetAll/Invalid Condition")
	}

	if err := db.Conn.Where("search_id=?", s.ID).Find(&items).Error; err != nil {
		return errors.New("SetAll/" + err.Error())
	}

	s.Items = items
	return nil
}

func (s Search) GetGridIDs() []int {
	gIDs := make([]int, 0, len(s.Items))
	for _, i := range s.Items {
		gIDs = append(gIDs, i.GridID)
	}
	return gIDs
}
