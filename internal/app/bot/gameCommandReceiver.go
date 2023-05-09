package bot

import (
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/blrrubik/vk_bot/internal/app/service"
)

func (bot *Bot) receiveGameCommands(obj events.MessageNewObject, game service.Game, user *User) {

	switch game.(type) {
	case service.TicTacToeInterface:
		bot.receiveTicTacToeCommands(obj, user)
	case service.RockPaperScissorsInterface:
		bot.receiveRPSCommands(obj, user)
	}

}
