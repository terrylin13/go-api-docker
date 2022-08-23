package ws

// message type
const (
	SystemType = iota
	TextType
	MediaType
)

type Message struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}
