package metro

import (
	"encoding/json"
	"errors"
	"github.com/upkodah/upkodah-api/pkg/db"
	"github.com/upkodah/upkodah-api/pkg/env"
	"github.com/upkodah/upkodah-api/pkg/util"
	"io/ioutil"
	"os"
)

func InitMetro() error {
	serviceKey = os.Getenv(env.MetroServiceKey)
	if serviceKey == "" {
		return errors.New("InitMetro/Invalid Env MetroServiceKey")
	}

	var metros []Metro
	if err := db.Conn.Find(&metros).Error; err != nil {
		return errors.New("InitMetro/" + err.Error())
	}

	metroMap = make(map[string][]Metro)

	for _, m := range metros {
		metroMap[m.Line] = append(metroMap[m.Line], m)
	}

	return nil
}

func GetMetrosFromMetroMap(line string) []Metro {
	return metroMap[line]
}

func GetNearbyStation(x string, y string) ([]Station, error) {
	queries := map[string]string{
		"page":                "1",
		"size":                "15",
		"sort":                "distance",
		"category_group_code": "SW8",
		"radius":              "300",
		"query":               "station",
		"x":                   x,
		"y":                   y,
	}

	resBytes, err := reqKakaoLocalAPI(KakaoKeywordApiUrl, queries)
	if err != nil {
		return nil, err
	}

	var stations []Station

	if err := json.Unmarshal(resBytes, &stations); err != nil {
		return nil, err
	}
	return stations, nil
}

func GetStationByName(name string) ([]Station, error) {
	queries := map[string]string{
		"page":                "1",
		"size":                "15",
		"sort":                "accuracy",
		"category_group_code": "SW8",
		"query":               name,
	}

	resBytes, err := reqKakaoLocalAPI(KakaoKeywordApiUrl, queries)
	if err != nil {
		return nil, err
	}

	var stations []Station

	if err := json.Unmarshal(resBytes, &stations); err != nil {
		return nil, err
	}
	return stations, nil

}

func reqKakaoLocalAPI(urlPath string, queries map[string]string) ([]byte, error) {
	header := map[string]string{
		"Authorization": serviceKey,
	}

	res, err := util.ReqGetWithHeader(urlPath, queries, header)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
