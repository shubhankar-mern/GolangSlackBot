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

	//os.Setenv("SLACK_BOT_TOKEN", "xoxb-6001447099542-6005180409685-SLjjEr1uRm2IvmqVMSFrGIrN")
	//os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0607V7HUJX-6031085918016-5f135689cd422193029bc08c09b989c0ce494035d94a3d924c4d0ad47d440ca5")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("hii", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello man, I am a Bot")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
