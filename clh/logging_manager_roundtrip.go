package clh

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
	LM      *LoggingManager
	AppName string
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request, appName string) (*http.Response, error) {
	var logger *log.Logger
	if lrt.LM.ShouldLog(lrt.AppName) {
		fmt.Printf("Request: %s %s\n\n\n", req.Method, req.URL.String())
		logger.Printf("%v", logrus.Fields{
			"Response":       "Request",
			"AppName":        appName,
			"request_method": req.Method,
			"request_uri":    req.URL.String(),
		})
	}
	resp, err := lrt.Proxied.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if lrt.LM.ShouldLog(lrt.AppName) {
		log.Printf("Response: %d", resp.StatusCode)
		fmt.Printf("Response: %d\n\n", resp.StatusCode)
		logger.Printf("%v", logrus.Fields{
			"Response":     "Response",
			"AppName":      appName,
			"ResponseBody": resp.Body,
		})
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
