package metrics

import "time"

// Collector はメトリクス収集の機能を提供します。
type Collector interface {
	// CollectProcessMetrics はプロセスのメトリクスを収集します。
	CollectProcessMetrics() (*ProcessMetrics, error)
	// CollectPerformanceMetrics はパフォーマンスメトリクスを収集します。
	CollectPerformanceMetrics() (*PerformanceMetrics, error)
}

// ProcessMetrics はプロセスのメトリクス情報を表します。
type ProcessMetrics struct {
	CPUUsage    float64
	MemoryUsage int64
	Duration    time.Duration
}
