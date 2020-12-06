package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upkodah/upkodah-api/pkg/room"
	"github.com/upkodah/upkodah-api/pkg/search"
	"log"
	"strconv"
)

func getRooms(c *gin.Context) {

	qs := getDefaultQueries(c, map[string]string{
		"facilities":  "",
		"trade_type":  "0",
		"estate_type": "0",
		"latitude":    "",
		"longitude":   "",
		"time":        "50",
		"price":       room.MaxPrice,
		"deposit":     room.MaxPrice,
	})

	s := search.Search{}

	if err := s.SetConditions(qs); err != nil {
		log.Printf("Error in getRooms/%s\n", err)
		resErr404InvalidParam(c)
		return
	}

	//fmt.Println(util.ObjectToString(s))

	if err := s.SetAll(); err != nil {
		// New Search
		x := strconv.FormatFloat(s.Lon, 'f', -1, 64)
		y := strconv.FormatFloat(s.Lat, 'f', -1, 64)

		// Search By Bus
		items, err := search.SearchByBus(x, y, s.Time, s.Keyword)
		if err != nil {
			log.Printf("Error in getRooms/%s\n", err)
			resErr500(c)
			return
		}
		s.Items = append(s.Items, items...)

		// Search By Metro
		items, err = search.SearchByMetro(x, y, s.Time, s.Keyword)
		if err != nil {
			log.Printf("Error in getRooms/%s\n", err)
			resErr500(c)
			return
		}
		s.Items = append(s.Items, items...)
	}

	s.Facilities = qs["facilities"]

	var rms []room.Room

	rms = room.GetRooms(s.GetGridIDs(), qs)
	//fmt.Println(len(rms))

	gridMap := room.GroupByGridID(rms, s.Facilities)

	resOK(c, map[string]interface{}{
		"grids":       gridMap,
		"longitude":   s.Lon,
		"latitude":    s.Lat,
		"estate_type": s.EstType,
		"trade_type":  s.TrdType,
		"facilities":  s.Facilities,
	})
}

func getRoomsBySearchID(c *gin.Context) {
	searchID := c.Param("search_id")

	s := search.Search{}
	var err error

	s.ID, err = strconv.Atoi(searchID)
	if err != nil {
		log.Printf("Error in getRoomsBySearchID/%s\n", err)
		resErr404InvalidParam(c)
		return
	}

	if err := s.SetAll(); err != nil {
		log.Printf("Error in getRoomsBySearchID/%s\n", err)
		resErr404InvalidParam(c)
		return
	}

	var rms []room.Room

	rms = room.GetRooms(s.GetGridIDs(), map[string]string{
		"estate_type": strconv.Itoa(s.EstType),
		"trade_type":  strconv.Itoa(s.TrdType),
		"price":       strconv.Itoa(s.Price),
		"deposit":     strconv.Itoa(s.Deposit),
	})

	//fmt.Println(len(rms))

	gridMap := room.GroupByGridID(rms, s.Facilities)

	resOK(c, map[string]interface{}{
		"grids":       gridMap,
		"longitude":   s.Lon,
		"latitude":    s.Lat,
		"estate_type": s.EstType,
		"trade_type":  s.TrdType,
		"facilities":  s.Facilities,
	})

}

func getInfo(c *gin.Context) {
	info := room.Info{}
	idStr := c.Param("id")
	if idStr == "" {
		resErr404InvalidParam(c)
		return
	}

	roomID, err := strconv.Atoi(idStr)
	if err != nil {
		resErr404InvalidParam(c)
		return
	}

	info.RoomID = roomID

	//fmt.Println(util.ObjectToString(info))

	if err := info.Get(); err != nil {
		log.Printf("Error in getInfo/%s\n", err)
		resErr404InvalidParam(c)
		return
	}

	resOK(c, info)
}
