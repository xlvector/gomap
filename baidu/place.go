package baidu

import (
	"encoding/json"
	"github.com/xlvector/gomap/data"
	"github.com/xlvector/gomap/util"
	"log"
	"net/url"
	"strconv"
	"time"
)

func AK() string {
	mod := time.Now().Unix() % 8
	if mod == 0 {
		return "F827baac9ff8d1e98178bcb0be60fc3b"
	} else if mod == 1 {
		return "693c4e0009c584eaafefdb5116e1b83e"
	} else if mod == 2 {
		return "F5f73aaf9942862674595224a4f486cb"
	} else if mod == 3 {
		return "CC4726bc48df6f635d4d69647b19ca02"
	} else if mod == 4 {
		return "458af52b49be366e290166aac63f25ea"
	} else if mod == 5 {
		return "D2ec390ea2301659e4a99313c584e094"
	} else if mod == 6 {
		return "D18ad00c81680c7ad7707af245b48c25"
	} else {
		return "AA8c70e83172c4eaf142b63c546c9bf7"
	}
}

type Location struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type DetailInfo struct {
	Tag              string `json:"tag"`
	DetailURL        string `json:"detail_url"`
	Type             string `json:"type"`
	Price            string `json:"price"`
	OverallRating    string `json:"overall_rating"`
	Distance         int    `json:"distance"`
	ShopHours        string `json:"shop_hours"`
	TasteRating      string `json:"taste_rating"`
	ServiceRating    string `json:"service_rating"`
	EnviromentRating string `json:"enviroment_rating"`
	FacilityRating   string `json:"facility_rating"`
	HygieneRating    string `json:"hygiene_rating"`
	TechnologyRating string `json:"technology_rating"`
	ImageNum         string `json:"image_num"`
	GrouponNum       int    `json:"groupon_num"`
	DiscountNum      int    `json:"discount_num"`
	CommentNum       string `json:"comment_num"`
	FavoriteNum      string `json:"favorite_num"`
	CheckinNum       string `json:"checkin_num"`
}

type Result struct {
	Name       string     `json:"name"`
	Location   Location   `json:"location"`
	Address    string     `json:"address"`
	Telephone  string     `json:"telephone"`
	Uid        string     `json:"uid"`
	DetailInfo DetailInfo `json:"detail_info"`
}

type PlaceAPIResp struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Results []Result `json:"results"`
}

func Decode(text string) *PlaceAPIResp {
	ret := PlaceAPIResp{}
	json.Unmarshal([]byte(text), &ret)
	return &ret
}

func buildEndpoint(api string, params map[string]string) string {
	endpoint := api
	hasScope := false
	hasAK := false
	for key, value := range params {
		endpoint += key + "=" + url.QueryEscape(value) + "&"
		if key == "scope" {
			hasScope = true
		}
		if key == "ak" {
			hasAK = true
		}
	}
	if !hasScope {
		endpoint += "scope=2&"
	}
	if !hasAK {
		endpoint += "ak=" + AK() + "&"
	}
	endpoint += "output=json"
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
	endpoint := buildEndpoint("http://api.map.baidu.com/place/v2/search?", params)
	log.Println(endpoint)
	for i := 0; i < 3; i++ {
		resp, err := self.dl.Download(endpoint)
		if err != nil {
			log.Println("Place Search", err)
			continue
		}
		return Decode(resp)
	}
	return nil
}

func (self *PlaceAPI) SearchInRegion(query, region string) *PlaceAPIResp {
	return self.Search(map[string]string{
		"query":  query,
		"region": region,
	})
}

func (self *PlaceAPI) SearchInBounds(query, bounds string) *PlaceAPIResp {
	return self.Search(map[string]string{
		"query":  query,
		"bounds": bounds,
	})
}

func (self *PlaceAPI) SearchInCicle(query string, location Location, radius int) *PlaceAPIResp {
	return self.Search(map[string]string{
		"query":    query,
		"location": strconv.FormatFloat(location.Lat, 'f', 5, 64) + "," + strconv.FormatFloat(location.Lng, 'f', 5, 64),
		"radius":   strconv.Itoa(radius),
	})
}

func (self *PlaceAPI) SearchPlace(address, region string) *data.PlaceResp {
	locations := self.SearchInRegion(address, region)
	if locations == nil || len(locations.Results) == 0 {
		return nil
	}
	bestMatchLocation := locations.Results[0].Location

	ret := data.PlaceResp{
		Status:  locations.Status,
		Message: locations.Message,
		Result: data.PlaceResult{
			Name: locations.Results[0].Name,
			Location: data.Location{
				Lng: bestMatchLocation.Lng,
				Lat: bestMatchLocation.Lat,
			},
			Address: locations.Results[0].Address,
		},
	}
	return &ret
}
