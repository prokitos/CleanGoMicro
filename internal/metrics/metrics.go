package metrics

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var RequestDuration *prometheus.SummaryVec = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "http_request_duration_seconds",
		Help:       "Duration of HTTP requests in seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"method", "path"},
)
var RequestStatus *prometheus.SummaryVec = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "http_status_total",
		Help:       "Total number of HTTP requests and status",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"status"},
)

func Observer(c *fiber.Ctx) error {
	start := time.Now()
	//c.Next()
	duration := time.Since(start).Seconds()

	RequestStatus.WithLabelValues(strconv.Itoa(c.Response().StatusCode())).Observe(1)
	RequestDuration.WithLabelValues(c.Method(), c.Path()).Observe(duration)

	return c.Next()
}
