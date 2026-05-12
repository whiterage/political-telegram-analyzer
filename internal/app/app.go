package app

import (
	"fmt"

	"sofiasoft/internal/config"
	"sofiasoft/internal/pipeline"
)

func Run() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		return
	}

	p := pipeline.New(cfg)

	if err := p.Run(); err != nil {
		fmt.Printf("pipeline failed: %v\n", err)
		return
	}
}
