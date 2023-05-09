package bot

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/blrrubik/vk_bot/internal/app/service"
	"strconv"
	"strings"
)

func (bot *Bot) receiveRPSCommands(obj events.MessageNewObject, user *User) {
	rps := user.GameService.Services["RPS"].(service.RockPaperScissorsInterface)

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	command := strings.Trim(obj.Message.Payload, "\"")

	if command == "stop" {
		rps.EndGame()
		messageToSend.Message("Игра закончена")

		bot.SendMessage(messageToSend.Params)
		bot.sendMenu(obj, user)

		return
	}

	playerChoice, err := strconv.Atoi(command)
	if err != nil {
		messageToSend.Message("Invalid input")
		bot.SendMessage(messageToSend.Params)

		return
	}

	messageToSend.Message(rps.IsWin(playerChoice))

	bot.SendMessage(messageToSend.Params)

	rps.EndGame()

	bot.startRPS(obj, user)
}
