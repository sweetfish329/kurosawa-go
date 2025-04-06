//go:build windows

package recorder

import "kurosawa-go/internal/ffmpeg"

func applyOSSettings(cmd *ffmpeg.Command) error {
	cmd.WithOption("-f", "gdigrab")
	cmd.WithOption("-i", "desktop")
	return nil
}
