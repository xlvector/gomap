package util

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Downloader struct {
	client *http.Client
}

func NewDownloader() *Downloader {
	ret := Downloader{}
	ret.client = &http.Client{
		Transport: &http.Transport{
			Dial:                  dialTimeout,
			DisableKeepAlives:     true,
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
	return &ret
}

func dialTimeout(network, addr string) (net.Conn, error) {
	timeout := 10 * time.Second
	deadline := time.Now().Add(timeout)
	c, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	c.SetDeadline(deadline)
	return c, nil
}

func (self *Downloader) Download(link string) (string, error){
	req, err := http.NewRequest("GET", link, nil)
	if err != nil || req == nil || req.Header == nil {
		return "", err
	}
	resp, err := self.client.Do(req)

	if err != nil || resp == nil || resp.Body == nil {
		return "", err
	} else {
		defer resp.Body.Close()
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(html), nil
	}
}