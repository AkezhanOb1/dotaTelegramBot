package main

import (
	"github.com/AkezhanOb1/dotaTelegramBot/config"
	"github.com/AkezhanOb1/dotaTelegramBot/router"
)

func main() {
	b := config.Bot
	b.Handle("/start", router.Start)
	b.Handle("/vote", router.Vote)
	b.Handle("/time", router.GameTime)
	b.Handle("/changetime", router.GameTimeChange)
	b.Handle("/in", router.InPlayers)
	b.Handle("/out", router.OutPlayers)
	b.Handle(&config.GoBtn, router.PollAgreement)
	b.Handle(&config.NopeBtn, router.PollDisAgreement)
	b.Start()
}
