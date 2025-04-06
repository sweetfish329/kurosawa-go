package pipeline

import "context"

type Stage interface {
	Process(ctx context.Context, in <-chan *media.Frame, out chan<- *media.Frame) error
}

type Pipeline struct {
	stages []Stage
}

func (p *Pipeline) AddStage(stage Stage) *Pipeline {
	p.stages = append(p.stages, stage)
	return p
}

func (p *Pipeline) Run(ctx context.Context) error {
	// ...pipeline implementation
}
