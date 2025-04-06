package film

import (
	"context"
	"time"

	"kurosawa-go/pkg/editor/clip"
)

// Film は複数のクリップから構成される動画を表します。
type Film struct {
	clips    []*clip.Clip
	width    int
	height   int
	duration time.Duration
}

type Builder struct {
	clips  []*clip.Clip
	width  int
	height int
}

// NewBuilder は新しい FilmBuilder を作成します。
func NewBuilder(width, height int) *Builder {
	return &Builder{
		clips:  make([]*clip.Clip, 0),
		width:  width,
		height: height,
	}
}

// ...existing code for builder methods...

// WriteFilm は Film を動画ファイルに書き出します。
func Write(ctx context.Context, film *Film, outputPath string) error {
	// ...existing WriteFilm implementation...
}
