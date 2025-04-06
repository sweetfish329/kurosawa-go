package effect

// Effect はクリップに適用可能なエフェクトを表します。
type Effect interface {
	Apply(clip *clip.Clip) string
	Validate() error
}

// ...existing effect implementations...
