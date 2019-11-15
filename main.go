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
	b.Handle("/inlist", router.InPlayers)
	b.Handle("/outlist", router.OutPlayers)
	b.Handle("/in", router.InPlayer)
	b.Handle("/out", router.OutPlayer)
	b.Handle(&config.GoBtn, router.PollAgreement)
	b.Handle(&config.NopeBtn, router.PollDisAgreement)
	b.Start()
}
