package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/nyanpasa-dev/bot/internal/service/product"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message)
		case "products":
			productsCommand(bot, update.Message, productService)
		default:
			defaultBehaviour(bot, update.Message)
		}
	}
}

func productsCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, service *product.Service) {
	products, err := service.List()
	var productsString string

	if err != nil {
		log.Panic(err)
	}

	for _, v := range products {
		productsString += "- " + string(v.Title) + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, productsString)

	bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Available commands:\n/help - show this message\n")

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "help - show this message\nlist - show list of commands\n")

	bot.Send(msg)
}

func defaultBehaviour(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	bot.Send(msg)
}
