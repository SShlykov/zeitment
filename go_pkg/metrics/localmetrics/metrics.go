package localmetrics

import (
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"log/slog"
	"sync"
)

type MetricValue struct {
	Count int
	Sum   float64
}

// LocalMetrics структура для локального хранения метрик.
type LocalMetrics struct {
	logger    loggerPkg.Logger
	mu        sync.Mutex
	counters  map[string]int
	summaries map[string]MetricValue
}

func NewLocalMetrics(logger loggerPkg.Logger) *LocalMetrics {
	return &LocalMetrics{
		logger:    logger,
		counters:  make(map[string]int),
		summaries: make(map[string]MetricValue),
	}
}

func (l *LocalMetrics) IncCounter(name string, labels ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	key := buildMetricKey(name, labels)
	if _, exists := l.counters[key]; !exists {
		l.counters[key]++
	} else {
		l.counters[key] = 1
	}
	l.logger.Info("IncCounter", slog.String("name", name))
}

func (l *LocalMetrics) ObserveHistogram(name string, value float64, labels ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	key := buildMetricKey(name, labels)
	summary, exists := l.summaries[key]
	if !exists {
		summary = MetricValue{}
	}
	summary.Count++
	summary.Sum += value
	l.summaries[key] = summary

	l.logger.Info("ObserveSummary", slog.String("name", name), slog.Float64("value", value))
}

// buildMetricKey генерирует уникальный ключ для метрики на основе имени и меток.
func buildMetricKey(name string, _ []string) string {
	return name
}
