package config

import "errors"

type Config struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
	TopLimit   int    `yaml:"top_limit"`
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
	return nil

}
