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

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var logger *log.Logger
	fmt.Printf("RoundTrip %s %+v\n\n\n", lrt.AppName, req)
	if lrt.LM.ShouldLog(lrt.AppName) {
		fmt.Printf("Request: %s %s\n\n\n", req.Method, req.URL.String())
		logger.Printf("%v", logrus.Fields{
			"request":     req.Method,
			"app_name":    lrt.AppName,
			"request_url": req.URL.String(),
		})
	}
	resp, err := lrt.Proxied.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if lrt.LM.ShouldLog(lrt.AppName) {
		fmt.Printf("Response: %d\n\n", resp.StatusCode)
		logger.Printf("%v", logrus.Fields{
			"Response":      resp.Status,
			"status_code":   resp.StatusCode,
			"app_name":      lrt.AppName,
			"response_body": resp.Body,
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
