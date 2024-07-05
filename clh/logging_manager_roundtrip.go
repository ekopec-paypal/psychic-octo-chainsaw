package clh

import (
	"log"
	"net/http"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
	LM      *LoggingManager
	AppName string
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if lrt.LM.ShouldLog(lrt.AppName) {
		log.Printf("Request: %s %s", req.Method, req.URL.String())
	}
	resp, err := lrt.Proxied.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if lrt.LM.ShouldLog(lrt.AppName) {
		log.Printf("Response: %d", resp.StatusCode)
	}
	return resp, nil
}

func NewLoggingRoundTripper(proxied http.RoundTripper, lm *LoggingManager, appName string) *LoggingRoundTripper {
	return &LoggingRoundTripper{
		Proxied: proxied,
		LM:      lm,
		AppName: appName,
	}
}
