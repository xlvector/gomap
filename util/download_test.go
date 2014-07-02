package util

import(
	"testing"
)

func Test(t *testing.T) {
	link := "http://api.map.baidu.com/place/v2/search?&q=%E5%98%89%E7%A6%BE%E8%B7%AF21%E5%8F%B71006%E5%AE%A4&scope=2&region=%E5%8E%A6%E9%97%A8&output=json&ak=F827baac9ff8d1e98178bcb0be60fc3b"
	dl := NewDownloader()
	html, err := dl.Download(link)
	if err != nil {
		t.Error(err)
	}
	if len(html) < 10 {
		t.Error("html to small")
	}

	placeResp := Decode(html)
	if placeResp.Status != 0 {
		t.Error("status is not 0")
	}
}