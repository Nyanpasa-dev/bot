package commander

import (
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

func (c *Commander) Product(inputMessage *tgbotapi.Message) {
	products, err := c.service.List()
	var productsString string

	if err != nil {
		log.Panic(err)
	}

	for _, v := range products {
		productsString += "- " + string(v.Title) + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, productsString)

	c.bot.Send(msg)
}

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Available commands:\n/help - show this message\n")

	c.bot.Send(msg)
}

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "help - show this message\nlist - show list of commands\n")

	c.bot.Send(msg)
}

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	c.bot.Send(msg)
}
