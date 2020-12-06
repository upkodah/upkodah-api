package bus

const (
	ApiURL = "http://ws.bus.go.kr/api/rest"

	StationPathURL = ApiURL + "/stationinfo"
	RoutePathURL   = ApiURL + "/busRouteInfo"
	PathInfoURL    = ApiURL + "/pathinfo"

	GetStationByPosURL   = StationPathURL + "/getStationByPos"
	GetRouteByStationURL = StationPathURL + "/getRouteByStation"
	GetStationByRouteURL = RoutePathURL + "/getStaionByRoute"
	GetPathInfoByBusURL  = PathInfoURL + "/getPathInfoByBus"

	LON_MAX = 127.269311 // 서울특별시 경도 상한 X
	LON_MIN = 126.734086 // 서울특별시 경도 하한
	LAT_MAX = 37.715133  // 서울특별시 위도 상한 Y
	LAT_MIN = 37.413294  // 서울특별시 위도 하한
)

var serviceKey string
