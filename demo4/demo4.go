package demo4

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func Init(app *fiber.App) {

	app.Get("/ws/demo4", websocket.New(func(c *websocket.Conn) {

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

	resp, err := getClient().CreateSpeech(
		context.Background(),
		openai.CreateSpeechRequest{
			Model:          openai.TTSModel1HD,
			Voice:          openai.VoiceAlloy,
			Input:          string(msg),
			ResponseFormat: openai.SpeechResponseFormatMp3,
		},
	)

	if err != nil {
		fmt.Print("====>", err)
		c.WriteMessage(1, []byte("Error :"+err.Error()))
		return
	}

	fmt.Printf("==>>> %+v", resp)

	fileName := ulid.Make().String()
	filePath := "./web/" + fileName + ".mp3"

	fl, err := os.Create(filePath)

	byt := []byte{}
	no, err := resp.Read(byt)

	fmt.Println("Read", no)
	for err == nil && no > 0 {
		no, err = fl.Write(byt)
		fmt.Println("Read", no)
	}

	fl.Close()

	c.WriteMessage(1, []byte("./"+fileName+".mp3"))
}
