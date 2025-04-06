//go:build linux

package recorder

import "kurosawa-go/internal/ffmpeg"

func applyOSSettings(cmd *ffmpeg.Command) error {
	cmd.WithOption("-f", "x11grab")
	cmd.WithOption("-i", ":0.0")
	return nil
}
