package demo2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	openai "github.com/sashabaranov/go-openai"
)

type UserEvent struct {
	Text     string `json:"text"`
	ImageURL string `json:"imageURL"`
}

var client *openai.Client

func Init(app *fiber.App) {

	app.Get("/ws/demo2", websocket.New(func(c *websocket.Conn) {

		var msg []byte
		var err error

		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("error read", err)
				break
			}

			log.Printf("recv: %s", msg)

			fmt.Println("masuk... 111111")
			handler(c, msg)
		}
	}))
}

func getClient() *openai.Client {

	if client == nil {
		client = openai.NewClient(os.Getenv("API_KEY"))
	}

	return client
}

func handler(c *websocket.Conn, msg []byte) {

	event := UserEvent{}

	err := json.Unmarshal(msg, &event)

	fmt.Println("masuk... 2222222")

	if err != nil {
		fmt.Println("masuk... error", err)
		c.WriteMessage(1, []byte("Error :"+err.Error()))
		return
	}

	userContent := []openai.ChatMessagePart{}

	if event.ImageURL != "" {
		userContent = append(userContent, openai.ChatMessagePart{
			Type: openai.ChatMessagePartTypeImageURL,
			ImageURL: &openai.ChatMessageImageURL{
				URL: event.ImageURL,
			},
		})
	}

	if event.Text != "" {
		userContent = append(userContent, openai.ChatMessagePart{
			Type: openai.ChatMessagePartTypeText,
			Text: event.Text,
		})
	}

	fmt.Println("masuk... 33333")

	resp, err := getClient().CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a polite chat bot. Response question with sydney sheldon style",
				},
				{
					Role:         openai.ChatMessageRoleUser,
					MultiContent: userContent,
				},
			},
		},
	)

	if err != nil {
		c.WriteMessage(1, []byte("Error :"+err.Error()))
		return
	}

	fmt.Printf("==>>> %+v", resp)
	c.WriteMessage(1, []byte(resp.Choices[0].Message.Content))
}
