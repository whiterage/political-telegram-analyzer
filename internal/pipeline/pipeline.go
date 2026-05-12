package pipeline

import (
	"fmt"
	"sofiasoft/internal/analysis"
	"sofiasoft/internal/config"
	"sofiasoft/internal/domain"
	"sofiasoft/internal/emotion/rulebased"
	"sofiasoft/internal/export"
	"sofiasoft/internal/input"
	"sofiasoft/internal/summary"
)

type Pipeline struct {
	cfg config.Config
}

func New(cfg config.Config) *Pipeline {
	return &Pipeline{
		cfg: cfg,
	}
}

func (p *Pipeline) Run() error {
	posts, err := input.ReadPostsJSON(p.cfg.InputFile)
	if err != nil {
		return fmt.Errorf("read posts json: %w", err)
	}

	classifier := rulebased.NewClassifier()

	analyzedPosts := analysis.AnalyzePosts(posts, classifier)

	analysis.SortByERR(analyzedPosts)

	topPosts := analysis.TopPosts(analyzedPosts, p.cfg.TopLimit)

	if err := export.WriteAnalyzedPostsCSV(p.cfg.OutputFile, topPosts); err != nil {
		return fmt.Errorf("write analyzed posts csv: %w", err)
	}

	fmt.Printf("CSV exported: %s\n", p.cfg.OutputFile)

	summaries := summary.BuildEmotionSummary(topPosts)

	if err := export.WriteEmotionSummaryCSV(p.cfg.SummaryOutputFile, summaries); err != nil {
		return fmt.Errorf("write emotion summary csv: %w", err)
	}

	fmt.Printf("Summary CSV exported: %s\n", p.cfg.SummaryOutputFile)

	channelSummaries := summary.BuildChannelSummary(topPosts)

	if err := export.WriteChannelSummaryCSV(p.cfg.ChannelSummaryOutputFile, channelSummaries); err != nil {
		return fmt.Errorf("write channel summary csv: %w", err)
	}

	fmt.Printf("Channel summary CSV exported: %s\n", p.cfg.ChannelSummaryOutputFile)

	printTopPosts(topPosts)

	return nil
}

func printTopPosts(posts []domain.AnalyzedPost) {
	for _, post := range posts {
		fmt.Println("----------")
		fmt.Printf("ID: %d\n", post.Post.ID)
		fmt.Printf("Text: %s\n", post.Post.Text)
		fmt.Printf("Views: %d\n", post.Post.Views)
		fmt.Printf("Total reactions: %d\n", post.TotalReactions)
		fmt.Printf("ERR: %.2f%%\n", post.ERR)
		fmt.Printf("Format: %s\n", post.FormatType)
		fmt.Printf("Emotion: %s\n", post.Emotion.Emotion)
		fmt.Printf("Frame: %s\n", post.Emotion.Frame)
		fmt.Printf("Confidence: %.2f\n", post.Emotion.Confidence)
		fmt.Printf("Method: %s\n", post.Emotion.Method)
		fmt.Printf("Reason: %s\n", post.Emotion.Reason)
		fmt.Printf("Markers: %v\n", post.Emotion.Markers)
		fmt.Printf("Dominant emoji: %s\n", post.ReactionEmotion.DominantEmoji)
		fmt.Printf("Reaction emotion: %s\n", post.ReactionEmotion.DominantEmotion)
	}
}
