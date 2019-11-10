package router

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func validTimeCheck(t string) (string, bool) {
	if len(t) == 1 || len(t) == 2 {
		t = t + ":00pm"
	} else {
		t = t + "pm"
	}
	r := regexp.MustCompile("[^:pm0-9]")
	correctCheck := r.ReplaceAllString(t, "")
	if t != correctCheck {
		return "", false
	}
	correctTime, err := time.Parse("3:04pm", correctCheck)
	if err != nil {
		return "", false
	}
	return correctTime.Format("15:04"), true
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
	if len(players) == 0 {
		return "нету голосов"
	}
	squad := fmt.Sprintf("всего %s игроков \n", strconv.Itoa(len(players)))
	for i, player := range players {
		squad = squad + strconv.Itoa(i+1) + ") @" + player + "\n"
	}
	return squad
}

func removePlayerFromList(player string, players []string) []string {
	for i, nick := range players {
		if player == nick {
			players = append(players[:i], players[i+1:]...)
		}
	}
	return players
}
