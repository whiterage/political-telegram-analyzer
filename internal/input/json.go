package input

import (
	"encoding/json"
	"os"
	"sofiasoft/internal/domain"
)

func ReadPostsJSON(filename string) ([]domain.Post, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var posts []domain.Post

	if err := json.Unmarshal(data, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
