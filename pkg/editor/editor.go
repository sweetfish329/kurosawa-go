package editor

import (
	"context"
	"fmt"
	"time"

	"kurosawa-go/internal/ffmpeg"
	"kurosawa-go/pkg/editor/effect"
)

// Editor は動画編集の主要なインターフェースを提供します。
type Editor struct {
	input    string
	output   string
	effects  []effect.Effect
	duration time.Duration
}

// New は新しい Editor インスタンスを作成します。
func New(input string) *Editor {
	return &Editor{
		input:   input,
		effects: make([]effect.Effect, 0),
	}
}

// Output は出力ファイルを設定します。
func (e *Editor) Output(path string) *Editor {
	e.output = path
	return e
}

// Trim は動画をトリミングします。
func (e *Editor) Trim(start, duration time.Duration) *Editor {
	e.duration = duration
	return e
}

// Resize は動画サイズを変更します。
func (e *Editor) Resize(width, height int) *Editor {
	e.effects = append(e.effects, &effect.ResizeEffect{
		Width:  width,
		Height: height,
	})
	return e
}

// AddEffect はエフェクトを追加します。
func (e *Editor) AddEffect(effect effect.Effect) *Editor {
	e.effects = append(e.effects, effect)
	return e
}

// Process は設定された編集内容を処理します。
func (e *Editor) Process(ctx context.Context) error {
	cmd := ffmpeg.NewCommand().
		WithInput(e.input).
		WithOutput(e.output)

	// エフェクトの適用
	if len(e.effects) > 0 {
		var filters []string
		for _, ef := range e.effects {
			filters = append(filters, ef.Apply(nil))
		}
		cmd.WithFilterComplex(fmt.Sprintf("%s", filters[0]))
	}

	return cmd.Run(ctx)
}
