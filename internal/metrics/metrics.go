package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var requestMetric = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace:  "clean",
	Subsystem:  "http",
	Name:       "request",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, // чтобы правильно видеть запросы, например если есть проблемы с редкими долгими запросами
}, []string{"status"})

func observeRequest(durat time.Duration, status int) {
	requestMetric.WithLabelValues(strconv.Itoa(status)).Observe(durat.Seconds())
}
