package observability

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"service", "status"},
	)
	once sync.Once
)

func InitPrometheus() {
	once.Do(func() {
		prometheus.MustRegister(HttpRequests)
	})
}
