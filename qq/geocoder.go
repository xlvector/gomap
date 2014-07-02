package qq

import (
    "github.com/xlvector/gomap/util"
    "encoding/json"
    "log"
    "net/url"
)

const (
    AK = "EPRBZ-KCPRF-6E5JQ-JAQGL-NSDKH-RXFZ6"
)

type LocationResp struct {
    Status      int     `json:"status"`
    Message     string  `json:"message"`
    Result      LocationResult  `json:result`
}

type LocationResult struct {
    Similarity  float64     `json:"similarity"`
    Location    Location    `json:"location"`
    AddressComponents   AddressComponents   `json:"address_components"`
}

type Location struct {
    Lat float64 `json:"lat"`
    Lng float64 `json:"lng"`
}

type AddressComponents struct {
    Province    string  `json:"province"`
    City        string  `json:"city"`
    District    string  `json:"district"`
    Street      string  `json:"street"`
    StreetNumber    string  `json:"street_number"`
}

type GeoCoderAPI struct {
    dl *util.Downloader
}

func NewGeoCoderAPI() *GeoCoderAPI {
    api := GeoCoderAPI{}
    api.dl = util.NewDownloader()
    return &api
}

func (self *GeoCoderAPI) GetLocationByAddress(address, region string) *LocationResp{
    endpoint := "http://apis.map.qq.com/ws/geocoder/v1/?region=" + url.QueryEscape(region) + "&address=" + url.QueryEscape(address) + "&key=" + AK
    log.Println(endpoint)
    resp, err := self.dl.Download(endpoint)
    if(err != nil){
        return nil
    }
    ret := LocationResp{}
    json.Unmarshal([]byte(resp), &ret)
    return &ret
}