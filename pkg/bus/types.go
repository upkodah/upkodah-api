package bus

import "encoding/json"

type ResBody struct {
	ServiceResult ServiceResult `json:"ServiceResult"`
}

type ServiceResult struct {
	MsgBody MsgBody `json:"MsgBody"`
}

type MsgBody struct {
	ItemList json.RawMessage `json:"itemList"`
}

type Station struct {
	ArsID     string `json:"arsId" korean:"정류소고유번호"`
	Dist      string `json:"dist" korean:"거리"`
	GpsX      string `json:"gpsX" korean:"정류소 좌표X (WGS84)"`
	GpsY      string `json:"gpsY" korean:"정류소 좌표Y (WGS84)"`
	PosX      string `json:"posX" korean:"정류소 좌표X (GRS80)"`
	PosY      string `json:"posY" korean:"정류소 좌표Y (GRS80)"`
	StationID string `json:"stationId" korean:"정류소 ID"`
	StationNm string `json:"stationNm" korean:"정류소명"`
	StationTp string `json:"stationTp" korean:"정류소타입 (0:공용, 1:일반형 시내/농어촌버스, 2:좌석형 시내/농어촌버스, 3:직행좌석형 시내/농어촌버스, 4:일반형 시외버스, 5:좌석형 시외버스, 6:고속형 시외버스, 7:마을버스)"`
}

type Route struct {
	BusRouteID string `json:"busRouteId" korean:"노선ID"`
	BusRouteNm string `json:"busRouteNm" korean:"노선명"`
}

type StationInfo struct {
	BusRouteID   string `json:"busRouteId" korean:"노선 ID"`
	BusRouteNm   string `json:"busRouteNm" korean:"노선명"`
	Seq          string `json:"length" korean:"순번"`
	Section      string `json:"section" korean:"구간 ID"`
	Station      string `json:"station" korean:"정류소 ID"`
	StationNm    string `json:"stationNm" korean:"정류소 이름"`
	GPSX         string `json:"gpsX" korean:"X좌표 (WGS 84)"`
	GPSY         string `json:"gpsY" korean:"Y좌표 (WGS 84)"`
	Direction    string `json:"direction" korean:"진행방향"`
	FullSectDist string `json:"fullSectDist" korean:"정류소간 거리"`
	StationNo    string `json:"stationNo" korean:"정류소 고유번호"`
	RouteType    string `json:"routeType" korean:"노선 유형 (1:공항, 2:마을, 3:간선, 4:지선, 5:순환, 6:광역, 7:인천, 8:경기, 9:폐지, 0:공용)"`
	PosX         string `json:"posX" korean:"좌표X (GRS80)"`
	PosY         string `json:"posY" korean:"좌표Y (GRS80)"`
	SectSpd      string `json:"sectSpd" korean:"구간속도"`
	ArsID        string `json:"arsId" korean:"정류소 고유번호"`
}

type Path struct {
	Distance string          `json:"distance" korean:"거리"`
	Time     string          `json:"time" korean:"소요시간"`
	PathList json.RawMessage `json:"pathList" korean:"경로목록"`
}

type PathInfo struct {
	RouteID string `json:"routeId" korean:"노선ID"`
	RouteNm string `json:"routeNm" korean:"노선명"`
	FID     string `json:"fid" korean:"탑승지ID"`
	FName   string `json:"fname" korean:"탑승지명"`
	FX      string `json:"fx" korean:"탑승지 X좌표 (WGS84)"`
	FY      string `json:"fy" korean:"탑승지 Y좌표 (WGS84)"`
	TID     string `json:"tid" korean:"하차지ID"`
	TName   string `json:"tname" korean:"하차지명"`
	TX      string `json:"tx" korean:"하차지 X좌표 (WGS84)"`
	TY      string `json:"ty" korean:"하차지 Y좌표 (WGS84)"`
	Time    int    `json:"time" korean:"이동 시간"`
}
