package resource

import (
	"context"
	"io"
)

// Manager はリソースのライフサイクルを管理します。
type Manager interface {
	// Register は新しいリソースを登録します。
	Register(resource io.Closer)
	// Cleanup は登録されたすべてのリソースを解放します。
	Cleanup() error
	// WithContext はコンテキストに基づいてリソースを管理します。
	WithContext(ctx context.Context) Manager
}

type resourceManager struct {
	resources []io.Closer
	ctx       context.Context
}
