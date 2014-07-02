package data

type Location struct {
    Lng float64 `json:"lng"`
    Lat float64 `json:"lat"`
}

type Result struct {
    Name        string      `json:"name"`
    Location    Location    `json:"location"`
    Address     string      `json:"address"`
    Distance	float64		`json:"distance"`
}

type NearByResp struct {
    Status      int       	`json:"status"`
    Message     string      `json:"message"`
    Results     []Result    `json:"results"`
}