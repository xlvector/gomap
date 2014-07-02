package gomap

import(
    "encoding/json"
    "github.com/xlvector/gomap/qq"
    "github.com/xlvector/gomap/baidu"
    "github.com/xlvector/gomap/data"
    "strconv"
    "net/http"
    "runtime/debug"
    "fmt"
)

type NearByService struct {
    qqApi      *qq.NearByAPI
    baiduApi   *baidu.NearByAPI
}

func NewNearByService() *NearByService {
    service := NearByService{
        qqApi: qq.NewNearByAPI(),
        baiduApi: baidu.NewNearByAPI(),
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
    provider := req.FormValue("provider")

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
        if provider == "baidu" {
            answers = self.baiduApi.NearBy(address, query, region, radiusValue)
        } else {
            answers = self.qqApi.NearBy(address, query, region, radiusValue)
        }
        
    }
    output, err := json.Marshal(answers)
    if err == nil{
        fmt.Fprint(w, string(output))
    }
}