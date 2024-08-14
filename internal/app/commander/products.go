package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
