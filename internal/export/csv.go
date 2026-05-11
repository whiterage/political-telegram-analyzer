package export

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"sofiasoft/internal/domain"
)

func WriteAnalyzedPostsCSV(filename string, posts []domain.AnalyzedPost) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"id",
		"channel_name",
		"text",
		"views",
		"total_reactions",
		"dominant_emoji",
		"reaction_emotion",
		"comments",
		"forwards",
		"err",
		"format",
		"emotion",
		"frame",
		"confidence",
		"method",
		"reason",
		"markers",
	}

	if err := writer.Write(header); err != nil {
		return err
	}

	for _, post := range posts {
		row := []string{
			strconv.FormatInt(post.Post.ID, 10),
			post.Post.ChannelName,
			post.Post.Text,
			strconv.Itoa(post.Post.Views),
			strconv.Itoa(post.TotalReactions),
			post.ReactionEmotion.DominantEmoji,
			post.ReactionEmotion.DominantEmotion,
			strconv.Itoa(post.Post.CommentsCount),
			strconv.Itoa(post.Post.Forwards),
			strconv.FormatFloat(post.ERR, 'f', 2, 64),
			post.FormatType,
			post.Emotion.Emotion,
			post.Emotion.Frame,
			strconv.FormatFloat(post.Emotion.Confidence, 'f', 2, 64),
			post.Emotion.Method,
			post.Emotion.Reason,
			strings.Join(post.Emotion.Markers, "|"),
		}

		if err := writer.Write(row); err != nil {
			return err
		}
	}

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
