package domain

type Reaction struct {
	Emoji string `json:"emoji"`
	Count int    `json:"count"`
}

type Post struct {
	ID              int64      `json:"id"`
	ChannelName     string     `json:"channel_name"`
	ChannelUsername string     `json:"channel_username"`
	URL             string     `json:"url"`
	Text            string     `json:"text"`
	PublishedAt     string     `json:"published_at"`
	Views           int        `json:"views"`
	Forwards        int        `json:"forwards"`
	CommentsCount   int        `json:"comments_count"`
	HasMedia        bool       `json:"has_media"`
	MediaType       string     `json:"media_type"`
	Reactions       []Reaction `json:"reactions"`
}
