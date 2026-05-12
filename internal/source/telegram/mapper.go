package telegram

// Telegram source must map raw Telegram messages into domain.Post.
//
// Required fields:
// - ID
// - ChannelName
// - ChannelUsername
// - ChannelCategory
// - ChannelActorType
// - URL
// - Text
// - PublishedAt
// - Views
// - Forwards
// - CommentsCount
// - HasMedia
// - MediaType
// - Reactions
//
// TODO: implement MTProto message mapper into domain.Post.
