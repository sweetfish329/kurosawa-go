package clip

import (
	"time"

	"kurosawa-go/pkg/editor/effect"
)

// MediaSource はメディアソースのインターフェースを定義します。
type MediaSource interface {
	Duration() time.Duration
	Path() string
	MediaType() MediaType
}

type MediaType int

const (
	MediaTypeVideo MediaType = iota
	MediaTypeImage
	MediaTypeText
)

// Clip はメディアクリップの基本実装を提供します。
type Clip struct {
	source   MediaSource
	start    time.Duration
	duration time.Duration
	position Position
	effects  []effect.Effect
}

// Position はクリップの位置情報を表します。
type Position struct {
	X int
	Y int
}

// NewClip は新しいクリップを作成します。
func NewClip(source MediaSource, start, duration time.Duration, position Position) *Clip {
	return &Clip{
		source:   source,
		start:    start,
		duration: duration,
		position: position,
		effects:  make([]effect.Effect, 0),
	}
}

// Source はメディアソースを返します。
func (c *Clip) Source() MediaSource {
	return c.source
}

// Start はクリップの開始時間を返します。
func (c *Clip) Start() time.Duration {
	return c.start
}

// Duration はクリップの長さを返します。
func (c *Clip) Duration() time.Duration {
	return c.duration
}

// Position はクリップの位置を返します。
func (c *Clip) Position() Position {
	return c.position
}

// Effects はクリップに適用されているエフェクトのスライスを返します。
func (c *Clip) Effects() []effect.Effect {
	return c.effects
}

// AddEffect はクリップにエフェクトを追加します。
func (c *Clip) AddEffect(e effect.Effect) {
	c.effects = append(c.effects, e)
}
