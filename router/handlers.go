package router

import (
	"fmt"
	"log"

	config "github.com/AkezhanOb1/dotaTelegramBot/config/bot"
	"github.com/AkezhanOb1/dotaTelegramBot/style/buttons"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot        = config.Bot
	inPlayers  = []string{}
	outPlayers = []string{}
	gameTime   = ""
)

//Start - responsible for handling /start comand
func Start(m *tb.Message) {
	message := fmt.Sprintf(`/game - начать голосование + обнулить результаты предыдущего голосования, надо указать время игры (/vote 5)
	/in - дать согласие на игру
	/out - отказаться от игры
	/inlist - список игроков на старте
	/outlist - список пиздюков
	/time - узнать время игры
	/changetime - изменить время игры время указывается через пробел (/changetime 5)`)
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
	gameTime = dotaTime

	message, _ := bot.Send(
		m.Chat,
		fmt.Sprintf("э супергеройлар, сегодня в %s дота 5x5", dotaTime),
		&tb.ReplyMarkup{InlineKeyboard: buttons.VoteKeys},
	)
	err := bot.Pin(message, tb.Silent)
	if err != nil {
		log.Println(err)
	}
}

//GameTime - list of all players not playing today
func GameTime(m *tb.Message) {
	if gameTime == "" {
		bot.Send(m.Chat, "время игры еще не выбрана")
		return
	}
	message := "Игра в " + gameTime
	bot.Send(m.Chat, message)
}

//GameTimeChange - changes the time of the game
func GameTimeChange(m *tb.Message) {
	if gameTime == "" {
		bot.Send(m.Chat, "сначала начните опрос командой /vote")
		return
	}
	newTime, ok := validTimeCheck(m.Payload)
	if !ok {
		bot.Send(m.Chat, "Петух введи в коректное время, пример для особо одаренных /changetime 5")
		return
	}
	gameTime = newTime
	bot.Send(m.Chat, "Время игры было изменено на "+gameTime)

}

//InPlayer - adds player to the list of players ready to play
func InPlayer(m *tb.Message) {
	nickname := m.Sender.Username
	outPlayers = removePlayerFromList(nickname, outPlayers)
	inPlayers = addPlayerToList(nickname, inPlayers)
	bot.Send(m.Chat, fmt.Sprintf("%s в игре, список работяг /inlist", nickname))
}

//OutPlayer - adds player to the list of players ready who is not playing today
func OutPlayer(m *tb.Message) {
	nickname := m.Sender.Username
	inPlayers = removePlayerFromList(nickname, inPlayers)
	outPlayers = addPlayerToList(nickname, outPlayers)
	bot.Send(m.Chat, fmt.Sprintf("%s минус, список петухов /outlist", nickname))
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
	c.Message.Sender = c.Sender
	InPlayer(c.Message)
}

//PollDisAgreement - handler for players who disagreed with the poll
func PollDisAgreement(c *tb.Callback) {
	bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	c.Message.Sender = c.Sender
	OutPlayer(c.Message)
}
