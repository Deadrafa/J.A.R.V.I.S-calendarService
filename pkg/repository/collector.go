package repository

import "github.com/prometheus/client_golang/prometheus"

type Collector interface {
	RecordHTTPRequest(method, path, status string)
	RecordHTTPDuration(path string, duration float64)
	RecordDBQuery(queryType string, duration float64)
	RegisterCustomMetric(metric prometheus.Collector)
}
