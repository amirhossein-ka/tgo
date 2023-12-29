package tgo

import (
	"bytes"
	"encoding/json"
)

// Sendable is an interface that represents any object that can be sent using an API client.
type Sendable interface {
	// GetChatID returns the chat ID associated with the sendable object.
	GetChatID() ChatID

	// SetChatID sets the chat ID for the sendable object.
	SetChatID(id int64)

	// Send sends the sendable object using the provided API client.
	// It returns the sent message and any error encountered.
	Send(api *API) (*Message, error)
}

// Replyable is an interface that represents any object that can be replied to.
type Replyable interface {
	Sendable
	SetReplyToMessageId(id int64)
}

// ParseModeSettable is an interface that represents any object that can have its ParseMode set
// Or in other words, messages with captions.
type ParseModeSettable interface {
	GetParseMode() ParseMode
	SetParseMode(mode ParseMode)
}

// ParseMode is a type that represents the parse mode of a message. eg. Markdown, HTML, etc.
type ParseMode string

const (
	// does not parse the message
	ParseModeNone ParseMode = ""
	// parses the message as Markdown
	ParseModeMarkdown ParseMode = "Markdown"
	// parses the message as Markdown but using telegram's V2 markdown
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	// parses the message as HTML
	ParseModeHTML ParseMode = "HTML"
)

type Username string

func (Username) IsChatID() {}

type ID int64

func (ID) IsChatID() {}

func unmarshalChatID(data []byte) (ChatID, error) {
	if bytes.HasPrefix(data, []byte("\"")) {
		var username Username
		err := json.Unmarshal(data, &username)
		return username, err
	} else {
		var id ID
		err := json.Unmarshal(data, &id)
		return id, err
	}
}

func (x *SendAnimation) GetChatID() ChatID { return x.ChatId }
func (x *SendAudio) GetChatID() ChatID     { return x.ChatId }
func (x *SendContact) GetChatID() ChatID   { return x.ChatId }
func (x *SendDice) GetChatID() ChatID      { return x.ChatId }
func (x *SendDocument) GetChatID() ChatID  { return x.ChatId }
func (x *SendGame) GetChatID() ChatID      { return ID(x.ChatId) }
func (x *SendInvoice) GetChatID() ChatID   { return x.ChatId }
func (x *SendLocation) GetChatID() ChatID  { return x.ChatId }
func (x *SendMessage) GetChatID() ChatID   { return x.ChatId }
func (x *SendPhoto) GetChatID() ChatID     { return x.ChatId }
func (x *SendPoll) GetChatID() ChatID      { return x.ChatId }
func (x *SendSticker) GetChatID() ChatID   { return x.ChatId }
func (x *SendVenue) GetChatID() ChatID     { return x.ChatId }
func (x *SendVideo) GetChatID() ChatID     { return x.ChatId }
func (x *SendVideoNote) GetChatID() ChatID { return x.ChatId }
func (x *SendVoice) GetChatID() ChatID     { return x.ChatId }

func (x *SendAnimation) SetChatID(id int64) { x.ChatId = ID(id) }
func (x *SendAudio) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendContact) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendDice) SetChatID(id int64)      { x.ChatId = ID(id) }
func (x *SendDocument) SetChatID(id int64)  { x.ChatId = ID(id) }
func (x *SendGame) SetChatID(id int64)      { x.ChatId = id }
func (x *SendInvoice) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendLocation) SetChatID(id int64)  { x.ChatId = ID(id) }
func (x *SendMessage) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendPhoto) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendPoll) SetChatID(id int64)      { x.ChatId = ID(id) }
func (x *SendSticker) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendVenue) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendVideo) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendVideoNote) SetChatID(id int64) { x.ChatId = ID(id) }
func (x *SendVoice) SetChatID(id int64)     { x.ChatId = ID(id) }

func (x *SendAnimation) Send(api *API) (*Message, error) { return api.SendAnimation(x) }
func (x *SendAudio) Send(api *API) (*Message, error)     { return api.SendAudio(x) }
func (x *SendContact) Send(api *API) (*Message, error)   { return api.SendContact(x) }
func (x *SendDice) Send(api *API) (*Message, error)      { return api.SendDice(x) }
func (x *SendDocument) Send(api *API) (*Message, error)  { return api.SendDocument(x) }
func (x *SendGame) Send(api *API)                        { api.SendGame(x) }
func (x *SendInvoice) Send(api *API) (*Message, error)   { return api.SendInvoice(x) }
func (x *SendLocation) Send(api *API) (*Message, error)  { return api.SendLocation(x) }
func (x *SendMessage) Send(api *API) (*Message, error)   { return api.SendMessage(x) }
func (x *SendPhoto) Send(api *API) (*Message, error)     { return api.SendPhoto(x) }
func (x *SendPoll) Send(api *API) (*Message, error)      { return api.SendPoll(x) }
func (x *SendSticker) Send(api *API) (*Message, error)   { return api.SendSticker(x) }
func (x *SendVenue) Send(api *API) (*Message, error)     { return api.SendVenue(x) }
func (x *SendVideo) Send(api *API) (*Message, error)     { return api.SendVideo(x) }
func (x *SendVideoNote) Send(api *API) (*Message, error) { return api.SendVideoNote(x) }
func (x *SendVoice) Send(api *API) (*Message, error)     { return api.SendVoice(x) }

func (x *SendAnimation) GetParseMode() ParseMode { return x.ParseMode }
func (x *SendAudio) GetParseMode() ParseMode     { return x.ParseMode }
func (x *SendDocument) GetParseMode() ParseMode  { return x.ParseMode }
func (x *SendMessage) GetParseMode() ParseMode   { return x.ParseMode }
func (x *SendPhoto) GetParseMode() ParseMode     { return x.ParseMode }
func (x *SendVideo) GetParseMode() ParseMode     { return x.ParseMode }
func (x *SendVoice) GetParseMode() ParseMode     { return x.ParseMode }

func (x *SendAnimation) SetParseMode(mode ParseMode) { x.ParseMode = mode }
func (x *SendAudio) SetParseMode(mode ParseMode)     { x.ParseMode = mode }
func (x *SendDocument) SetParseMode(mode ParseMode)  { x.ParseMode = mode }
func (x *SendMessage) SetParseMode(mode ParseMode)   { x.ParseMode = mode }
func (x *SendPhoto) SetParseMode(mode ParseMode)     { x.ParseMode = mode }
func (x *SendVideo) SetParseMode(mode ParseMode)     { x.ParseMode = mode }
func (x *SendVoice) SetParseMode(mode ParseMode)     { x.ParseMode = mode }

func (x *SendAnimation) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendAudio) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendContact) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendDice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendDocument) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendGame) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendInvoice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendLocation) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendMessage) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendPhoto) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendPoll) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendSticker) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendVenue) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendVideo) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendVideoNote) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
func (x *SendVoice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// Send sends a message with the preferred ParseMode.
func (b *Bot) Send(msg Sendable) (*Message, error) {
	if x, ok := msg.(ParseModeSettable); ok {
		if x.GetParseMode() == ParseModeNone {
			x.SetParseMode(b.DefaultParseMode)
		}
	}

	return msg.Send(b.API)
}
