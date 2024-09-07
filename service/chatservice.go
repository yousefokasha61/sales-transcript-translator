package service

import (
	"bytes"
	"chat/ctx"
	"chat/model"
	httpClient "chat/pkg/http/client"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"
)

type TranscriptTranslatorService struct {
	logger     *logrus.Logger
	httpClient *httpClient.HTTPClient
	openAIKey  string
}

func NewTranscriptTranslatorService(serviceContext ctx.ServiceContext) *TranscriptTranslatorService {
	return &TranscriptTranslatorService{
		httpClient: serviceContext.HTTPClient(),
		logger:     serviceContext.Logger(),
		openAIKey:  serviceContext.Conf().OpenAIKey(),
	}
}

func (s *TranscriptTranslatorService) Translate(req *model.TranslateTranscriptRequest) (*model.TranslateTranscriptResponse, error) {
	chunkedTranscript := chunkTranscript(req.Transcript)

	response := &model.TranslateTranscriptResponse{
		TranslatedTranscript: make([]model.TranscriptSegment, 0),
	}

	for _, chunk := range chunkedTranscript {
		translatedSegments, err := s.translateBatch(chunk)
		if err != nil {
			return nil, err
		}
		response.TranslatedTranscript = append(response.TranslatedTranscript, translatedSegments...)
	}

	return response, nil
}

func (s *TranscriptTranslatorService) translateBatch(batch []model.TranscriptSegment) ([]model.TranscriptSegment, error) {
	var batchText strings.Builder
	for _, segment := range batch {
		batchText.WriteString(segment.Sentence)
		batchText.WriteString("\n")
	}

	translatedText, err := s.translateToEnglish(batchText.String())
	if err != nil {
		return nil, err
	}

	translatedSentences := strings.Split(translatedText, "\n")
	translatedBatch := make([]model.TranscriptSegment, len(batch))

	for i, segment := range batch {
		translatedBatch[i] = model.TranscriptSegment{
			Speaker:  segment.Speaker,
			Time:     segment.Time,
			Sentence: strings.TrimSpace(translatedSentences[i]),
		}
	}

	return translatedBatch, nil
}

func (s *TranscriptTranslatorService) translateToEnglish(text string) (string, error) {

	requestBody := model.OpenAIRequest{
		Model: "gpt-4o-mini",
		Messages: []model.OpenAIRequestMessage{
			{
				Role:    "system",
				Content: "You are a translator. Translate the given Arabic text to English. If the text is already in English, return it as is. Maintain the original structure of the input, including line breaks."},
			{
				Role:    "user",
				Content: text,
			},
		},
		Temperature: 0.7,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.openAIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error while closing")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var openAIResp model.OpenAIResponse
	err = json.Unmarshal(body, &openAIResp)
	if err != nil {
		return "", err
	}

	if len(openAIResp.Choices) == 0 {
		return "", fmt.Errorf("no translation provided by OpenAI")
	}

	return openAIResp.Choices[0].Message.Content, nil
}

func countTokens(text string) int {
	// Naive token count, usually based on words or characters
	return utf8.RuneCountInString(text) / 4 // Rough estimate
}

func chunkTranscript(transcript []model.TranscriptSegment) [][]model.TranscriptSegment {
	var chunks [][]model.TranscriptSegment
	var currentChunk []model.TranscriptSegment
	currentTokens := 0

	for _, segment := range transcript {
		segmentTokens := countTokens(segment.Sentence)
		if currentTokens+segmentTokens > model.MaxChunkSize {
			// Add current chunk to chunks and start a new one
			chunks = append(chunks, currentChunk)
			currentChunk = []model.TranscriptSegment{}
			currentTokens = 0
		}
		currentChunk = append(currentChunk, segment)
		currentTokens += segmentTokens
	}

	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
