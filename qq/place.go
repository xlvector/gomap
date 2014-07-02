package qq

import(
    "encoding/json"
    "log"
    "strconv"
    "github.com/xlvector/gomap/util"
    "net/url"
)

type Data struct {
    Id      string      `json:"id"`
    Title   string      `json:"title"`
    Category    string  `json:"category"`
    Type    int     `json:"type"`
    Location    Location    `json:"location"`
    Address     string      `json:"address"`
    Distance	float64		`json:"_distance"`
}

type PlaceAPIResp struct {
    Status      int     `json:"status"`
    Message     string      `json:"message"`
    Data        []Data  `json:"data"`
    Count       int `json:"count"`
}

func Decode(text string) *PlaceAPIResp {
    ret := PlaceAPIResp{}
    json.Unmarshal([]byte(text), &ret)
    return &ret
}

func buildEndpoint(api string, params map[string]string) string{
    endpoint := api
    for key, value := range params {
        endpoint += key + "=" + value + "&"
    }
    endpoint += "output=json&key=" + AK
    return endpoint
}

type PlaceAPI struct {
    dl *util.Downloader
}

func NewPlaceAPI() *PlaceAPI {
    api := PlaceAPI{}
    api.dl = util.NewDownloader()
    return &api
}

func (self *PlaceAPI) Search(params map[string]string) *PlaceAPIResp {
    endpoint := buildEndpoint("http://apis.map.qq.com/ws/place/v1/search?", params)
    log.Println(endpoint)
    resp, err := self.dl.Download(endpoint)
    if(err != nil){
        return nil
    }
    return Decode(resp)
}

func (self *PlaceAPI) SearchInRegion(query, region string) *PlaceAPIResp {
    return self.Search(map[string]string{
        "keyword": url.QueryEscape(query),
        "boundary": "region(" + url.QueryEscape(region) + ",0)",
        })
}

func (self *PlaceAPI) SearchInBounds(query, bounds string) *PlaceAPIResp {
    return self.Search(map[string]string{
        "keyword": url.QueryEscape(query),
        "boundary": "rectangle(" + bounds + ")",
        })
}

func (self *PlaceAPI) SearchInCicle(query string, location Location, radius int) *PlaceAPIResp {
    return self.Search(map[string]string{
        "keyword": url.QueryEscape(query),
        "boundary": "nearby(" + strconv.FormatFloat(location.Lat, 'f', 5, 64) + "," + strconv.FormatFloat(location.Lng, 'f', 5, 64) + "," + strconv.Itoa(radius) + ")",
        })
}

