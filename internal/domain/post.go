package domain

type Reaction struct {
	Emoji string `json:"emoji"`
	Count int    `json:"count"`
}

type Post struct {
	ID            int64      `json:"id"`
	ChannelName   string     `json:"channel_name"`
	Text          string     `json:"text"`
	Views         int        `json:"views"`
	Forwards      int        `json:"forwards"`
	CommentsCount int        `json:"comments_count"`
	Reactions     []Reaction `json:"reactions"`
}
