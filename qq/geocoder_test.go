package qq

import(
	"testing"
)

func TestGeoCoderAPI(t *testing.T) {
	geoCoderAPI := NewGeoCoderAPI()
	resp := geoCoderAPI.GetLocationByAddress("清河橡树湾", "北京")
	if resp.Status != 0{
		t.Error()
	}
}