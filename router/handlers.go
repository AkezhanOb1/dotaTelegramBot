package router

import (
	"fmt"

	"github.com/AkezhanOb1/dotaTelegramBot/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot        = config.Bot
	inPlayers  = []string{}
	outPlayers = []string{}
)

//Start - responsible for handling /start comand
func Start(m *tb.Message) {
	message := fmt.Sprintf("/vote - начать голосование + обнулить результаты предыдущего голосования надо указать время игры (/vote 5)\n /in - список игроков на старте \n /out - список пиздюков \n /restart - обнулить результаты предыдущего голосования")
	bot.Send(m.Chat, message)
}

// Vote - responsible for creating a poll
func Vote(m *tb.Message) {
	inPlayers = []string{}
	outPlayers = []string{}
	dotaTime, ok := validTimeCheck(m.Payload)
	if !ok {
		bot.Send(m.Chat, "Петух введи в какое время хочешь катать доту, пример для особо одаренных /vote 5")
		return
	}
	bot.Send(
		m.Chat,
		fmt.Sprintf("Паца хотят в дотку 5х5, сегодня в %s", dotaTime),
		&tb.ReplyMarkup{InlineKeyboard: config.VoteKeys},
	)
}

//InPlayers - list of all players ready to play today
func InPlayers(m *tb.Message) {
	squad := listOfPlayers(inPlayers)
	bot.Send(m.Chat, squad)
}

//OutPlayers - list of all players not playing today
func OutPlayers(m *tb.Message) {
	squad := listOfPlayers(outPlayers)
	bot.Send(m.Chat, squad)
}

//PollAgreement - handler for players who agreed with the poll
func PollAgreement(c *tb.Callback) {
	bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	nickname := c.Sender.Username
	outPlayers = removePlayerFromList(nickname, outPlayers)
	if ok := checkPlayer(nickname, inPlayers); !ok {
		inPlayers = append(inPlayers, nickname)
	}
}

//PollDisAgreement - handler for players who disagreed with the poll
func PollDisAgreement(c *tb.Callback) {
	bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	nickname := c.Sender.Username
	inPlayers = removePlayerFromList(nickname, inPlayers)
	if ok := checkPlayer(nickname, outPlayers); !ok {
		outPlayers = append(outPlayers, nickname)
	}
}