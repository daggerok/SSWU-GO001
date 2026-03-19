package customclient

import (
	"net"
	"net/http"
	"time"
)

func NewHTTPClient() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:     	 1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  false,
		MaxConnsPerHost: 	 1000,
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second, // overall request timeout
	}
	return client
}
