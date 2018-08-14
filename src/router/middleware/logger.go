package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func LoggingMw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log request
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		// use recorder to observer response
		c := httptest.NewRecorder()
		start := time.Now()
		next.ServeHTTP(c, r)

		// log response
		log.Printf("Response: %s %d %s", time.Since(start), c.Code, string(c.Body.Bytes()))

		// write response from recorder to actual response writer
		for k, v := range c.HeaderMap {
			w.Header().Set(k, strings.Join(v, ""))
		}
		w.WriteHeader(c.Code)
		c.Body.WriteTo(w)
	}
}
