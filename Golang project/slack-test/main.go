package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4416536459077-4613819663154-k4MbsOYK3D9rxpIUZdSWbzhn")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04JECAM5RP-4616571973012-040bfc3f441ca44683b83363d9e8647b416ca4721c7dc8f1e3256777d88f1128")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
