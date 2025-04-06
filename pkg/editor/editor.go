// Package editor provides video editing functionality using FFmpeg.
// Example usage:
//
//	editor := editor.NewEditor("input.mp4")
//	err := editor.
//		Trim("00:00:00", "00:00:30").
//		Resize("1280x720").
//		SaveAs("output.mp4")
package editor

import (
	"fmt"
	"os/exec"
	"time"

	"kurosawa-go/ffmpeg"
)

// Editor represents a video editor instance
type Editor struct {
	inputPath string
	startTime string
	duration  string
	size      string
	filters   []string
}

// NewEditor creates a new video editor instance
func NewEditor(inputPath string) *Editor {
	return &Editor{
		inputPath: inputPath,
		filters:   make([]string, 0),
	}
}

// Trim sets the start time and duration for video trimming
func (e *Editor) Trim(start, duration string) *Editor {
	e.startTime = start
	e.duration = duration
	return e
}

// Resize sets the output video size
func (e *Editor) Resize(size string) *Editor {
	e.size = size
	return e
}

// SaveAs processes the video with the specified settings and saves to outputPath
func (e *Editor) SaveAs(outputPath string) error {
	args := []string{"-i", e.inputPath}

	if e.startTime != "" {
		args = append(args, "-ss", e.startTime)
	}
	if e.duration != "" {
		args = append(args, "-t", e.duration)
	}
	if e.size != "" {
		args = append(args, "-s", e.size)
	}

	args = append(args, "-y", outputPath)

	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to process video: %w", err)
	}

	return nil
}

// Clip represents a video clip.
type Clip struct {
	path   string
	editor *Editor
}

// Clip creates a new video clip.
func (e *Editor) Clip(path string) *Clip {
	return &Clip{path: path, editor: e}
}

// TrimOptions represents options for trimming a video.
type TrimOptions struct {
	StartTime time.Duration
	EndTime   time.Duration
}

// Trim trims a video clip.
func (c *Clip) Trim(outputPath string, options TrimOptions) error {
	// FFmpeg command.
	cmd := ffmpeg.NewCommand().
		WithInput(c.path).
		WithOption("-ss", fmt.Sprintf("%.3f", options.StartTime.Seconds())).
		WithOption("-to", fmt.Sprintf("%.3f", options.EndTime.Seconds())).
		WithOutput(outputPath)

	return cmd.Run()
}

// Image represents an image.
type Image struct {
	path   string
	editor *Editor
}

// Image creates a new image.
func (e *Editor) Image(path string) *Image {
	return &Image{path: path, editor: e}
}

// ResizeOptions represents options for resizing an image.
type ResizeOptions struct {
	Width  int
	Height int
}

// Resize resizes an image.
func (i *Image) Resize(outputPath string, options ResizeOptions) error {
	// FFmpeg command.
	cmd := ffmpeg.NewCommand().
		WithInput(i.path).
		WithOption("-vf", fmt.Sprintf("scale=%d:%d", options.Width, options.Height)).
		WithOutput(outputPath)

	return cmd.Run()
}
