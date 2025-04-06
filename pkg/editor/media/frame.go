package media

import "time"

// Frame はメディアフレームを表します。
type Frame struct {
	Data      []byte
	Timestamp time.Duration
	Width     int
	Height    int
	Format    string
}

func NewFrame(data []byte, timestamp time.Duration) *Frame {
	return &Frame{
		Data:      data,
		Timestamp: timestamp,
	}
}
