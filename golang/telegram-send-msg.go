package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
)

var chatID string
var apiToken string

func Init(chatId string, token string) {
	chatID = chatId
	apiToken = token
}

type Message struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(text string) (*http.Response, error) {
	msg := Message{
		ChatID: chartID,
		Text:   text,
	}

	body, err := json.Marshal(msg)
	FailOnError(err)

	// https://api.telegram.org/bot-TOKEN/sendMessage
	sendMessageURL := getSendMessageURL()
	return http.Post(sendMessageURL, "application/json", bytes.NewReader(body))
}

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getSendMessageURL() string {
	baseURL := fmt.Sprintf("https://api.telegram.org/bot%s/", apiToken)
	apiURL, err := url.Parse(baseURL)
	FailOnError(err)

	apiURL.Path = path.Join(apiURL.Path, "/sendMessage")
	return apiURL.String()
}

// for reference
func WriteResponse(w http.ResponseWriter, res map[string]interface{}, httpStatus int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(res)
}

// Example:
//
// telegram.Init("chartId", "token")
//
// resp, err := telegram.SendMessage("Hello")
// if resp.StatusCode == http.StatusOK {
//    fmt.Println("Notification has been sent")
// }
