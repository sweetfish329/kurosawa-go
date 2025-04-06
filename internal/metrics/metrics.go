package metrics

import "time"

type Metrics interface {
	RecordDuration(name string, duration time.Duration)
	RecordError(name string, err error)
	RecordFFmpegCommand(args []string)
}
