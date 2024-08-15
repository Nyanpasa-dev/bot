package commander

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Callback(callback *tgbotapi.CallbackQuery) {
	parsedData := CommandData{}

	err := json.Unmarshal([]byte(callback.Data), &parsedData)

	if err != nil {
		log.Panic("Cannot parse data")
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v", parsedData))
	c.bot.Send(msg)

}
