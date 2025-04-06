package media

import (
	"context"
	"time"
)

// Processor はメディア処理の基本インターフェースを定義します。
type Processor interface {
	Process(ctx context.Context) error
}

// Source はメディアソースの基本インターフェースを定義します。
type Source interface {
	// Open はメディアソースを開きます。
	Open(ctx context.Context) error
	// Close はメディアソースを閉じます。
	Close() error
	// Info はメディア情報を返します。
	Info() (*MediaInfo, error)
}

// MediaInfo はメディアの基本情報を表します。
type MediaInfo struct {
	Duration  time.Duration
	Width     int
	Height    int
	FrameRate float64
	Format    string
}
