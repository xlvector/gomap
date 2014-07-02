package baidu

import(
	"testing"
)

func TestPlaceAPI(t *testing.T) {
	placeAPI := NewPlaceAPI()
	resp := placeAPI.Search(map[string]string{
		"query" : "清河橡树湾",
		"region" : "北京",
		})
	if resp.Status != 0 {
		t.Error()
	}

	if len(resp.Results) == 0 {
		t.Error()
	}

	resp = placeAPI.SearchInRegion("清河橡树湾", "北京")

	if resp.Status != 0 {
		t.Error()
	}

	if len(resp.Results) == 0 {
		t.Error()
	}
}