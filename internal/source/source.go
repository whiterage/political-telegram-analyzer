package source

import "sofiasoft/internal/domain"

type Source interface {
	LoadPosts() ([]domain.Post, error)
}
