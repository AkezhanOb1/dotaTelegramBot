package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "923261617:AAGIM0wr4rrUZxFL7xCAn1i62eQK5h7pCUE",
		URL:    "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	inPlayers := []string{}
	outPlayers := []string{}

	b.Handle("/start", func(m *tb.Message) {
		message := "/vote - starts poll, /in - list of players ready to play, /out - list of players who will not play today, /restart - nullifing voices"
		b.Send(m.Chat, message)
	})

	goBtn := tb.InlineButton{
		Unique: "Go",
		Text:   "Go",
	}

	nopeBtn := tb.InlineButton{
		Unique: "Nope",
		Text:   "Nope ",
	}

	b.Handle(&goBtn, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		nickname := c.Sender.Username
		outPlayers = removePlayerFromList(nickname, outPlayers)
		if ok := checkPlayer(nickname, inPlayers); !ok {
			inPlayers = append(inPlayers, nickname)
		}
	})

	b.Handle(&nopeBtn, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		nickname := c.Sender.Username
		inPlayers = removePlayerFromList(nickname, inPlayers)
		if ok := checkPlayer(nickname, outPlayers); !ok {
			outPlayers = append(outPlayers, nickname)
		}
	})

	voteKeys := [][]tb.InlineButton{
		[]tb.InlineButton{goBtn, nopeBtn},
	}
	b.Handle("/vote", func(m *tb.Message) {
		b.Send(
			m.Chat,
			"We are going to play dota today, are you with us?",
			&tb.ReplyMarkup{InlineKeyboard: voteKeys})
	})

	b.Handle("/restart", func(m *tb.Message) {
		inPlayers = []string{}
		outPlayers = []string{}

		b.Send(m.Chat, "nullifing voices")
	})

	b.Handle("/in", func(m *tb.Message) {
		squad := listOfPlayers(inPlayers)
		b.Send(m.Chat, squad)
	})
	b.Handle("/out", func(m *tb.Message) {
		squad := listOfPlayers(outPlayers)
		b.Send(m.Chat, squad)
	})

	b.Start()
}

func checkPlayer(player string, players []string) bool {
	for _, nick := range players {
		if nick == player {
			return true
		}
	}
	return false
}

func listOfPlayers(players []string) string {
	log.Println("LIST", players)
	if len(players) == 0 {
		return "no players"
	}
	squad := ""
	for _, player := range players {
		squad = squad + player + " "
	}
	return squad
}

func removePlayerFromList(player string, players []string) []string {
	log.Println(players, "before")
	for i, nick := range players {
		if player == nick {
			players = append(players[:i], players[i+1:]...)
		}
	}
	return players
}
