// Package recorder provides screen recording functionality using FFmpeg.
// Example usage:
//
//	recorder := recorder.NewRecorder("output.mp4", 30)
//	recorder.WithArea("1920x1080")
//	err := recorder.Start()
//	if err != nil {
//		log.Fatal(err)
//	}
//	// ... recording ...
//	err = recorder.Stop()
package recorder

import (
	"fmt"
	"os"

	"kurosawa-go/internal/ffmpeg"
)

// Recorder represents a screen recorder.
type Recorder struct {
	area       string
	framerate  int
	outputPath string
	cmd        *ffmpeg.Command
}

// NewRecorder creates a new screen recorder with specified output path and framerate.
func NewRecorder(outputPath string, framerate int) *Recorder {
	return &Recorder{
		framerate:  framerate,
		outputPath: outputPath,
	}
}

// WithArea sets the recording area.
func (r *Recorder) WithArea(area string) *Recorder {
	r.area = area
	return r
}

// WithFramerate sets the frame rate.
func (r *Recorder) WithFramerate(framerate int) *Recorder {
	r.framerate = framerate
	return r
}

// Start starts the screen recording.
func (r *Recorder) Start() error {
	cmd := ffmpeg.NewCommand()

	// Apply OS-specific settings
	if err := applyOSSettings(cmd); err != nil {
		return fmt.Errorf("failed to apply OS settings: %w", err)
	}

	cmd.WithOption("-framerate", fmt.Sprintf("%d", r.framerate))

	if r.area != "" {
		cmd.WithOption("-video_size", r.area)
	}

	cmd.WithOutput(r.outputPath)

	process, err := cmd.StartProcess()
	if err != nil {
		return fmt.Errorf("failed to start recording: %w", err)
	}

	r.cmd = process
	return nil
}

// Stop stops the screen recording.
func (r *Recorder) Stop() error {
	if r.cmd == nil || r.cmd.Process == nil {
		return fmt.Errorf("recording is not running")
	}

	// Send SIGINT to stop ffmpeg
	if err := r.cmd.Process.Signal(os.Interrupt); err != nil {
		return fmt.Errorf("failed to stop recording: %w", err)
	}

	// Wait for the command to exit.
	if err := r.cmd.Wait(); err != nil {
		return fmt.Errorf("failed to wait for recording to stop: %w", err)
	}

	return nil
}
