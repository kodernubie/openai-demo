package demo3

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func Init(app *fiber.App) {

	app.Get("/ws/demo3", websocket.New(func(c *websocket.Conn) {

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

	resp, err := getClient().CreateImage(
		context.Background(),
		openai.ImageRequest{
			Model:   openai.CreateImageModelDallE3,
			Prompt:  string(msg),
			N:       1,
			Size:    openai.CreateImageSize1024x1024,
			Quality: openai.CreateImageQualityHD,
			Style:   openai.CreateImageStyleNatural,
		},
	)

	if err != nil {
		fmt.Print("====>", err)
		c.WriteMessage(1, []byte("Error :"+err.Error()))
		return
	}

	fmt.Printf("==>>> %+v", resp)
	c.WriteMessage(1, []byte(resp.Data[0].URL))
}
