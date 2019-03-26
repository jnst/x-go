package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	url, ok := os.LookupEnv("WEBHOOK_URL")
	if !ok {
		log.Fatal("missing WEBHOOK_URL")
	}

	attachment := slack.Attachment{
		Color: "good",
		Text:  "<!channel> <@jnst> test text",
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(url, &msg)
	if err != nil {
		log.Fatal("failed to post slack api")
	}
}
