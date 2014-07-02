package gomap

import(
    "encoding/json"
    "github.com/xlvector/gomap/qq"
    "github.com/xlvector/gomap/data"
    "strconv"
    "net/http"
    "runtime/debug"
    "fmt"
)

type NearByService struct {
    api     *qq.NearByAPI
}

func NewNearByService() *NearByService {
    service := NearByService{
        api : qq.NewNearByAPI(),
    }
    return &service
}

func (self *NearByService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    defer func() {
        if recvr := recover(); recvr != nil {
            fmt.Println("fatal error and recover:", recvr)
            debug.PrintStack()
        }
    }()

    address := req.FormValue("address")
    region := req.FormValue("region")
    query := req.FormValue("query")
    radius := req.FormValue("radius")

    var answers *data.NearByResp
    if address == "" || region == "" || query == "" {
        answers = &data.NearByResp{
            Status: 2,
            Message: "Invalid Parameters",
        }
    } else {
        if radius == "" {
            radius = "5000"
        }
        radiusValue, _ := strconv.Atoi(radius)
        answers = self.api.NearBy(address, query, region, radiusValue)
    }
    output, err := json.Marshal(answers)
    if err == nil{
        fmt.Fprint(w, string(output))
    }
}