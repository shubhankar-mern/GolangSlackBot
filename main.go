package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events", event.Timestamp, event.Command, event.Parameters, event.Event)
	}
}

func main() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	//os.Setenv("SLACK_BOT_TOKEN", "")
	//os.Setenv("SLACK_APP_TOKEN", "")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("hii", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello man, I am a Bot")
		},
	})

	bot.Command("bye", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Bye man, have a nice day")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
