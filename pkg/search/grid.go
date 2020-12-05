package search

import (
	"strconv"
)

const (
	BoundLeft = 124.17972222222222
	BoundLow  = 32.12277777777778
	GridSize  = 0.0036
)

func getGridID(latitude string, longitude string) (int, error) {
	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return 0, err
	}
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return 0, err
	}

	lon = lon - BoundLeft
	lat = lat - BoundLow

	gridId := (int(lon/GridSize) * 10000) + int(lat/GridSize)

	return gridId, nil
}
