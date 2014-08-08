package baidu

import (
	"encoding/json"
	"github.com/xlvector/gomap/data"
	"github.com/xlvector/gomap/util"
	"log"
	"net/url"
)

type DirectionResp struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Type    int             `json:"type"`
	Result  DirectionResult `json:"result"`
}

type DirectionResult struct {
	Routes      []DirectionRoute `json:"routes"`
	Origin      OriginPt         `json:"origin"`
	Destination DestinationPt    `json:"destination"`
}

type OriginPt struct {
	OriginPt Location `json:"originPt"`
}

type DestinationPt struct {
	DestinationPt Location `json:"destinationPt"`
}

type DirectionRoute struct {
	Scheme []DirectionScheme `json:"scheme"`
}

type DirectionScheme struct {
	Distance            int               `json:"distance"`
	Duration            int               `json:"duration"`
	OriginLocation      Location          `json:"originLocation"`
	DestinationLocation Location          `json:"destinationLocation"`
	Steps               [][]DirectionStep `json:"steps"`
}

type DirectionStep struct {
	Distance                int      `json:"distance"`
	Duration                int      `json:"duration"`
	Path                    string   `json:"path"`
	Type                    int      `json:"type"`
	Vehicle                 Vehicle  `json:"vehicle"`
	StepOriginLocation      Location `json:"stepOriginLocation"`
	StepDestinationLocation Location `json:"stepDestinationLocation"`
	StepInstruction         string   `json:"stepInstruction"`
}

type Vehicle struct {
	EndName    string  `json:"end_name"`
	EndTime    string  `json:"end_time"`
	EndUid     string  `json:"end_uid"`
	Name       string  `json:"name"`
	StartName  string  `json:"start_name"`
	StartTime  string  `json:"start_time"`
	StartUid   string  `json:"start_uid"`
	StopNum    int     `json:"stop_num"`
	TotalPrice float64 `json:"total_price"`
	Uid        string  `json:"uid"`
	ZonePrice  float64 `json:"zone_price"`
}

type DirectionAPI struct {
	dl *util.Downloader
}

func NewDirectionAPI() *DirectionAPI {
	api := DirectionAPI{}
	api.dl = util.NewDownloader()
	return &api
}

func (self *DirectionAPI) Transit(origin, destination, region string) *DirectionResp {
	endpoint := "http://api.map.baidu.com/direction/v1?mode=transit&origin=" + url.QueryEscape(origin) + "&destination=" + url.QueryEscape(destination) + "&region=" + url.QueryEscape(region) + "&output=json&ak=" + AK()
	log.Println(endpoint)
	resp, err := self.dl.Download(endpoint)
	if err != nil {
		log.Println(err)
		return nil
	}
	ret := DirectionResp{}
	json.Unmarshal([]byte(resp), &ret)
	return &ret
}

func (self *DirectionAPI) TransitSequence(origin, destination, region string) *data.DirectionResp {
	baiduRet := self.Transit(origin, destination, region)
	if baiduRet == nil {
		return nil
	}
	ret := data.DirectionResp{
		Status:  baiduRet.Status,
		Message: baiduRet.Message,
		Result: data.DirectionResult{
			Routes: []data.DirectionRoute{},
		},
	}
	for _, route := range baiduRet.Result.Routes {
		for _, scheme := range route.Scheme {
			seq := []string{}
			prev := ""
			for _, steps := range scheme.Steps {
				if steps[0].Vehicle.StartName != "" && steps[0].Vehicle.StartName != prev {
					seq = append(seq, steps[0].Vehicle.StartName)
				}
				if steps[0].Vehicle.Name != "" {
					seq = append(seq, steps[0].Vehicle.Name)
				}
				if steps[0].Vehicle.EndName != "" {
					seq = append(seq, steps[0].Vehicle.EndName)
					prev = steps[0].Vehicle.EndName
				}
			}

			ret.Result.Routes = append(ret.Result.Routes, data.DirectionRoute{Sequence: seq})
		}
	}
	return &ret
}
