package main

import (
	"fmt"

	raven "github.com/getsentry/raven-go"
)

func init() {
}

func main() {
	raven.SetDSN("http://de044df139c3400285ad85ac70038e00@127.0.0.1:9000/2")

	raven.CapturePanicAndWait(func() {
		panic(fmt.Errorf("panic"))
	}, map[string]string{"browser": "Chrome"}, &raven.Http{
		Method: "GET",
		URL:    "https://example.com/raven-go",
	})
}
