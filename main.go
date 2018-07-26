package main

import (
	"fmt"
	slack "github.com/monochromegane/slack-incoming-webhooks"
	// "github.com/iamucil/go-signal/signal"
)

func main() {
	fmt.Println("vim-go")
	// payload := {
	// 	Text: "User ",
	//	IconEmoji: ":calling:",
	//}
	//attachment := &{
	//	Text: "Twilio",
	//	Color: "#3AA3E3",
	//	Title: "Twilio",
	//}
	// slack.Payload.AddAttachment(attachment)

	slack.Client{
		WebhookURL: "https://hooks.slack.com/services/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",
	}.Post(&slack.Payload{
		Text:      "User +phone has been registered into radiumchat.",
		IconEmoji: ":calling:",
	})
}
