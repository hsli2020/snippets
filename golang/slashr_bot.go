// https://github.com/tudurom/slashr_bot
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"
)

type Config struct {
	Token string      `json:"token"`
	Env   Environment `json:"env"`
}

type Environment int

const (
	Debug      Environment = iota
	Production Environment = iota
)

var configPath = "config.json"
var env = Debug
var bot *tgbotapi.BotAPI

const subredditRegex = `(?:\A|\s)(?:\/)?(?P<sublink>[ru])\/(?P<subspec>[a-zA-Z0-9_\-]*)(?:\s|\z)`

func (e *Environment) UnmarshalJSON(data []byte) error {
	var str string

	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	if strings.EqualFold(str, "debug") {
		*e = Debug
	} else if strings.EqualFold(str, "production") {
		*e = Production
	} else {
		return errors.New("Invalid value for Environment type")
	}
	return nil
}

func (c *Config) readConfig(fp string) error {
	buf, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, c)
	if err != nil {
		return err
	}

	return nil
}

func initLogger(env Environment) {
	if env == Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func initBotAPI(token string, env Environment) error {
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	logrus.WithField("username", bot.Self.UserName).Debug("Successfully connected to Telegram API")
	return nil
}

func main() {
	var conf Config
	var rx = regexp.MustCompile(subredditRegex)

	err := conf.readConfig(configPath)
	if err != nil {
		logrus.WithField(
			"configPath", configPath,
		).Fatalf("Failed to read config file: %v", err)
	}
	initLogger(conf.Env)
	err = initBotAPI(conf.Token, conf.Env)
	if err != nil {
		logrus.WithField("token", conf.Token).Fatalf("Couldn't connect to Telegram API: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logrus.WithField("update", u).Fatalf("Couldn't get updates channel: %v", err)
	}

	for update := range updates {
		var text string
		var msg *tgbotapi.Message
		var q *tgbotapi.InlineQuery
		if update.Message != nil {
			msg = update.Message

			text = msg.Text
		} else if update.InlineQuery != nil {
			q = update.InlineQuery
			text = q.Query
		}
		matches := rx.FindAllStringSubmatch(text, -1)

		// Found subreddits
		if len(matches) > 0 {
			s := ""
			for _, m := range matches {
				// get sublink and sub/user from regex match
				logrus.WithField("match", m[0]).Debug("Got match")
				result := make(map[string]string)
				for i, name := range rx.SubexpNames() {
					if i != 0 {
						result[name] = m[i]
					}
				}
				link := result["sublink"]
				sub := result["subspec"]
				newlink := fmt.Sprintf("https://reddit.com/%s/%s", link, sub)
				md := fmt.Sprintf("[/%s/%s](%s)", link, sub, newlink)

				if update.Message != nil {
					foundEntity := false
					if msg.Entities != nil {
						for _, ent := range *msg.Entities {
							if ent.URL == newlink {
								foundEntity = true
								break
							}
						}
					}
					if !foundEntity {
						s += md + "\n"
					}
				}

				text = string(rx.ReplaceAll([]byte(text), []byte(" "+md+" ")))
			}

			if update.Message != nil && s != "" {
				reply := tgbotapi.NewMessage(msg.Chat.ID, s)
				reply.ReplyToMessageID = msg.MessageID
				reply.ParseMode = "markdown"

				_, err = bot.Send(reply)
			} else if update.InlineQuery != nil {
				doc := tgbotapi.NewInlineQueryResultArticleMarkdown(time.Now().Format("slashr_%s"), "Result", text)
				doc.Description = text
				results := make([]interface{}, 1)
				results[0] = doc
				ic := tgbotapi.InlineConfig{}
				ic.InlineQueryID = q.ID
				ic.Results = results
				ic.IsPersonal = true
				ic.CacheTime = 86400
				_, err = bot.AnswerInlineQuery(ic)
			}
			if err != nil {
				logrus.Infof("Couldn't send message: %v", err)
			}

		}
	}
}
