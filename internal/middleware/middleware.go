package middleware

import (
	"context"
	"kurosawa-go/internal/metrics"
	"time"
)

type ProcessFunc func(context.Context) error

type Middleware func(ProcessFunc) ProcessFunc

// WithMetrics はメトリクス収集を行うミドルウェアです。
func WithMetrics(metrics metrics.Metrics) Middleware {
	return func(next ProcessFunc) ProcessFunc {
		return func(ctx context.Context) error {
			start := time.Now()
			err := next(ctx)
			metrics.RecordDuration("process_time", time.Since(start))
			return err
		}
	}
}
