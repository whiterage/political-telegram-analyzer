package app

import (
	"log/slog"
	"os"

	"sofiasoft/internal/config"
	"sofiasoft/internal/pipeline"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load("config.yaml")
	if err != nil {
		logger.Error("failed to load config", "error", err)
		return
	}

	p := pipeline.New(cfg, logger)

	if err := p.Run(); err != nil {
		logger.Error("pipeline failed", "error", err)
		return
	}
}
