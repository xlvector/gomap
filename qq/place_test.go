package qq

import(
    "testing"
)

func TestPlaceAPI(t *testing.T) {
    placeAPI := NewPlaceAPI()

    resp := placeAPI.SearchInRegion("清河橡树湾", "北京")

    if resp.Status != 0 {
        t.Error()
    }
}