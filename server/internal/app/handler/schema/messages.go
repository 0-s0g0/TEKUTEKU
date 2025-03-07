package schema

type MessagePOSTRequest struct {
	Message string `json:"text"`
	School  int    `json:"school"`
}

type MessagePOSTResponse struct {
}

type MessageGETResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	ID        string    `json:"id"`
	Message   string    `json:"text"`
	Likes     int       `json:"likes"`
	X         int       `json:"X"`
	Y         int       `json:"Y"`
	FloatTime float32   `json:"float_time"`
	CreatedAt string    `json:"created_at"`
	School    int       `json:"school"`
	Reply     []Message `json:"reply"`
}

type ReplyPOSTRequest struct {
	Message string `json:"text"`
	To      string `json:"to"`
	School  int    `json:"school"`
}
