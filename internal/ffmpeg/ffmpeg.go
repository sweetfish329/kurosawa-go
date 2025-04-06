package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// Command は FFmpeg コマンドビルダーを表します。
type Command struct {
	inputs  []string
	outputs []string
	options []string
	filters []string
	process *exec.Cmd // 実行中のプロセスを保持
}

// NewCommand は新しい FFmpeg コマンドビルダーを作成します。
func NewCommand() *Command {
	return &Command{
		inputs:  make([]string, 0),
		outputs: make([]string, 0),
		options: make([]string, 0),
		filters: make([]string, 0),
	}
}

// WithInput は入力ファイルをコマンドに追加します。
func (c *Command) WithInput(input string) *Command {
	c.inputs = append(c.inputs, input)
	return c
}

// WithOutput は出力ファイルをコマンドに追加します。
func (c *Command) WithOutput(output string) *Command {
	c.outputs = append(c.outputs, output)
	return c
}

// WithOption は FFmpeg オプションをコマンドに追加します。
func (c *Command) WithOption(name, value string) *Command {
	c.options = append(c.options, name, value)
	return c
}

// WithFilter はフィルターをコマンドに追加します。
func (c *Command) WithFilter(filter string) *Command {
	c.filters = append(c.filters, filter)
	return c
}

// WithFilterComplex は複雑なフィルターグラフをコマンドに追加します。
func (c *Command) WithFilterComplex(filter string) *Command {
	if filter != "" {
		c.options = append(c.options, "-filter_complex", filter)
	}
	return c
}

// buildArgs は FFmpeg コマンドの引数を構築します。
func (c *Command) buildArgs() []string {
	args := make([]string, 0)

	// Add global options (before inputs)
	args = append(args, "-y") // 既存ファイルを上書き

	// Add inputs
	for _, input := range c.inputs {
		args = append(args, "-i", input)
	}

	// Add options and filter complex
	args = append(args, c.options...)

	// Add outputs
	args = append(args, c.outputs...)

	return args
}

// Start はコンテキスト付きで FFmpeg プロセスを開始します。
func (c *Command) Start(ctx context.Context) error {
	args := c.buildArgs()
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg process: %w", err)
	}

	c.process = cmd
	return nil
}

// Stop は FFmpeg プロセスを正常に停止します。
func (c *Command) Stop() error {
	if c.process == nil || c.process.Process == nil {
		return fmt.Errorf("no running ffmpeg process")
	}

	if err := c.process.Process.Signal(os.Interrupt); err != nil {
		return fmt.Errorf("failed to stop ffmpeg process: %w", err)
	}

	if err := c.process.Wait(); err != nil {
		return fmt.Errorf("failed to wait for ffmpeg process: %w", err)
	}

	return nil
}

// Run はコンテキスト付きで FFmpeg コマンドを実行し、完了を待ちます。
func (c *Command) Run(ctx context.Context) error {
	if err := c.Start(ctx); err != nil {
		return err
	}
	return c.process.Wait()
}

// Process は基底の exec.Cmd を返します。
func (c *Command) Process() *exec.Cmd {
	return c.process
}
