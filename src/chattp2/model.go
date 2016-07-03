package chattp2

// Message represents a chat message
type Message struct {
	Content  string `json:"content"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}
