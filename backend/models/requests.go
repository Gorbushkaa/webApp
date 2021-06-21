package models

import (
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
	db "webApp/db"
	"webApp/settings"
)

type Request_config struct {
	Urls           []string
	Requests_state bool
	RPM            float64
}

var Config = Request_config{
	Urls:           []string{settings.STATUS_URL, settings.DELAY_URL},
	Requests_state: false,
	RPM:            settings.RPM}

func make_request(url string) {
	req, _ := http.NewRequest("GET", url, nil)
	var start time.Time
	var conn_time time.Duration
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			conn_time = time.Duration(time.Since(start).Nanoseconds())
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}
	db.SetRequest(db.Request{Timestamp: time.Now(), Req_time: conn_time, Url: url})
}

func Run_requests() {
	for Config.Requests_state {
		for _, url := range Config.Urls {
			time.Sleep(time.Duration(Config.RPM*float64(time.Second)))
			go make_request(url)
		}
	}
}
