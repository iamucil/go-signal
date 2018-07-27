package main

import (
	// "fmt"
	slack "github.com/monochromegane/slack-incoming-webhooks"
	// "github.com/iamucil/go-signal/signal"
)

func main() {
	payload := slack.Payload{
		Text:      "User with phone number +phone-number has been registered into radiumchat",
		IconEmoji: ":robot_face:",
		Username:  "Radium-Bot",
	}
	attachment := &slack.Attachment{}
	attachment.AuthorName = "Radium-messenger"
	attachment.AuthorLink = "https://radium.id"
	attachment.Title = "Twilio Payload"
	attachment.Pretext = "Pretext Twilio Payload"
	attachment.Text = "```Twilio Payload sender name ```"
	attachment.Color = "#3AA3E3"
	attachment.Fallback = "Twilio Payload."
	attachment.FooterIcon = "https://platform.slack-edge.com/img/default_application_icon.png"
	attachment.Footer = "Twilio-Messenger"

	payload.AddAttachment(attachment)

	slack.Client{
		WebhookURL: "https://hooks.slack.com/services/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",
	}.Post(&payload)
}
