package bot

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
	"strings"
)

func (bot *Bot) sendMenu(obj events.MessageNewObject, user *User) {
	user.inMenu = true

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	keyboard := object.NewMessagesKeyboardInline()

	row1 := keyboard.AddRow()
	row1.AddTextButton("Крестики-нолики", "TTT", object.ButtonBlue)
	row1.AddTextButton("Камень-ножницы-бумага", "RPS", object.ButtonBlue)
	row2 := keyboard.AddRow()
	row2.AddTextButton("Правила", "RULES", object.ButtonBlue)
	row2.AddTextButton("Другое", "OTHER", object.ButtonBlue)

	messageToSend.Keyboard(keyboard)
	messageToSend.Message("Меню")

	bot.SendMessage(messageToSend.Params)
}

func (bot *Bot) receiveMenuCommands(obj events.MessageNewObject, user *User) {
	command := strings.Trim(obj.Message.Payload, "\"")

	switch command {
	case "TTT":
		user.inMenu = false
		bot.startTTT(obj, user)

		return
	case "RPS":
		user.inMenu = false
		bot.startRPS(obj, user)

		return
	case "RULES":
		user.inMenu = false
		bot.sendRules(obj, user)

		return
	case "OTHER":
		user.inMenu = false
		bot.sendOtherMenu(obj, user)

		return
	}
}
