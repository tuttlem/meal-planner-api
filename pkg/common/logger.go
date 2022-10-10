package common

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type LogItems struct {
	ISOTime        time.Time   `json:"isoTime"`
	UnixTime       int64       `json:"unitTime"`
	IP             string      `json:"ip"`
	Method         string      `json:"method"`
	Host           string      `json:"host"`
	User           string      `json:"user"`
	Path           string      `json:"path"`
	Query          string      `json:"query"`
	Protocol       string      `json:"protocol"`
	ContentType    string      `json:"contentType"`
	ContentLength  int64       `json:"contentLength"`
	ResponseStatus int         `json:"responseStatus"`
	ResponseSize   int         `json:"responseSize"`
	Headers        http.Header `json:"headers"`
	TLSData        TLSData     `json:"tls"`

	RequestProcessingTime int64 `json:"requestProcessingTime"`
	LogProcessingTime     int64 `json:"logProcessingTime"`
}

type TLSData struct {
	TLSVersion     uint16 `json:"version"`
	TLSCipherUsed  uint16 `json:"cipher"`
	TLSMutualProto bool   `json:"mutualProtocol"`
}

var FormatJSON = func(log LogItems) string {
	logline, _ := json.Marshal(log)
	return fmt.Sprintf("%s\n", logline)
}

func JSONLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// All time values are in nanoseconds
		start := time.Now()
		c.Next() // Request is processed here
		stop := time.Now().UnixNano()

		log := LogItems{
			ISOTime:        start,
			UnixTime:       start.UnixNano(),
			IP:             c.ClientIP(),
			Method:         c.Request.Method,
			Host:           c.Request.Host,
			User:           c.Request.URL.User.Username(),
			Path:           c.Request.URL.EscapedPath(),
			Query:          c.Request.URL.RawQuery,
			Protocol:       c.Request.Proto,
			ContentType:    c.ContentType(),
			ContentLength:  c.Request.ContentLength,
			ResponseStatus: c.Writer.Status(),
			ResponseSize:   c.Writer.Size(),
			// Headers are now placed arrays, eg. "Dnt":["1"]
			// Header type is not a struct, but is the format a problem?
			Headers: c.Request.Header,
		}

		if c.Request.TLS != nil {
			// https://golang.org/pkg/crypto/tls/#pkg-constants
			log.TLSData = TLSData{
				TLSVersion:     c.Request.TLS.Version,
				TLSCipherUsed:  c.Request.TLS.CipherSuite,
				TLSMutualProto: c.Request.TLS.NegotiatedProtocolIsMutual,
			}
		}

		log.RequestProcessingTime = stop - log.UnixTime
		log.LogProcessingTime = time.Now().UnixNano() - stop
		fmt.Fprint(os.Stdout, FormatJSON(log))
	}
}
