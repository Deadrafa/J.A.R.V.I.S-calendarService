package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HttpRequests *prometheus.CounterVec
	HttpDuration *prometheus.HistogramVec
	DbQueryTime  *prometheus.HistogramVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		HttpRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		HttpDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration distribution",
				Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2},
			},
			[]string{"path"},
		),
		DbQueryTime: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "db_query_duration_seconds",
				Help:    "Database query duration",
				Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1},
			},
			[]string{"query_type"},
		),
	}
}

func (m *Metrics) Register() {
	prometheus.MustRegister(
		m.HttpRequests,
		m.HttpDuration,
		m.DbQueryTime,
	)
}

func (m *Metrics) RecordHTTPRequest(method, path, status string) {
	m.HttpRequests.WithLabelValues(method, path, status).Inc()
}

func (m *Metrics) RecordHTTPDuration(path string, duration float64) {
	m.HttpDuration.WithLabelValues(path).Observe(duration)
}

func (m *Metrics) RecordDBQuery(queryType string, duration float64) {
	m.DbQueryTime.WithLabelValues(queryType).Observe(duration)
}

func (m *Metrics) RegisterCustomMetric(metric prometheus.Collector) {
	prometheus.MustRegister(metric)
}
