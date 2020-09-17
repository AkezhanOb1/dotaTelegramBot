package buttons

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//GoBtn - represents agreement with the poll
var GoBtn = tb.InlineButton{
	Unique: "Go",
	Text:   "Go",
}

//NopeBtn - represents disagreement with the poll
var NopeBtn = tb.InlineButton{
	Unique: "Nope",
	Text:   "Nope",
}

//VoteKeys - represents keys for the poll
var VoteKeys = [][]tb.InlineButton{
	[]tb.InlineButton{GoBtn, NopeBtn},
}
