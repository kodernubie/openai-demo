package demo1

import (
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

			handler(msg)
		}
	}))
}

func getClient() *openai.Client {

	if client == nil {
		client = openai.NewClient(os.Getenv("API_KEY"))
	}

	return client
}

func handler(msg []byte) {

	getClient()
}
