package app

import (
	"fmt"
	"sofiasoft/src/internal/emotion/mock"
	"sofiasoft/src/internal/export"

	"sofiasoft/src/internal/analysis"
	"sofiasoft/src/internal/domain"
)

func Run() {
	posts := []domain.Post{
		{
			ID:            1,
			ChannelName:   "Тестовый канал",
			Text:          "Они снова врут. Это угроза для всех. Не молчи!",
			Views:         10000,
			Forwards:      120,
			CommentsCount: 80,
			Reactions: []domain.Reaction{
				{Emoji: "😡", Count: 500},
				{Emoji: "🔥", Count: 300},
				{Emoji: "👍", Count: 200},
			},
		},
		{
			ID:            2,
			ChannelName:   "Тестовый канал",
			Text:          "Наши герои снова добились победы. Спасибо всем, кто поддерживает!",
			Views:         20000,
			Forwards:      300,
			CommentsCount: 150,
			Reactions: []domain.Reaction{
				{Emoji: "🔥", Count: 1200},
				{Emoji: "👍", Count: 900},
				{Emoji: "❤️", Count: 400},
			},
		},
		{
			ID:            3,
			ChannelName:   "Тестовый канал",
			Text:          "Сегодня опубликован подробный аналитический разбор ситуации. Эксперты отмечают несколько факторов, которые могут повлиять на дальнейшее развитие событий.",
			Views:         5000,
			Forwards:      20,
			CommentsCount: 15,
			Reactions: []domain.Reaction{
				{Emoji: "👍", Count: 100},
			},
		},
	}

	classifier := mock.NewClassifier()
	analyzedPosts := analysis.AnalyzePosts(posts, classifier)
	analysis.SortByERR(analyzedPosts)
	topPosts := analysis.TopPosts(analyzedPosts, 2)
	if err := export.WriteAnalyzedPostsCSV("out_posts.csv", topPosts); err != nil {
		fmt.Printf("failed to write csv: %v\n", err)
		return
	}
	fmt.Println("CSV exported: out_posts.csv")

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
	}
}
