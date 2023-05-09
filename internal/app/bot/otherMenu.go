package bot

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
	"strings"
	"time"
)

func (bot *Bot) sendOtherMenu(obj events.MessageNewObject, user *User) {
	user.inOtherMenu = true

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	keyboard := object.NewMessagesKeyboardInline()

	row := keyboard.AddRow()
	row.AddTextButton("ping-pong", "PP", object.ButtonBlue)
	row.AddTextButton("Время", "TIME", object.ButtonBlue)
	row.AddTextButton("В меню", "MENU", object.ButtonBlue)

	messageToSend.Keyboard(keyboard)
	messageToSend.Message("Прочее")

	bot.SendMessage(messageToSend.Params)
}

func (bot *Bot) receiveOtherMenuCommands(obj events.MessageNewObject, user *User) {
	command := strings.Trim(obj.Message.Payload, "\"")

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	switch command {
	case "PP":
		keyboard := object.NewMessagesKeyboardInline()
		row := keyboard.AddRow()
		row.AddTextButton("ping", "PING", object.ButtonBlue)

		messageToSend.Keyboard(keyboard)
		messageToSend.Message("Жми")

		bot.SendMessage(messageToSend.Params)

		return
	case "PING":
		messageToSend.Message("pong")
		bot.SendMessage(messageToSend.Params)

		bot.sendOtherMenu(obj, user)

		return
	case "TIME":
		loc, _ := time.LoadLocation("Europe/Moscow")
		messageToSend.Message(time.Now().In(loc).Format("02-01-2006 15:04:05"))
		bot.SendMessage(messageToSend.Params)

		bot.sendOtherMenu(obj, user)

		return
	case "MENU":
		user.inOtherMenu = false
		bot.sendMenu(obj, user)

		return
	}
}
