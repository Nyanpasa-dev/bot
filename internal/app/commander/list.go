package commander

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	products, err := c.service.List()
	var productsString string

	if err != nil {
		log.Panic(err)
	}

	for _, v := range products {
		productsString += "- " + string(v.Title) + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, productsString)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next", string(serializedData)),
		),
	)

	msg.ReplyMarkup = numericKeyboard

	c.bot.Send(msg)
}
