package bot

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/blrrubik/vk_bot/internal/app/service"
	"regexp"
	"strconv"
	"strings"
)

func (bot *Bot) receiveTicTacToeCommands(obj events.MessageNewObject, user *User) {
	ticTacToe := user.GameService.Services["TTT"].(service.TicTacToeInterface)

	gameCommand := strings.Trim(obj.Message.Payload, "\"")

	messageToSend := params.NewMessagesSendBuilder()
	messageToSend.RandomID(0)
	messageToSend.PeerID(obj.Message.PeerID)

	if gameCommand == "stop" {
		ticTacToe.EndGame()

		messageToSend.Message("Игра закончена")

		bot.SendMessage(messageToSend.Params)
		bot.sendMenu(obj, user)

		return
	}

	re := regexp.MustCompile(`^[0-9]\s[0-9]$`)
	if re.MatchString(gameCommand) {
		row, err := strconv.Atoi(strings.Split(gameCommand, " ")[0])
		if err != nil {
			messageToSend.Message("Invalid input")
			bot.SendMessage(messageToSend.Params)

			return
		}

		col, err := strconv.Atoi(strings.Split(gameCommand, " ")[1])
		if err != nil {
			messageToSend.Message("Invalid input")
			bot.SendMessage(messageToSend.Params)

			return
		}

		board, status, error := ticTacToe.PlayerTurn(row, col)
		if error != nil {
			messageToSend.Message(error.Message)
			bot.SendMessage(messageToSend.Params)

			return
		}

		keyboard := object.NewMessagesKeyboardInline()

		for i := 0; i < 3; i++ {
			row := keyboard.AddRow()
			row.AddTextButton(board[0][i], fmt.Sprintf("%d %d", 0, i), object.ButtonBlue)
			row.AddTextButton(board[1][i], fmt.Sprintf("%d %d", 1, i), object.ButtonBlue)
			row.AddTextButton(board[2][i], fmt.Sprintf("%d %d", 2, i), object.ButtonBlue)
		}

		row2 := keyboard.AddRow()
		row2.AddTextButton("Закончить", "stop", object.ButtonRed)

		if status != "" {
			messageToSend.Message(status)
			messageToSend.Keyboard(keyboard)

			bot.SendMessage(messageToSend.Params)

			bot.startTTT(obj, user)

			return
		}

		messageToSend.Message("Твой ход")
		messageToSend.Keyboard(keyboard)
		bot.SendMessage(messageToSend.Params)

		return
	}
}
