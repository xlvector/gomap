package gomap

import (
	"encoding/json"
	"fmt"
	"github.com/xlvector/gomap/baidu"
	"github.com/xlvector/gomap/data"
	"net/http"
	"runtime/debug"
)

type PlaceService struct {
	api *baidu.PlaceAPI
}

func NewPlaceService() *PlaceService {
	service := PlaceService{
		api: baidu.NewPlaceAPI(),
	}
	return &service
}

func (self *PlaceService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if recvr := recover(); recvr != nil {
			fmt.Println("fatal error and recover:", recvr)
			debug.PrintStack()
		}
	}()

	address := req.FormValue("address")
	region := req.FormValue("region")

	var answers *data.PlaceResp
	if address == "" || region == "" {
		answers = &data.PlaceResp{
			Status:  2,
			Message: "Invalid Parameters",
		}
	} else {
		answers = self.api.SearchPlace(address, region)

	}
	output, err := json.Marshal(answers)
	if err == nil {
		fmt.Fprint(w, string(output))
	}
}
