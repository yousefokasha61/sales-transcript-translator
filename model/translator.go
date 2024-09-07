package model

const (
	InvalidRequest = "invalid_request"
	MaxChunkSize   = 4000 // Maximum characters per chunk
)

type HttpMainResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type EmptyData struct{}

type TranslateTranscriptRequest struct {
	Transcript []TranscriptSegment `json:"transcript"`
}

type TranslateTranscriptResponse struct {
	TranslatedTranscript []TranscriptSegment `json:"translatedTranscript"`
}

type TranscriptSegment struct {
	Speaker  string `json:"speaker"`
	Time     string `json:"time"`
	Sentence string `json:"sentence"`
}

type OpenAIRequest struct {
	Model       string                 `json:"model"`
	Messages    []OpenAIRequestMessage `json:"messages"`
	Temperature float64                `json:"temperature"`
}

type OpenAIRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content string `json:"content"`
}
