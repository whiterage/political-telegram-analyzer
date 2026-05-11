package domain

type Reaction struct {
	Emoji string
	Count int
}

type Post struct {
	ID            int64
	ChannelName   string
	Text          string
	Views         int
	Forwards      int
	CommentsCount int
	Reactions     []Reaction
}
