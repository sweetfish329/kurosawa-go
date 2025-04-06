package config

import "os"

// Config はグローバル設定を表します。
type Config struct {
	FFmpegPath    string
	DefaultWidth  int
	DefaultHeight int
	TempDir       string
}

var globalConfig = &Config{
	FFmpegPath:    "ffmpeg",
	DefaultWidth:  1920,
	DefaultHeight: 1080,
	TempDir:       os.TempDir(),
}

func SetFFmpegPath(path string) { globalConfig.FFmpegPath = path }
