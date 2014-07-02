package baidu

import(
	"testing"
)

func TestNearbyAPI(t *testing.T) {
	api := NewNearByAPI()
	resp := api.NearBy("清河橡树湾","地铁", "北京", 5000)
	if resp.Status != 0 {
		t.Error()
	}

	if len(resp.Results) == 0 {
		t.Error()
	}
}