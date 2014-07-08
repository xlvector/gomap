package qq

import (
	"github.com/xlvector/gomap/data"
)

type NearByAPI struct {
	geoCoderAPI *GeoCoderAPI
	placeAPI *PlaceAPI
}

func NewNearByAPI() *NearByAPI {
	api := NearByAPI{}
	api.geoCoderAPI = NewGeoCoderAPI()
	api.placeAPI = NewPlaceAPI()
	return &api
}

func (self *NearByAPI) NearBy(address, query, region string, radius int) *data.NearByResp {
	locations := self.geoCoderAPI.GetLocationByAddress(address, region)
	if locations == nil {
		return nil
	}
	bestMatchLocation := locations.Result.Location

	places := self.placeAPI.SearchInCicle(query, bestMatchLocation, radius)
	if places == nil {
		return nil
	}

	ret := data.NearByResp{
		Status: places.Status,
		Message: places.Message,
		Results: []data.NearyByResult{},
	}

	for _, result := range places.Data {
		dret := data.NearyByResult{
			Name: result.Title,
			Location: data.Location{
				Lng: result.Location.Lng,
				Lat: result.Location.Lat,
			},
			Address: result.Address,
			Distance: result.Distance,
		}
		ret.Results = append(ret.Results, dret)
	}
	return &ret
}