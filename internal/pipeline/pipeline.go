package pipeline

import (
	"fmt"
	"log/slog"
	"sofiasoft/internal/analysis"
	"sofiasoft/internal/config"
	"sofiasoft/internal/domain"
	"sofiasoft/internal/emotion/rulebased"
	"sofiasoft/internal/export"
	"sofiasoft/internal/source"
	jsonsource "sofiasoft/internal/source/json"
	"sofiasoft/internal/summary"
)

type Pipeline struct {
	cfg    config.Config
	logger *slog.Logger
}

func New(cfg config.Config, logger *slog.Logger) *Pipeline {
	if logger == nil {
		logger = slog.Default()
	}

	return &Pipeline{
		cfg:    cfg,
		logger: logger,
	}
}

func (p *Pipeline) Run() error {
	posts, err := p.loadPosts()
	if err != nil {
		return err
	}

	classifier := rulebased.NewClassifier()

	analyzedPosts := analysis.AnalyzePosts(posts, classifier)

	analysis.SortByERR(analyzedPosts)

	topPosts := analysis.TopPosts(analyzedPosts, p.cfg.TopLimit)

	if err := export.WriteAnalyzedPostsCSV(p.cfg.OutputFile, topPosts); err != nil {
		return fmt.Errorf("write analyzed posts csv: %w", err)
	}

	p.logger.Info("analyzed posts csv exported", "file", p.cfg.OutputFile)

	summaries := summary.BuildEmotionSummary(topPosts)

	if err := export.WriteEmotionSummaryCSV(p.cfg.SummaryOutputFile, summaries); err != nil {
		return fmt.Errorf("write emotion summary csv: %w", err)
	}

	p.logger.Info("emotion summary csv exported", "file", p.cfg.SummaryOutputFile)

	channelSummaries := summary.BuildChannelSummary(topPosts)

	if err := export.WriteChannelSummaryCSV(p.cfg.ChannelSummaryOutputFile, channelSummaries); err != nil {
		return fmt.Errorf("write channel summary csv: %w", err)
	}

	p.logger.Info(
		"channel summary csv exported",
		"file",
		p.cfg.ChannelSummaryOutputFile,
	)

	p.logTopPosts(topPosts)

	return nil
}

func (p *Pipeline) logTopPosts(posts []domain.AnalyzedPost) {
	for _, post := range posts {
		p.logger.Info(
			"top post analyzed",
			"id",
			post.Post.ID,
			"channel_name",
			post.Post.ChannelName,
			"text",
			post.Post.Text,
			"views",
			post.Post.Views,
			"total_reactions",
			post.TotalReactions,
			"err",
			post.ERR,
			"format",
			post.FormatType,
			"emotion",
			post.Emotion.Emotion,
			"frame",
			post.Emotion.Frame,
			"confidence",
			post.Emotion.Confidence,
			"method",
			post.Emotion.Method,
			"reason",
			post.Emotion.Reason,
			"markers",
			post.Emotion.Markers,
			"dominant_emoji",
			post.ReactionEmotion.DominantEmoji,
			"reaction_emotion",
			post.ReactionEmotion.DominantEmotion,
		)
	}
}

func (p *Pipeline) loadPosts() ([]domain.Post, error) {
	src, err := p.newSource()
	if err != nil {
		return nil, err
	}

	posts, err := src.LoadPosts()
	if err != nil {
		return nil, fmt.Errorf("load posts from %s source: %w", p.cfg.Source, err)
	}

	return posts, nil
}

func (p *Pipeline) newSource() (source.Source, error) {
	switch p.cfg.Source {
	case "json":
		return jsonsource.New(p.cfg.InputFile), nil

	case "telegram":
		return nil, fmt.Errorf("telegram source is not implemented yet")

	default:
		return nil, fmt.Errorf("unsupported source: %s", p.cfg.Source)
	}
}
