package bot

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/blrrubik/vk_bot/internal/app/service"
	"math/rand"
)

func (bot *Bot) startRPS(obj events.MessageNewObject, user *User) {
	rps := user.GameService.Services["RPS"]

	rps.(service.RockPaperScissorsInterface).StartGame()

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	keyboard := object.NewMessagesKeyboardInline()

	row1 := keyboard.AddRow()
	row1.AddTextButton("\U0001FAA8", 1, object.ButtonBlue)
	row1.AddTextButton("\u2702\ufe0f", 2, object.ButtonBlue)
	row1.AddTextButton("\U0001F9FB", 3, object.ButtonBlue)

	row2 := keyboard.AddRow()
	row2.AddTextButton("Закончить", "stop", object.ButtonRed)

	messageToSend.Keyboard(keyboard)
	messageToSend.Message("Выбирай")

	bot.SendMessage(messageToSend.Params)
}

func (bot *Bot) startTTT(obj events.MessageNewObject, user *User) {

	ticTacToe := user.GameService.Services["TTT"].(service.TicTacToeInterface)

	board := ticTacToe.StartGame()

	if rand.Intn(2)+1 == 2 {
		ticTacToe.ComputerTurn()
	}

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	keyboard := object.NewMessagesKeyboardInline()

	for i := 0; i < 3; i++ {
		row := keyboard.AddRow()
		row.AddTextButton(board[0][i], fmt.Sprintf("%d %d", 0, i), object.ButtonBlue)
		row.AddTextButton(board[1][i], fmt.Sprintf("%d %d", 1, i), object.ButtonBlue)
		row.AddTextButton(board[2][i], fmt.Sprintf("%d %d", 2, i), object.ButtonBlue)
	}

	row := keyboard.AddRow()
	row.AddTextButton("Закончить", "stop", object.ButtonRed)

	messageToSend.Keyboard(keyboard)
	messageToSend.Message("Новая игра")

	bot.SendMessage(messageToSend.Params)

	return
}
