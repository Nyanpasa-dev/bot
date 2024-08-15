package commander

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nyanpasa-dev/bot/internal/service/product"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *product.Service
}

func NewCommandRouter(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot:     bot,
		service: service,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Printf("Recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Commander) handleMessage(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}

func (c *Commander) handleCallback(callback *tgbotapi.CallbackQuery) {
	parsedData := CommandData{}

	err := json.Unmarshal([]byte(callback.Data), &parsedData)

	if err != nil {
		log.Panic("Cannot parse data")
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v", parsedData))
	c.bot.Send(msg)

}
