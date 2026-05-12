package config

import "errors"

type Config struct {
	InputFile                string `yaml:"input_file"`
	OutputFile               string `yaml:"output_file"`
	SummaryOutputFile        string `yaml:"summary_output_file"`
	TopLimit                 int    `yaml:"top_limit"`
	ChannelSummaryOutputFile string `yaml:"channel_summary_output_file"`
}

func (c Config) Validate() error {
	if c.InputFile == "" {
		return errors.New("input_file is required")
	}
	if c.OutputFile == "" {
		return errors.New("output_file is required")
	}
	if c.TopLimit <= 0 {
		return errors.New("top_limit must be greater than zero")
	}
	if c.SummaryOutputFile == "" {
		return errors.New("summary_output_file is required")
	}
	if c.ChannelSummaryOutputFile == "" {
		return errors.New("channel_summary_output_file is required")
	}
	return nil

}
