package prom

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var HttpRequestTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_http_request_total",
		Help: "The total number of http requests",
	},
	[]string{"code", "method", "handler"})

var HttpRequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "api_http_request_duration",
		Help: "The duration of http requests",
	},
	[]string{"code", "method", "handler"})

func UpdateMetrics(statusCode int, r *http.Request, begin time.Time) {
	code := strconv.Itoa(statusCode)
	HttpRequestTotal.With(prometheus.Labels{
		"code":    code,
		"method":  r.Method,
		"handler": r.URL.Path,
	}).Inc()
	HttpRequestDuration.With(prometheus.Labels{
		"code":    code,
		"method":  r.Method,
		"handler": r.URL.Path,
	}).Observe(float64(time.Now().Sub(begin).Seconds()))
}
