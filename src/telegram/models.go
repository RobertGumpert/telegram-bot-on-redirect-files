package telegram

type jsonModelSubscribeWebhook struct {
	URL string `json:"url"`
}

type (
	ChatType string
	LangCode string
)

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
	//
	LangRU LangCode = "ru"
	LangEN LangCode = "en"
)

type jsonModelUser struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	IsBot     bool     `json:"is_bot"`
	LangCode  LangCode `json:"language_code"`
}

type jsonModelChat struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	ChatType  ChatType `json:"type"`
}

type jsonModelPhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	FileSize     int64  `json:"file_size"`
}

type jsonModelAudio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int64      `json:"duration"`
	Performer    string     `json:"performer"`
	Title        string     `json:"title"`
	MimeType     string     `json:"mime_type"`
	FileSize     int64      `json:"file_size"`
	Thumb        *jsonModelPhotoSize `json:"thumb"`
}

type jsonModelDocument struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *jsonModelPhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int64      `json:"file_size"`
}

type jsonModelVideo struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	Duration     int64      `json:"duration"`
	Thumb        *jsonModelPhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int64      `json:"file_size"`
}

type jsonModelMessage struct {
	From      *jsonModelUser `json:"from"`
	Chat      *jsonModelChat `json:"chat"`
	MessageID int   `json:"message_id"`
	Date      int64 `json:"date"`
	//
	// Content
	//
	Text     string       `json:"text"`
	Photo    []*jsonModelPhotoSize `json:"photo"`
	Audio    *jsonModelAudio       `json:"audio"`
	Video    *jsonModelVideo       `json:"video"`
	Document *jsonModelDocument    `json:"document"`
}

type ClientMessage struct {
	UpdateID int64    `json:"update_id"`
	Message  *jsonModelMessage `json:"message"`
}
