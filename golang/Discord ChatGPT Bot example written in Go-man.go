// Package main  https://gist.github.com/haoel/0d5440e953743a1b359300fff408311c
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

// Token is the token for the discord bot and chatgpt
type Token struct {
	Discord string `yaml:"discord"`
	ChatGPT string `yaml:"chatgpt"`
}

var token Token

// ReadConfig reads the config file and unmarshals it into the config variable
func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = yaml.Unmarshal(file, &token)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Config file read successfully!")

	return nil

}

// Start starts the bot
func Start() {
	dg, err := discordgo.New("Bot " + token.Discord)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageHandler)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	select {
	case <-done:
		fmt.Println("Received the exit signal, exiting...")
	}
	// Cleanly close down the Discord session.
	fmt.Println("Closing Discord session...")
	dg.Close()

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't mention the bot
	mentiond := false
	for _, u := range m.Mentions {
		if u.ID == s.State.User.ID {
			mentiond = true
			break
		}
	}
	if !mentiond {
		fmt.Printf("Not mentioned in message: [%s] %s\n", m.Author.Username, m.Content)
		return
	}

	message := fmt.Sprintf("Message: %s, Author: %s", m.Content, m.Author.Username)

	fmt.Println(message)
	ChatGPTResponse, err := callChatGPT(m.Content)
	if err != nil {
		fmt.Println(err.Error())
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	//s.ChannelMessageSend(m.ChannelID, ChatGPTResponse)
	s.ChannelMessageSendReply(m.ChannelID, ChatGPTResponse, m.Reference())

}

// ChatGPTResponse is the response from the chatgpt api
type ChatGPTResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Model  string `json:"model"`
	Usage  struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

func callChatGPT(msg string) (string, error) {
	api := "https://api.openai.com/v1/chat/completions"
	body := []byte(`{
		"model": "gpt-3.5-turbo",
		"messages": [
		  {
			"role": "user",
			"content": "` + msg + `"
		  }
		]
	  }`)

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error creating request", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.ChatGPT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return "", err
	}

	chatGPTData := ChatGPTResponse{}
	err = json.Unmarshal(body, &chatGPTData)
	if err != nil {
		fmt.Println("Error unmarshalling response", err)
		return "", err
	}
	return chatGPTData.Choices[0].Message.Content, nil
}

func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()
}
