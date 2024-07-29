package store

import "gopl.io/hw/mocks/models"

type Store interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}

type MetricsStorage interface {
	ReadMetric(name string) (models.Metrics, error)
	ReadMetrics() []models.Metrics
	UpdateMetric(metric models.Metrics) error
	SaveMetrics(metrics []models.Metrics)
	PingDB() error
}

type FileStorage interface {
	WriteMetrics(isLoop bool)
	ReadMetrics()
}
