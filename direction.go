package gomap

import(
    "encoding/json"
    "github.com/xlvector/gomap/baidu"
    "github.com/xlvector/gomap/data"
    "net/http"
    "fmt"
    "runtime/debug"
)

type DirectionService struct {
    api      *baidu.DirectionAPI
}

func NewDirectionService() *DirectionService {
    service := DirectionService{
        api: baidu.NewDirectionAPI(),
    }
    return &service
}

func (self *DirectionService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    defer func() {
        if recvr := recover(); recvr != nil {
            fmt.Println("fatal error and recover:", recvr)
            debug.PrintStack()
        }
    }()

    origin := req.FormValue("origin")
    destination := req.FormValue("destination")
    region := req.FormValue("region")

    var answers *data.DirectionResp
    if origin == "" || region == "" || destination == "" {
        answers = &data.DirectionResp{
            Status: 2,
            Message: "Invalid Parameters",
        }
    } else {
        answers = self.api.TransitSequence(origin, destination, region)        
    }
    output, err := json.Marshal(answers)
    if err == nil{
        fmt.Fprint(w, string(output))
    }
}