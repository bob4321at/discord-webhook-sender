package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gtuk/discordwebhook"
)

var url = ""

type ReqMessege struct {
	Url      string
	Name     string
	Img_Url  string
	Messegee string
	Fields   [][]string
	Footer   string
}

func sendEmbededMessege(name string, img_url string, title string, content discordwebhook.Footer, fields []discordwebhook.Field) {
	username := name

	red := "16774912"

	image := discordwebhook.Image{
		Url: &img_url,
	}

	message := discordwebhook.Message{
		Username: &username,
		Embeds: &[]discordwebhook.Embed{
			{Title: &title, Color: &red, Footer: &content, Fields: &fields, Image: &image},
		},
	}
	if err := discordwebhook.SendMessage(url, message); err != nil {
		fmt.Println(err)
	}
}

func newFooter(text string) discordwebhook.Footer {
	return discordwebhook.Footer{Text: &text}
}

func newFields(name, value []string) []discordwebhook.Field {
	fields := []discordwebhook.Field{}
	for i := 0; i < len(name); i++ {
		Name := name[i]
		Value := value[i]
		Inline := false
		fields = append(fields, discordwebhook.Field{Name: &Name, Value: &Value, Inline: &Inline})
	}
	return fields
}

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("./templates/index.tmpl")
	router.StaticFile("/js", "./templates/index.js")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	router.POST("/sendMessege", func(c *gin.Context) {
		req := c.Request.Body
		temp_info, err := io.ReadAll(req)
		if err != nil {
			panic(err)
		}

		info := ReqMessege{}

		json.Unmarshal(temp_info, &info)

		fmt.Println(info)

		fields := []discordwebhook.Field{}

		for i := 0; i < len(info.Fields); i++ {
			inline := false
			fields = append(fields, discordwebhook.Field{Name: &info.Fields[i][0], Value: &info.Fields[i][1], Inline: &inline})
		}

		url = info.Url

		sendEmbededMessege(info.Name, info.Img_Url, info.Messegee, newFooter(info.Footer), fields)
	})

	router.Run()
}
