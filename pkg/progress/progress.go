package progress

import "time"

// Progress は処理の進捗状況を表します。
type Progress struct {
	// 進捗率（0-100）
	Percent float64
	// 現在の処理ステージ（例：「エンコード中」「フィルター適用中」）
	Stage string
	// 処理開始からの経過時間
	Elapsed time.Duration
	// 予測残り時間
	Remaining time.Duration
	// 処理中のエラー（もしあれば）
	Error error
}

// Reporter は進捗状況を報告するインターフェースです。
type Reporter interface {
	// Update は進捗状況を更新します。
	Update(progress Progress)
}

// Options は進捗報告のオプションを定義します。
type Options struct {
	// 進捗更新の間隔
	UpdateInterval time.Duration
	// 進捗報告を受け取るチャネル
	Channel chan<- Progress
	// カスタムレポーター
	Reporter Reporter
}
