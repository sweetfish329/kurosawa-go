package errors

import "errors"

var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrProcessNotFound  = errors.New("ffmpeg process not found")
	ErrRecordingRunning = errors.New("recording is already running")
	ErrInvalidEffect    = errors.New("invalid effect parameters")
)
