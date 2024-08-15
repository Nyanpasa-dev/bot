package commander

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		return
	}

	offset, err := strconv.Atoi(args)

	if err != nil {
		log.Panic("Error while converting offset string into int")
	}

	item, err := c.service.GetByOffset(offset)

	if err != nil {
		errmsg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(errmsg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, item.Title)

	c.bot.Send(msg)
}
