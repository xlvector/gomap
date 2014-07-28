package util

import (
	"github.com/xlvector/gomap/data"
	"math"
)

const (
	EARTH_RADIUS = 6378.137
	PI           = 3.1415926535898
	EAST         = "东"
	WEST         = "西"
	NORTH        = "北"
	SOUTH        = "南"
	UNKNOWN      = ""
	NORTHEAST    = "东北"
	NORTHWEST    = "西北"
	SOUTHEAST    = "东南"
	SOUTHWEST    = "西南"
)

func rad(d float64) float64 {
	return d * PI / 180.0
}

func Distance(lat1, lng1, lat2, lng2 float64) float64 {
	radLat1 := rad(lat1)
	radLat2 := rad(lat2)
	a := radLat1 - radLat2
	b := rad(lng1) - rad(lng2)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2.0), 2.0)+
		math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2.0), 2.0)))
	s = s * EARTH_RADIUS
	return math.Abs(s)
}

func OrientationInChina(src *data.Location, dst *data.Location) string {
	dlat := dst.Lat - src.Lat
	dlng := dst.Lng - src.Lng

	dis := Distance(src.Lat, src.Lng, dst.Lat, dst.Lng)
	if dis < 1.0 {
		return UNKNOWN
	}

	d1 := Distance(src.Lat, src.Lng, src.Lat, dst.Lng)
	d2 := Distance(src.Lat, src.Lng, dst.Lat, src.Lng)
	if dlng >= 0 {
		if d2 < 0.4142*d1 {
			return EAST
		} else if d1 < d2*0.4142 {
			if dlat > 0 {
				return NORTH
			} else {
				return SOUTH
			}
		} else {
			if dlat > 0 {
				return NORTHEAST
			} else {
				return SOUTHEAST
			}
		}
	} else {
		if d2 < 0.4142*d1 {
			return WEST
		} else if d1 < d2*0.4142 {
			if dlat > 0 {
				return NORTH
			} else {
				return SOUTH
			}
		} else {
			if dlat > 0 {
				return NORTHWEST
			} else {
				return SOUTHWEST
			}
		}
	}
}
