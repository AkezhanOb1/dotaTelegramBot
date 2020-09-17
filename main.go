package main

import (
	config "github.com/AkezhanOb1/dotaTelegramBot/config/bot"
	"github.com/AkezhanOb1/dotaTelegramBot/router"
	"github.com/AkezhanOb1/dotaTelegramBot/style/buttons"
)

func main() {
	bot := config.Bot
	bot.Handle("/start", router.Start)
	bot.Handle("/game", router.Vote)
	bot.Handle("/time", router.GameTime)
	bot.Handle("/changetime", router.GameTimeChange)
	bot.Handle("/inlist", router.InPlayers)
	bot.Handle("/outlist", router.OutPlayers)
	bot.Handle("/in", router.InPlayer)
	bot.Handle("/out", router.OutPlayer)
	bot.Handle(&buttons.GoBtn, router.PollAgreement)
	bot.Handle(&buttons.NopeBtn, router.PollDisAgreement)
	bot.Start()
}
