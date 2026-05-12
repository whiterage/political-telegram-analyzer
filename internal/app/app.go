package app

import (
	"fmt"
	"sofiasoft/internal/emotion/rulebased"
	"sofiasoft/internal/summary"

	"sofiasoft/internal/analysis"
	"sofiasoft/internal/config"
	"sofiasoft/internal/export"
	"sofiasoft/internal/input"
)

func Run() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		return
	}

	posts, err := input.ReadPostsJSON(cfg.InputFile)
	if err != nil {
		fmt.Printf("failed to read posts json: %v\n", err)
		return
	}

	classifier := rulebased.NewClassifier()

	analyzedPosts := analysis.AnalyzePosts(posts, classifier)

	analysis.SortByERR(analyzedPosts)

	topPosts := analysis.TopPosts(analyzedPosts, cfg.TopLimit)

	if err := export.WriteAnalyzedPostsCSV(cfg.OutputFile, topPosts); err != nil {
		fmt.Printf("failed to write csv: %v\n", err)
		return
	}

	fmt.Printf("CSV exported: %s\n", cfg.OutputFile)

	summaries := summary.BuildEmotionSummary(topPosts)

	if err := export.WriteEmotionSummaryCSV(cfg.SummaryOutputFile, summaries); err != nil {
		fmt.Printf("failed to write summary csv: %v\n", err)
		return
	}

	fmt.Printf("Summary CSV exported: %s\n", cfg.SummaryOutputFile)

	channelSummaries := summary.BuildChannelSummary(topPosts)

	if err := export.WriteChannelSummaryCSV(cfg.ChannelSummaryOutputFile, channelSummaries); err != nil {
		fmt.Printf("failed to write channel summary csv: %v\n", err)
		return
	}

	fmt.Printf("Channel summary CSV exported: %s\n", cfg.ChannelSummaryOutputFile)

	for _, post := range topPosts {
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
