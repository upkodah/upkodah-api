# Upkodah Api
## Purpose
사용자의 직장으로부터 일정 시간이 걸리는 위치의 매물을 보여주는 부동산 애플리케이션의 API

API List
1. GET http://{ip_address}/ping

2. GET http://{ip_address}/setting
    - 앱의 초기 세팅을 가져와 주는 api 입니다. 현재는 편의시설 목록을 가져와주고 있습니다.
    - request  queries : 없음
    - 예시 : GET {ip_address}/setting

3. GET http://{ip_address}/v1/rooms
    - upkodah db에 있는 매물들을 원하는 latitude, longitude에서 time만큼 떨어진 매물을 가져와주는 api입니다.
    - request queries :
        - `facilities` : 편의시설의 keyword를 ,로 나눈 string 형태입니다. 총 9개의 keyword가 있습니다. (MT-대형마트, PS3-유치원, BK9-은행, CE7-카페, HP8-병원, PM9-약국, 경찰서-경찰서, 헬스장-헬스장, 세탁소-세탁소)
          `default` 값은 "" 으로 편의시설에 상관없이 매물을 가져옵니다. ex) "PS3, 세탁소"
        - `trade_type` : 매물 거래 종류입니다. 임대, 매매, 전세가 있으나 현재는 임대만 지원하고 있습니다. 무조건 0을 보내며 됩니다.
          `default` 값은 0 입니다. ex) 3
        - `estate_type` : 매물의 건물 종류입니다. 원룸, 아파트, 오피스텔이 있으나 현재는 원룸과 오피스텔만 지원하고 있습니다.
          `default` 값은 0 입니다. ex) 0 or 3 (원룸 or 오피스텔)
        - `latitude` : 위도 값을 의미합니다. float64 형태로 query string 에 담아주시면 됩니다. (37.413294 ~ 37.715133)
          `필수 요소` 입니다. 넣지 않을 시 404 Bad Request를 반환합니다. ex) 37.5759689663327
        - longitude : 경도 값을 의미합니다. float64 형태로 query string 에 담아주시면 됩니다. (126.734086 ~ 127.269311)
          `필수 요소` 입니다. 넣지 않을 시 404 Bad Request를 반환합니다. ex) 126.976861018866
        - time : latitude, longitude 로부터타 매물이 얼마나 떨어져 있는지를 나타내는 시간입니다. 분 단위이며 20, 30, 40, 50을 받습니다.
          default 값은 50 입니다. ex) 30
        - price : 사용자가 원하는 매물의 월세의 상향가를 의미합니다. 만원 단위이며 입력하지 않을 시 최대 월세로 지정됩니다.
          `default`값은 10000000 입니다. ex) 50 
        - deposit : 사용자가 원하는 매물 보증금의 상향가를 의미합니다. 만원 단위이며 입력하지 않을 시 최대 월세로 지정됩니다.
        `default`값은 10000000 입니다. ex) 50
    - 요청 예시 :
        <http://34.64.166.133:80/v1/rooms/?facilities=헬스클럽&estate_type=0&trade_type=0&latitude=37.5759689663327&longitude=126.976861018866&price=50&deposit=10000&time=30>

4. GET http://{ip_address}/v1/rooms/:search_id
    - 기존에 검색했던 search_id를 통해 매물을 가져오는 쿼리 입니다. 캐싱을 하기 때문에 자원을 소모하지 않고 더 빠르게 가져올 수 있습니다.
    - request queries : 없음
    - 요청 예시 :
        <http://34.64.166.133:80/v1/rooms/1>

5. GET http://{ip_address}/v1/room/:id/info
    - upkodah db의 매물에 해당하는 상세 정보를 가져와 주는 api입니다. room의 id를 parameter로 받아 사용합니다.
    - request queries : 없음
    - 요청 예시 :
        <http://34.64.166.133:80/v1/room/300/info>
