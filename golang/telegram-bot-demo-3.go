// https://dev.to/talentlessguy/create-and-deploy-golang-telegram-bot-2dl0

package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hi!")
		// Getting user input
		b.Send(m.Sender, "You entered "+m.Text)
		// Any text after the command will also be shown.
		// For text without a command, use Payload:
		b.Send(m.Sender, "You entered "+m.Payload)
	})

	// Buttons
	inlineBtn1 := tb.InlineButton{Unique: "moon", Text: "Moon 🌚"}
	inlineBtn2 := tb.InlineButton{Unique: "sun", Text: "Sun 🌞"}

	b.Handle(&inlineBtn1, func(c *tb.Callback) {
		// Required for proper work
		b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		// Send messages here
		b.Send(c.Sender, "Moon says 'Hi'!")
	})

	// Let's do the same for second button:
	b.Handle(&inlineBtn2, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		b.Send(c.Sender, "Sun says 'Hi'!")
	})

	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn1, inlineBtn2},
	}

	b.Handle("/pick_time", func(m *tb.Message) {
		b.Send(m.Sender, "Day or night, you choose",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
	})

	b.Start()
}
