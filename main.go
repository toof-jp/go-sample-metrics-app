package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sample_http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"path", "method", "status"},
	)
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues("/hello", r.Method, "200").Inc()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello"))
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Println("listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
