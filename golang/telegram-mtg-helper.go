package main  // https://github.com/masticore252/telegram-mtg-helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	scryfall "github.com/BlueMonday/go-scryfall"
	env "github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

const startMessage = "Hi! I'm your Magic: the gathering bot, I work in inline mode to search for Magic: The Gathering cards in scryfall.com\n" +
	"I support advanced syntax to filter results like color, type, artist, mana value, etc\n\n" +
	"tap the button bellow to start using me\n\n" +
	"Type /help to learn more about inline bots and the advanced syntax you can use to filter your searches"

const helpMessage = "Helpful links:"

const defaultMessage = "I only work in inline mode, tab the button bellow to search \"%s\""

func main() {
	env.Load()
	bot, _ := makeBot()
	bot.Start()
}

func makeBot() (*tb.Bot, error) {
	Api := os.Getenv("TELEGRAM_API_URL")
	token := os.Getenv("TELEGRAM_TOKEN")
	poller := makePoller(os.Getenv("POLLER_MODE"))
	isVerbose, _ := strconv.ParseBool(os.Getenv("VERBOSE_OUTPUT"))

	bot, _ := tb.NewBot(tb.Settings{
		URL:       Api,
		Token:     token,
		Poller:    poller,
		Verbose:   isVerbose,
		ParseMode: tb.ModeHTML,
	})

	// Handle inline queries
	bot.Handle(tb.OnQuery, func(q *tb.Query) {

		results := tb.Results{}

		if len(q.Text) == 0 {
			return
		}

		cards, _ := cardSearch(q.Text)

		if len(cards) == 0 {
			emptyResult := &tb.ArticleResult{
				Title:       "No results",
				Description: "Your query returned no results",
				Text:        "Your query returned no results",
			}
			emptyResult.SetResultID("0")
			results = append(results, emptyResult)
		}

		for _, card := range cards {
			if isDoubleFacedLayout(card.Layout) {
				backFace := newResultFromFace(card, 0)
				frontFace := newResultFromFace(card, 1)
				results = append(results, frontFace, backFace)
			} else {
				singleResult := newResultFromCard(card)
				results = append(results, singleResult)
			}
		}

		err := bot.Answer(q, &tb.QueryResponse{
			Results:   results,
			CacheTime: 60, // a minute
		})

		if err != nil {
			log.Println(err)
		}
	})

	// Handle /start command
	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Chat, startMessage, &tb.SendOptions{
			ReplyMarkup:           makeReplyMarkupForStart(),
			DisableWebPagePreview: true,
		})
	})

	// Handle /help command
	bot.Handle("/help", func(m *tb.Message) {
		bot.Send(m.Chat, helpMessage, &tb.SendOptions{
			ReplyMarkup:           makeReplyMarkupForHelp(),
			DisableWebPagePreview: true,
		})
	})

	// Handle all other messages
	bot.Handle(tb.OnText, func(m *tb.Message) {
		msg := fmt.Sprintf(defaultMessage, m.Text)
		bot.Send(m.Chat, msg, &tb.SendOptions{
			ReplyMarkup: makeReplyMarkupForText(m.Text),
		})
	})

	return bot, nil
}

func makePoller(pollerType string) tb.Poller {
	if pollerType == "webhook" {
		port := os.Getenv("PORT")
		route := os.Getenv("ROUTE")
		// url that the web server will listen to
		listen := fmt.Sprintf(":%s/%s", port, route)
		// the webhook to be set to telegram API using setWebhook method
		webhook := os.Getenv("WEBHOOK_URL") + route

		return &tb.Webhook{Listen: listen, Endpoint: &tb.WebhookEndpoint{PublicURL: webhook}}
	}

	return &tb.LongPoller{Timeout: 10 * time.Second}
}

func cardSearch(query string) ([]scryfall.Card, error) {
	client, err := scryfall.NewClient()

	if err != nil {
		return nil, err
	}

	context := context.Background()

	options := scryfall.SearchCardsOptions{
		// Unique:        scryfall.UniqueModeCards,
		// Order:         scryfall.OrderName,
		// Dir:           scryfall.DirAuto,
	}

	result, err := client.SearchCards(
		context,
		query,
		options,
	)

	if err != nil {
		return nil, err
	}

	// show no more than 50 results, as per telegram's limitation of inline queries
	max := 50

	if length := len(result.Cards); length < 50 {
		max = length
	}

	cards := result.Cards[0:max]

	return cards, nil
}

func isDoubleFacedLayout(layout scryfall.Layout) bool {
	doubleFacedLayouts := []scryfall.Layout{"modal_dfc", "transform", "double_faced_token", "art_series"}

	for _, val := range doubleFacedLayouts {
		if layout == val {
			return true
		}
	}

	return false
}

// create a photo result for a Card
func newResultFromCard(card scryfall.Card) *tb.PhotoResult {
	result := &tb.PhotoResult{
		URL:      card.ImageURIs.Normal,
		ThumbURL: card.ImageURIs.Small,
		ResultBase: tb.ResultBase{
			ID:          card.ID,
			ReplyMarkup: makeReplyMarkupForResult(card),
		},
	}
	return result
}

// create a photo result for a face of a double-faced Card
func newResultFromFace(card scryfall.Card, faceIndex int) *tb.PhotoResult {
	face := card.CardFaces[faceIndex]
	faceID := fmt.Sprintf("%s-face-%d", card.ID, faceIndex)

	result := &tb.PhotoResult{
		URL:      face.ImageURIs.Normal,
		ThumbURL: face.ImageURIs.Small,
		ResultBase: tb.ResultBase{
			ID:          faceID,
			ReplyMarkup: makeReplyMarkupForResult(card),
		},
	}
	return result
}

// replay markup for inline query results
func makeReplyMarkupForResult(card scryfall.Card) *tb.InlineKeyboardMarkup {
	btn := tb.InlineButton{Text: "Details", URL: card.ScryfallURI}
	row := []tb.InlineButton{btn}
	grid := [][]tb.InlineButton{row}

	return &tb.InlineKeyboardMarkup{
		InlineKeyboard: grid,
	}
}

// reply markup for /start command
func makeReplyMarkupForStart() *tb.ReplyMarkup {
	btn := tb.InlineButton{Text: "Select a chat to see me in action", InlineQuery: "liliana"}
	row := []tb.InlineButton{btn}
	grid := [][]tb.InlineButton{row}

	return &tb.ReplyMarkup{
		InlineKeyboard: grid,
	}
}

// reply markup for /help command
func makeReplyMarkupForHelp() *tb.ReplyMarkup {
	btn1 := tb.InlineButton{Text: "Advanced Syntax Guide", URL: "https://scryfall.com/docs/syntax"}
	btn2 := tb.InlineButton{Text: "Learn about inline bots", URL: "https://telegram.org/blog/inline-bots"}
	row1 := []tb.InlineButton{btn1}
	row2 := []tb.InlineButton{btn2}
	grid := [][]tb.InlineButton{row1, row2}

	return &tb.ReplyMarkup{
		InlineKeyboard: grid,
	}
}

// reply markup for all text messages received
func makeReplyMarkupForText(text string) *tb.ReplyMarkup {
	btn := tb.InlineButton{Text: "Search this text", InlineQueryChat: text}
	row := []tb.InlineButton{btn}
	grid := [][]tb.InlineButton{row}

	return &tb.ReplyMarkup{
		InlineKeyboard: grid,
	}
}
