package baidu

import (
	"github.com/xlvector/gomap/data"
	"github.com/xlvector/gomap/util"
	"sort"
)

type NearyByResultList []data.NearyByResult

func (self NearyByResultList) Len() int {
	return len(self)
}

func (self NearyByResultList) Less(i, j int) bool {
	return self[i].Distance < self[j].Distance
}

func (self NearyByResultList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

type NearByAPI struct {
	placeAPI *PlaceAPI
}

func NewNearByAPI() *NearByAPI {
	api := NearByAPI{}
	api.placeAPI = NewPlaceAPI()
	return &api
}

func (self *NearByAPI) NearBy(address, query, region string, radius int) *data.NearByResp {
	locations := self.placeAPI.SearchInRegion(address, region)
	if locations == nil || len(locations.Results) == 0 {
		return nil
	}
	bestMatchLocation := locations.Results[0].Location

	places := self.placeAPI.SearchInCicle(query, bestMatchLocation, radius)
	if places == nil {
		return nil
	}

	ret := data.NearByResp{
		Status:  places.Status,
		Message: places.Message,
		Results: []data.NearyByResult{},
		Center: data.PlaceResult{
			Name: locations.Results[0].Name,
			Location: data.Location{
				Lng: bestMatchLocation.Lng,
				Lat: bestMatchLocation.Lat,
			},
			Address: locations.Results[0].Address,
		},
	}

	for _, result := range places.Results {
		dret := data.NearyByResult{
			Name:        result.Name,
			Orientation: util.OrientationInChina(&data.Location{Lng: bestMatchLocation.Lng, Lat: bestMatchLocation.Lat}, &data.Location{Lng: result.Location.Lng, Lat: result.Location.Lat}),
			Location: data.Location{
				Lng: result.Location.Lng,
				Lat: result.Location.Lat,
			},
			Address:  result.Address,
			Distance: float64(result.DetailInfo.Distance),
		}
		ret.Results = append(ret.Results, dret)
	}
	sort.Sort((NearyByResultList)(ret.Results))
	return &ret
}
