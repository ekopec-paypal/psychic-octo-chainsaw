package clh

import (
	"log"
	"net/http"
)

func LoggingMiddleware(lm *LoggingManager, appName string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if lm.ShouldLog(appName) {
			// Log the request
			log.Printf("Request: %s %s", r.Method, r.URL.Path)

			// Log the response (this example assumes a simple logging, you can enhance it)
			lrw := NewLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)
			log.Printf("Response: %d", lrw.statusCode)
		} else {
			next.ServeHTTP(w, r) // No-op: Just call the next handler
		}
	})
}

// LoggingResponseWriter wraps http.ResponseWriter to capture the status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
