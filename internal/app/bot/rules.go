package bot

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
	"strings"
)

func (bot *Bot) sendRules(obj events.MessageNewObject, user *User) {
	user.inRules = true

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	keyboard := object.NewMessagesKeyboardInline()

	row := keyboard.AddRow()
	row.AddTextButton("Крестики-нолики", "TTT_RULES", object.ButtonBlue)
	row.AddTextButton("Камень-ножницы-бумага", "RPS_RULES", object.ButtonBlue)
	row.AddTextButton("В меню", "MENU", object.ButtonBlue)

	messageToSend.Keyboard(keyboard)
	messageToSend.Message("Правила")

	bot.SendMessage(messageToSend.Params)
}

func (bot *Bot) receiveRulesCommands(obj events.MessageNewObject, user *User) {
	command := strings.Trim(obj.Message.Payload, "\"")

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	switch command {
	case "TTT_RULES":

		messageToSend.Message("Правила максимально просты: " +
			"\n1. кто будет первый ходить выбирается рандомно, " +
			"\n2. ты играешь за Х, " +
			"\n3. выбор клетки происходит строго с кнопок клавиатуры, " +
			"\n4. чтобы завершить игру нажми завершить.")

		bot.SendMessage(messageToSend.Params)

		bot.sendRules(obj, user)

		return
	case "RPS_RULES":
		messageToSend.Message("Правила: " +
			"\n1. камень бьет ножницы," +
			"\n2. ножницы бьют бумагу, " +
			"\n3. бумага бьет камень," +
			"\n4. чтобы завершить игру нажми завершить.")

		bot.SendMessage(messageToSend.Params)

		bot.sendRules(obj, user)

		return
	case "MENU":
		user.inRules = false
		bot.sendMenu(obj, user)

		return
	}
}
