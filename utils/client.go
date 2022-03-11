package utils

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	withRetry    *resty.Client
	withoutRetry *resty.Client
}

func New(host string, retryCount int, debugMode bool) *Client {
	withRetry := resty.
		New().
		SetHostURL(host).
		SetRetryCount(retryCount).
		SetTimeout(120 * time.Second).
		SetRetryWaitTime(100 * time.Millisecond).
		SetRetryMaxWaitTime(500 * time.Millisecond).
		SetDebug(debugMode)
	withoutRetry := resty.
		New().
		SetHostURL(host).
		SetRetryCount(0).
		SetTimeout(360 * time.Second).
		SetDebug(debugMode)
	return &Client{
		withRetry,
		withoutRetry,
	}
}
