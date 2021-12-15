// https://github.com/SlitiBrahim/go-telegram-notifier

// ********** go-telegram-notifier/main.go
package main

import (
	"go-telegram-notifier/api"
)

func main() {
	api.Start()
}

// ********** go-telegram-notifier/config/config.go
package config

import (
	"fmt"
	"os"
)

// TODO: better solution ?
var Config = make(map[string]interface{})

func init() {
	Config["APP_PORT"] = os.Getenv("APP_PORT")
	Config["TG_CHAT_ID"] = os.Getenv("TG_CHAT_ID")
	Config["TG_BOT_TOKEN"] = os.Getenv("TG_BOT_TOKEN")
	Config["TOKEN"] = os.Getenv("TOKEN")
	Config["TG_API_BOT_BASE_URL"] = fmt.Sprintf("https://api.telegram.org/bot%s/",
        Config["TG_BOT_TOKEN"])
}

// ********** go-telegram-notifier/helper/error.go
package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SendApiError(w http.ResponseWriter, err error, httpStatus int) {
	if err != nil {
		httpErr := map[string]interface{}{
			"error": err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)

		json.NewEncoder(w).Encode(httpErr)
	}
}

// ********** go-telegram-notifier/api/Notification.go
package api

type Notification struct {
	Message string `json:"message"`
}

// ********** go-telegram-notifier/api/api.go
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go-telegram-notifier/config"
	"go-telegram-notifier/helper"
	"log"
	"net/http"
)

func authenticatedReq(r *http.Request) bool {
	return r.Header.Get("token") == config.Config["TOKEN"]
}

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if authenticatedReq(r) == false {
		helper.SendApiError(w, errors.New("invalid token"), http.StatusForbidden)
		return
	}

	var notification Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		helper.SendApiError(w,
            errors.New("invalid request body: cannot parse body to Notification object"),
            http.StatusBadRequest)
		return
	}

	if notification.Message == "" {
		helper.SendApiError(w, errors.New("empty message passed"), http.StatusBadRequest)
		return
	}

	msg := Message{
		ChatID: config.Config["TG_CHAT_ID"].(string),
		Text:   notification.Message,
	}

	telegramResponse, err := sendMessage(msg)
	helper.SendApiError(w, err, http.StatusInternalServerError)

	if telegramResponse.StatusCode == http.StatusOK {
		res := map[string]interface{}{
			"message": "Notification has been sent.",
		}

		err = ReturnResponse(w, res, http.StatusOK)
		helper.SendApiError(w, err, http.StatusInternalServerError)
	} else {
		res := map[string]interface{}{
			"message": "Notification cannot be sent.",
			"error":   err.Error(),
		}

		err = ReturnResponse(w, res, http.StatusBadRequest)
		helper.SendApiError(w, err, http.StatusInternalServerError)
	}
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/send-notification", sendNotificationHandler).Methods("POST")

	log.Printf("Listening on localhost:%v\n", config.Config["APP_PORT"])
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config["APP_PORT"]), router)
	helper.FailOnError(err)
}

func ReturnResponse(w http.ResponseWriter, res map[string]interface{}, httpStatus int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	err := json.NewEncoder(w).Encode(res)

	return err
}

// ********** go-telegram-notifier/api/telegram.go
package api

import (
	"bytes"
	"encoding/json"
	"go-telegram-notifier/config"
	"go-telegram-notifier/helper"
	"net/http"
	"net/url"
	"path"
)

type Message struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func getSendMessageURL() string {
	baseURL, err := url.Parse(config.Config["TG_API_BOT_BASE_URL"].(string))
	helper.FailOnError(err)

	baseURL.Path = path.Join(baseURL.Path, "/sendMessage")

	return baseURL.String()
}

func sendMessage(message Message) (*http.Response, error) {
	body, err := json.Marshal(message)
	helper.FailOnError(err)

	return http.Post(getSendMessageURL(), "application/json", bytes.NewReader(body))
}
