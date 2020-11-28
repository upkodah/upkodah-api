package bus

import (
	"encoding/json"
	"fmt"
	"github.com/upkodah/upkodah-api/pkg/util"
	"strconv"
)

//"127.0565884"
//"37.5838699"
var count int

func GetStationsByPos(colX string, colY string) ([]Station, error) {
	queries := map[string]string{
		"tmX":        colX,
		"tmY":        colY,
		"radius":     "300",
		"serviceKey": ServiceKey,
	}
	resBytes, err := reqSeoulBusAPI(GetStationByPosURL, queries)
	if err != nil || resBytes == nil {
		return nil, err
	}

	var stations []Station

	if err := json.Unmarshal(resBytes, &stations); err != nil {
		return nil, err
	}

	return stations, nil
}

func GetRoutesByStation(arsID string) ([]Route, error) {
	queries := map[string]string{
		"serviceKey": ServiceKey,
		"arsId":      arsID,
	}

	resBytes, err := reqSeoulBusAPI(GetRouteByStationURL, queries)
	if err != nil {
		return nil, err
	}

	var routes []Route

	if err := json.Unmarshal(resBytes, &routes); err != nil {
		return nil, err
	}

	return routes, nil
}

func GetStationsByRoute(routeID string) ([]StationInfo, error) {
	queries := map[string]string{
		"serviceKey": ServiceKey,
		"busRouteId": routeID,
	}

	resBytes, err := reqSeoulBusAPI(GetStationByRouteURL, queries)
	if err != nil {
		fmt.Printf("Error in reqSeoulBusAPI : %s\n", err)
		return nil, err
	}

	var stationInfos []StationInfo

	if err := json.Unmarshal(resBytes, &stationInfos); err != nil {
		fmt.Printf("Error in reqSeoulBusAPI : %s\n", err)
		return nil, err
	}

	return stationInfos, nil
}

func GetPathsToBus(sX string, sY string, eX string, eY string) ([]Path, error) {

	if !checkXYValid(sX, sY) || !checkXYValid(eX, eY) {
		return nil, nil // TODO: make new Error type
	}
	queries := map[string]string{
		"serviceKey": ServiceKey,
		"startX":     sX,
		"startY":     sY,
		"endX":       eX,
		"endY":       eY,
	}

	resBytes, err := reqSeoulBusAPI(GetPathInfoByBusURL, queries)
	if err != nil {
		return nil, err
	}

	var paths []Path

	if err := json.Unmarshal(resBytes, &paths); err != nil {
		fmt.Println("error in fnlsnflsdn")
		return nil, err
	}

	return paths, nil
}

func reqSeoulBusAPI(urlPath string, queries map[string]string) (json.RawMessage, error) {
	count++
	res, err := util.ReqGet(urlPath, queries)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("Error in reqSeoulBusAPI : %s\n", err)
		return nil, err
	}

	buf, err := util.Xml2json(res.Body)
	if err != nil {
		return nil, err
	}

	body := ResBody{}

	if err := json.Unmarshal(buf.Bytes(), &body); err != nil {
		return nil, err
	}

	jsonStr := string(body.ServiceResult.MsgBody.ItemList)
	if jsonStr[0] != '[' {
		jsonStr = "[" + jsonStr + "]"
	}
	fmt.Printf("api count : %d\n", count)

	return []byte(jsonStr), nil
}

func checkXYValid(x string, y string) bool {
	// check x
	fx, err := strconv.ParseFloat(x, 64)
	if err != nil || fx > LON_MAX || fx < LON_MIN {
		fmt.Printf("x : %f, LON_MAX : %f, LON_MIN : %f, %b", fx, LON_MAX, LON_MIN)
		return false
	}
	// check y
	fy, err := strconv.ParseFloat(y, 64)
	if err != nil || fy > LAT_MAX || fy < LAT_MIN {
		fmt.Printf("x : %f, LAT_MAX : %f, LAT_MIN : %f, %b", fx, LAT_MAX, LAT_MIN)
		return false
	}
	return true
}
