package search

import (
	"errors"
	"github.com/upkodah/upkodah-api/pkg/metro"
	"strings"
)

func SearchByMetro(startX string, startY string, time int, keyword string) ([]Item, error) {
	items := make([]Item, 0, 10)

	stations, err := metro.GetNearbyStation(startX, startY)
	if err != nil {
		return nil, errors.New("SearchByMetro/" + err.Error())
	}
	for _, s := range stations {
		name := strings.Split(s.PlaceName, " ")[0]
		line := strings.Split(s.PlaceName, " ")[1]
		metros := metro.GetMetrosFromMetroMap(line)

		var startIdx int
		for i, m := range metros {
			if m.Name == name {
				startIdx = i
				break
			}
		}

		// find from right
		t := 0
		for i := startIdx; i < len(metros); i++ {
			t += metros[i].Time
			if t < time-10 {
				continue
			} else if t <= time {
				if i+1 != len(metros) {
					gridId, err := getGridIDFromName(metros[i+1].Name)
					if err != nil {
						return nil, errors.New("SearchByMetro/" + err.Error())
					}
					items = append(items, Item{
						Keyword:   keyword,
						TransType: 1,
						Name:      metros[i+1].Line,
						Station:   metros[i+1].Name,
						GridID:    gridId,
					})
				}
			} else {
				break
			}
		}

		// find from left
		t = 0
		for i := startIdx - 1; i >= 0; i-- {
			t += metros[i].Time
			if t < time-10 {
				continue
			} else if t <= time {
				gridId, err := getGridIDFromName(metros[i].Name)
				if err != nil {
					return nil, errors.New("SearchByMetro/" + err.Error())
				}
				items = append(items, Item{
					Keyword:   keyword,
					TransType: 1,
					Name:      metros[i].Line,
					Station:   metros[i].Name,
					GridID:    gridId,
				})
			} else {
				break
			}
		}
	}

	return items, nil
}

func getGridIDFromName(name string) (int, error) {
	stations, err := metro.GetStationByName(name)
	if err != nil {
		return 0, errors.New("getGridIDFromName/" + err.Error())
	}

	if len(stations) == 0 {
		return 0, errors.New("getGridIDFromName/No Stations Exist")
	}

	return calcGridIDFromStr(stations[0].Y, stations[0].X), nil
}
