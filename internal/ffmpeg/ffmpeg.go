package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Command represents an FFmpeg command builder
type Command struct {
	inputs  []string
	outputs []string
	options []string
	filters []string
}

// NewCommand creates a new FFmpeg command builder
func NewCommand() *Command {
	return &Command{
		inputs:  make([]string, 0),
		outputs: make([]string, 0),
		options: make([]string, 0),
		filters: make([]string, 0),
	}
}

// WithInput adds an input file to the command
func (c *Command) WithInput(input string) *Command {
	c.inputs = append(c.inputs, input)
	return c
}

// WithOutput adds an output file to the command
func (c *Command) WithOutput(output string) *Command {
	c.outputs = append(c.outputs, output)
	return c
}

// WithOption adds an FFmpeg option to the command
func (c *Command) WithOption(name, value string) *Command {
	c.options = append(c.options, name, value)
	return c
}

// WithFilter adds a filter to the command
func (c *Command) WithFilter(filter string) *Command {
	c.filters = append(c.filters, filter)
	return c
}

// buildArgs builds the FFmpeg command arguments
func (c *Command) buildArgs() []string {
	args := make([]string, 0)

	// Add inputs
	for _, input := range c.inputs {
		args = append(args, "-i", input)
	}

	// Add options
	args = append(args, c.options...)

	// Add filters if any
	if len(c.filters) > 0 {
		filterStr := strings.Join(c.filters, ",")
		args = append(args, "-vf", filterStr)
	}

	// Add outputs
	args = append(args, c.outputs...)

	return args
}

// Run executes the FFmpeg command
func (c *Command) Run() error {
	args := c.buildArgs()
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg command failed: %w", err)
	}

	return nil
}

// StartProcess starts an FFmpeg process and returns the command
func (c *Command) StartProcess() (*exec.Cmd, error) {
	args := c.buildArgs()
	cmd := exec.Command("ffmpeg", args...)

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start ffmpeg process: %w", err)
	}

	return cmd, nil
}
