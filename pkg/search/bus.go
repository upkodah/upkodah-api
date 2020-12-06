package search

import (
	"errors"
	"github.com/upkodah/upkodah-api/pkg/bus"
)

func SearchByBus(startX string, startY string, time int, keyword string) ([]Item, error) {
	items := make([]Item, 0, 0)

	var stations []bus.Station
	stations, err := bus.GetNearbyStation(startX, startY)
	if err != nil {
		return nil, errors.New("SearchByBus/" + err.Error())
	}

	routeMap := make(map[string]bool)

	cnt := 0

	for _, s := range stations {
		routes, err := bus.GetRoutesByStation(s.ArsID)
		if err != nil {
			return nil, errors.New("SearchByBus/" + err.Error())
		}
		for _, r := range routes {
			if routeMap[r.BusRouteID] {
				cnt++
				continue
			}

			routeMap[r.BusRouteID] = true

			stationInfos, err := bus.GetStationsByRoute(r.BusRouteID)
			if err != nil {
				return nil, errors.New("SearchByBus/" + err.Error())
			}

			pInfos, err := findDest(s.ArsID, stationInfos, time)
			if err != nil {
				return nil, errors.New("SearchByBus/" + err.Error())
			}

			pathInfos2Items(&items, keyword, pInfos)
		}
	}
	return items, nil
}

func pathInfos2Items(items *[]Item, keyword string, pathInfos []bus.PathInfo) {
	for _, pi := range pathInfos {
		*items = append(*items, Item{
			Keyword:   keyword,
			TransType: 0,
			Name:      pi.RouteNm,
			Station:   pi.TName,
			GridID:    calcGridIDFromStr(pi.TY, pi.TX),
			Time:      pi.Time,
		})
	}
}

func findDest(arsID string, stationInfos []bus.StationInfo, time int) ([]bus.PathInfo, error) {

	trgStations := make([]bus.PathInfo, 0, 10)
	idxDist := int(float32(time-10) / 1.5) // time - 10 <= targetTime <= time 으로 하기 위해

	startIdx := 0
	var direct string
	// find start points
	for i := 0; i < len(stationInfos); i++ {
		if stationInfos[i].ArsID == arsID {
			startIdx = i
			direct = stationInfos[i].Direction
			break
		}
	}

	leftIdx := startIdx - idxDist
	for leftIdx >= 0 && stationInfos[leftIdx].Direction == direct {
		p, t, err := bus.GetPathNTime(stationInfos[startIdx], stationInfos[leftIdx])
		if err != nil {
			return nil, errors.New("findDest/" + err.Error())
		}
		if t == 0 {
			return nil, errors.New("findDest/time error")
		}
		if t < time-10 {
			leftIdx--
			continue
		}
		if t > time {
			break
		}
		p.Time = t
		trgStations = append(trgStations, p)
		leftIdx--
	}

	rightIdx := startIdx + idxDist
	for rightIdx < len(stationInfos) && stationInfos[rightIdx].Direction == direct {
		p, t, err := bus.GetPathNTime(stationInfos[startIdx], stationInfos[rightIdx])
		if err != nil {
			return nil, errors.New("findDest/" + err.Error())
		}
		if t == 0 {
			return nil, errors.New("findDest/time error")
		}
		if t < time-10 {
			rightIdx++
			continue
		}
		if t > time {
			break
		}
		p.Time = t
		trgStations = append(trgStations, p)
		rightIdx++
	}

	return trgStations, nil
}
