package main

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = "6677"
		publicURL = "https://dotapollbot.herokuapp.com"
		token     = "923261617:AAGIM0wr4rrUZxFL7xCAn1i62eQK5h7pCUE"
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hi!")
	})
}
