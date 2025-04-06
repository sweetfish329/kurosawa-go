package main

import (
	"context"
	"log"
	"time"

	"kurosawa-go/pkg/editor"
	"kurosawa-go/pkg/recorder"
)

func main() {
	// 画面録画の例
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := recorder.NewOptions()
	opts.Area = "1280x720"
	if err := recorder.Record(ctx, "screen.mp4", opts); err != nil {
		log.Fatal(err)
	}

	// 動画編集の例
	editor := editor.New("input.mp4")
	err := editor.
		Resize(1280, 720).
		Trim(0, 30*time.Second).
		Output("output.mp4").
		Process(context.Background())

	if err != nil {
		log.Fatal(err)
	}
}
