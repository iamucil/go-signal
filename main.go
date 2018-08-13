package main

import (
	"fmt"

	"encoding/json"
	"github.com/RadicalApp/libsignal-protocol-go/util/keyhelper"
	"github.com/graphql-go/graphql"
	"github.com/iamucil/go-signal/signal"
	"log"
	// slack "github.com/monochromegane/slack-incoming-webhooks"
)

func main() {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
	signal.Serializing()
	signal.Fingerprint()

	// schema
	fields := graphql.Fields{
		"Hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `{
        Hello
    }`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
	// payload := slack.Payload{
	// 	Text:      "User with phone number +phone-number has been registered into radiumchat",
	// 	IconEmoji: ":robot_face:",
	// 	Username:  "Radium-Bot",
	// }
	// attachment := &slack.Attachment{}
	// attachment.AuthorName = "Radium-messenger"
	// attachment.AuthorLink = "https://radium.id"
	// attachment.Title = "Twilio Payload"
	// attachment.Pretext = "Pretext Twilio Payload"
	// attachment.Text = "```Twilio Payload sender name ```"
	// attachment.Color = "#3AA3E3"
	// attachment.Fallback = "Twilio Payload."
	// attachment.FooterIcon = "https://platform.slack-edge.com/img/default_application_icon.png"
	// attachment.Footer = "Twilio-Messenger"

	// payload.AddAttachment(attachment)

	// slack.Client{
	// 	WebhookURL: "https://hooks.slack.com/services/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",
	// }.Post(&payload)
}
