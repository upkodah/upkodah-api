package bus

import (
	"encoding/json"
	"errors"
	"github.com/upkodah/upkodah-api/pkg/env"
	"github.com/upkodah/upkodah-api/pkg/util"
	"log"
	"os"
	"strconv"
)

//"127.0565884"
//"37.5838699"
func InitBus() error {
	serviceKey = os.Getenv(env.BusServiceKey)
	if serviceKey == "" {
		return errors.New("InitMetro/Invalid Env BusServiceKey")
	}
	return nil
}

func GetNearbyStation(x string, y string) ([]Station, error) {
	queries := map[string]string{
		"tmX":        x,
		"tmY":        y,
		"radius":     "300",
		"serviceKey": serviceKey,
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
		"serviceKey": serviceKey,
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
		"serviceKey": serviceKey,
		"busRouteId": routeID,
	}

	resBytes, err := reqSeoulBusAPI(GetStationByRouteURL, queries)
	if err != nil {
		log.Printf("Error in reqSeoulBusAPI : %s\n", err)
		return nil, err
	}

	var stationInfos []StationInfo

	if err := json.Unmarshal(resBytes, &stationInfos); err != nil {
		log.Printf("Error in reqSeoulBusAPI : %s\n", err)
		return nil, err
	}

	return stationInfos, nil
}

func GetPathNTime(startS StationInfo, endS StationInfo) (PathInfo, int, error) {

	paths, err := GetPathsToBus(startS.GPSX, startS.GPSY, endS.GPSX, endS.GPSY)

	if err != nil {
		return PathInfo{}, 0, err
	}

	for _, p := range paths {
		var pathInfo PathInfo
		if err := json.Unmarshal(p.PathList, &pathInfo); err != nil {
			continue
		}

		if pathInfo.RouteID == startS.BusRouteID {
			time, err := strconv.Atoi(p.Time)
			if err != nil {
				return PathInfo{}, 0, err
			}
			return pathInfo, time, nil
		}
	}

	return PathInfo{}, 0, nil
}

func GetPathsToBus(sX string, sY string, eX string, eY string) ([]Path, error) {

	if !checkXYValid(sX, sY) || !checkXYValid(eX, eY) {
		return nil, errors.New("GetPathsToBus/X and Y are out of range")
	}
	queries := map[string]string{
		"serviceKey": serviceKey,
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
		return nil, err
	}

	return paths, nil
}

func reqSeoulBusAPI(urlPath string, queries map[string]string) (json.RawMessage, error) {
	res, err := util.ReqGet(urlPath, queries)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Error in reqSeoulBusAPI : %s\n", err)
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

	return []byte(jsonStr), nil
}

func checkXYValid(x string, y string) bool {
	// check x
	fx, err := strconv.ParseFloat(x, 64)

	if err != nil || fx > LON_MAX || fx < LON_MIN {
		log.Printf("x : %f, LON_MAX : %f, LON_MIN : %f, err : %v", fx, LON_MAX, LON_MIN, err)
		return false
	}
	// check y
	fy, err := strconv.ParseFloat(y, 64)
	if err != nil || fy > LAT_MAX || fy < LAT_MIN {
		log.Printf("x : %f, LAT_MAX : %f, LAT_MIN : %f, err : %v", fy, LAT_MAX, LAT_MIN, err)
		return false
	}
	return true
}
