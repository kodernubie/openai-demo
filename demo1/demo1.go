package demo1

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func Init(app *fiber.App) {

	app.Get("/ws/demo1", websocket.New(func(c *websocket.Conn) {

		var msg []byte
		var err error

		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("error read", err)
				break
			}

			log.Printf("recv: %s", msg)

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
					Role:    openai.ChatMessageRoleUser,
					Content: string(msg),
				},
			},
		},
	)

	if err != nil {
		c.WriteMessage(1, []byte("Error :"+err.Error()))
		return
	}

	c.WriteMessage(1, []byte(resp.Choices[0].Message.Content))
}
