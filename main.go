package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env variables")
	}
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("YOB <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Examples:    []string{"YOB 2000", "YOB 1995"},
		Handler: func(botctx slacker.BotContext, req slacker.Request, resp slacker.ResponseWriter) {
			y, err := strconv.Atoi(req.Param("year"))
			if err != nil {
				fmt.Printf("Error while parsing year: %v", err)
			}
			age := 2024 - y
			r := fmt.Sprintf("age is %v", age)
			resp.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
