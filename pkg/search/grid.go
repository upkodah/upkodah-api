package search

import "strconv"

const (
	BoundLeft = 124.17972222222222
	BoundLow  = 32.12277777777778
	GridSize  = 0.0036
)

func calcGridIDFromStr(latitude string, longitude string) int {
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return 0
	}
	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return 0
	}
	return calcGridID(lat, lon)
}

func calcGridID(lat float64, lon float64) int {
	lon = lon - BoundLeft
	lat = lat - BoundLow

	gridId := (int(lon/GridSize) * 10000) + int(lat/GridSize)

	return gridId
}
