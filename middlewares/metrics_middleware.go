package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
)

func InitPrometheusMetrics() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		timer := prometheus.NewTimer(httpRequestDuration.WithLabelValues(
			c.Path(), c.Request().Method, "200",
		))
		defer timer.ObserveDuration()
		err := next(c)

		status := http.StatusOK
		if err != nil {
			status = c.Response().Status
		}
		httpRequestsTotal.WithLabelValues(c.Path(), c.Request().Method, http.StatusText(status)).Inc()

		return err
	}
}

func PrometheusHandler() echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}
