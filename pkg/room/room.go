package room

import (
	"errors"
	"github.com/upkodah/upkodah-api/pkg/db"
	"strings"
)

func GetRooms(gridIDs []int, qs map[string]string) []Room {
	var rms []Room
	q := "est_type=" + qs["estate_type"] + " AND trd_type=" + qs["trade_type"]

	if qs["price"] != MaxPrice {
		q = q + " AND price<=" + qs["price"]
	}
	if qs["deposit"] != MaxPrice {
		q = q + " AND deposit<=" + qs["deposit"]
	}

	if gridIDs == nil || len(gridIDs) == 0 {
		db.Conn.Where(q).Find(&rms)
	} else {
		db.Conn.Where(q+" AND grid_id IN (?)", gridIDs).Find(&rms)
	}
	return rms
}

func GroupByGridID(rms []Room, facilities string) map[int][]Room {
	gridMap := make(map[int][]Room)

	fcs := strings.Split(strings.ReplaceAll(facilities, " ", ""), ",")
	if len(facilities) != 0 {
		for _, r := range rms {
			fcMap := make(map[string]bool)
			tempFcs := strings.Split(strings.ReplaceAll(r.Facilities, " ", ""), ",")

			isContain := true

			for _, f := range tempFcs {
				fcMap[f] = true
			}

			for _, f := range fcs {
				if !fcMap[f] {
					isContain = false
					break
				}
			}

			if !isContain {
				continue
			}
			gridMap[r.GridID] = append(gridMap[r.GridID], r)
		}
	} else {
		for _, r := range rms {
			gridMap[r.GridID] = append(gridMap[r.GridID], r)
		}
	}
	return gridMap
}

func (i *Info) Get() error {
	if i.ID != 0 {
		return db.Conn.First(i).Error
	} else if i.RoomID != 0 {
		return db.Conn.Where("room_id=?", i.RoomID).First(i).Error
	} else {
		return errors.New("Get/Invalid Conditions")
	}
}
