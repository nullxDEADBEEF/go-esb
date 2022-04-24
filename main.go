package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Steps:
// - Pass message to the ESB, (XML, JSON, TSV)
// - Register users and provide them with a token,
//   ESB will transform the message depending on who is consuming it
// - Consumer needs to identify itself, this is done via token
// - Consumer informs ESB to get messages from a specific provider

type User struct {
	Id    string `json:"id" xml:"Id"`
	Token string `json:"token" xml:"Token"`
}

type ProviderMessage struct {
	Id        string `json:"id" xml:"Id"`
	Message   string `json:"message" xml:"Message"`
	CreatedAt string `json:"created_at" xml:"CreatedAt"`
}

var users = []User{
	{
		Id:    "1",
		Token: "1212",
	},
	{
		Id:    "2",
		Token: "3333",
	},
}

var messages = map[string][]ProviderMessage{
	"1": {
		{
			Id: "1a76658a-7e4c-4f24-9a96-f68ef3526008", Message: "I am message 1", CreatedAt: "1650807757",
		},
		{

			Id: "87a17661-13ae-45c3-bfb2-041afa15234a", Message: "I am message 2", CreatedAt: "1650807758",
		},
		{
			Id: "aa229257-48f3-46e2-88dc-beb210e7f9e4", Message: "I am message 3", CreatedAt: "1650807759",
		},
		{
			Id: "410cc421-71c0-414a-99d6-6e603e741692", Message: "I am message 4", CreatedAt: "1650807760",
		},
	},
}

func main() {
	router := gin.Default()
	router.GET("/provider/:id/token/:token", esb)
	router.Run(":9999")
}

func esb(c *gin.Context) {
	token := c.Param("token")
	msg_id := c.Param("id")
	for _, v := range users {
		if v.Token == token {
			// transform message based on consumer
			transformMessage(c, token, messages[msg_id])
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"info": "Invalid token"})
}

func transformMessage(c *gin.Context, consumerToken string, message []ProviderMessage) {
	if consumerToken == "1212" {
		JSONTransformer(c, message)
		return
	} else {
		XMLTransformer(c, message)
	}
}
