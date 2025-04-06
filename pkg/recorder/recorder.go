// パッケージ recorder は FFmpeg を使用した画面録画機能を提供します。
// 使用例：
//
//	recorder := recorder.NewRecorder("output.mp4", 30)
//	recorder.WithArea("1920x1080")
//	err := recorder.Start()
//	if (err != nil) {
//		log.Fatal(err)
//	}
//	// ... 録画中 ...
//	err = recorder.Stop()
package recorder

import (
	"context"
	"fmt"

	"kurosawa-go/internal/ffmpeg"
)

// Options は録画のオプションを定義します。
type Options struct {
	Area      string
	Framerate int
	Quality   string
}

// NewOptions はデフォルトのオプションを返します。
func NewOptions() *Options {
	return &Options{
		Area:      "1920x1080",
		Framerate: 30,
		Quality:   "medium",
	}
}

// Recorder は画面レコーダーを表します。
type Recorder struct {
	area       string
	framerate  int
	outputPath string
	cmd        *ffmpeg.Command
	cancel     context.CancelFunc // doneチャネルの代わりにcancel funcを保持
}

// NewRecorder は指定された出力パスとフレームレートで新しい画面レコーダーを作成します。
func NewRecorder(outputPath string, framerate int) *Recorder {
	return &Recorder{
		framerate:  framerate,
		outputPath: outputPath,
	}
}

// WithArea は録画領域を設定します。
func (r *Recorder) WithArea(area string) *Recorder {
	r.area = area
	return r
}

// WithFramerate はフレームレートを設定します。
func (r *Recorder) WithFramerate(framerate int) *Recorder {
	r.framerate = framerate
	return r
}

// Start はコンテキスト付きで画面録画を開始します。
func (r *Recorder) Start(ctx context.Context) error {
	// 親contextから新しいcancelable contextを作成
	recordCtx, cancel := context.WithCancel(ctx)
	r.cancel = cancel

	cmd := ffmpeg.NewCommand()

	if err := applyOSSettings(cmd); err != nil {
		return fmt.Errorf("failed to apply OS settings: %w", err)
	}

	cmd.WithOption("-framerate", fmt.Sprintf("%d", r.framerate))

	if r.area != "" {
		cmd.WithOption("-video_size", r.area)
	}

	cmd.WithOutput(r.outputPath)

	if err := cmd.Start(recordCtx); err != nil {
		r.cancel()
		return fmt.Errorf("failed to start recording: %w", err)
	}

	r.cmd = cmd

	return nil
}

// Stop は画面録画を停止します。
func (r *Recorder) Stop() error {
	if r.cmd == nil {
		return fmt.Errorf("recording is not running")
	}
	if r.cancel != nil {
		r.cancel() // contextをキャンセルして録画を停止
	}
	return r.cmd.Stop()
}

// Recorder はシンプルな画面録画インターフェースを提供します。
func Record(ctx context.Context, output string, opts *Options) error {
	if opts == nil {
		opts = NewOptions()
	}

	cmd := ffmpeg.NewCommand()
	if err := applyOSSettings(cmd); err != nil {
		return fmt.Errorf("failed to configure recorder: %w", err)
	}

	cmd.WithOption("-framerate", fmt.Sprintf("%d", opts.Framerate))
	if opts.Area != "" {
		cmd.WithOption("-video_size", opts.Area)
	}
	cmd.WithOption("-preset", opts.Quality)
	cmd.WithOutput(output)

	return cmd.Run(ctx)
}
