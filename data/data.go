package data

type Location struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type NearyByResult struct {
	Name        string   `json:"name"`
	Location    Location `json:"location"`
	Address     string   `json:"address"`
	Distance    float64  `json:"distance"`
	Orientation string   `json:"orientation"`
}

type NearByResp struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Results []NearyByResult `json:"results"`
	Center  NearyByResult   `json:"center"`
}

type DirectionResp struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Result  DirectionResult `json:"result"`
}

type DirectionResult struct {
	Routes []DirectionRoute `json:"routes"`
}

type DirectionRoute struct {
	Sequence []string `json:"seq"`
}
