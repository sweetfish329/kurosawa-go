//go:build darwin

package recorder

import "kurosawa-go/internal/ffmpeg"

func applyOSSettings(cmd *ffmpeg.Command) error {
	cmd.WithOption("-f", "avfoundation")
	cmd.WithOption("-i", "1") // screen capture device
	return nil
}
