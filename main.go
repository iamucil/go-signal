package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RadicalApp/libsignal-protocol-go/util/keyhelper"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/iamucil/go-signal/signal"
	"io/ioutil"
	"log"
	"net/http"
	// slack "github.com/monochromegane/slack-incoming-webhooks"
)

// Post is post
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Comment is comment
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Account
type Account struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Type      string `json:"type"`
}

func main() {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
	signal.Serializing()
	signal.Fingerprint()

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			createQueryType(
				createPostType(
					createCommentType(),
				),
			),
		),
	})

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})
	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
	// schema
	/*
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

			Query
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
	*/
	/*
	   Implementing slack webhook from go-lang
	   using package from github.com/monochromegane/slack-incoming-webhooks
	   go get github.com/monochromegane/slack-incoming-webhooks
	   import (
	       slack github.com/monochromegane/slack-incoming-webhooks
	   )
	*/
	/*
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
			}.Post(&payload) */
}

func createQueryType(postType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "QueryType", Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: postType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"]
				v, _ := id.(int)
				log.Printf("fetching post with id: %d", v)

				return fetchPostByiD(v)
			},
		},
		"Hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"iam": &graphql.Field{
			Type: postType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"]
				v, _ := id.(int)
				log.Printf("Fetching account id: %v", v)

				// return "world", nil
				return Account{}, nil
			},
		},
	}}
}

func createPostType(commentType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"userId": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					post, _ := p.Source.(*Post)
					log.Printf("Fetching comments of post with id: %d", post.ID)
					return fetchCommentsByPostID(post.ID)
				},
			},
		},
	})
}

func createCommentType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"postid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func fetchPostByiD(id int) (*Post, error) {
	resp, err := http.Get(fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%d", id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "Could not fetch data", resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read data")
	}

	result := Post{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("Could not unmarshal data")
	}
	return &result, nil
}

func fetchCommentsByPostID(id int) ([]Comment, error) {
	resp, err := http.Get(fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%d/comments", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "Could not fetch data", resp.Status)

	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read data")
	}
	result := []Comment{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("Could not unmarshal data")
	}

	return result, nil
}
