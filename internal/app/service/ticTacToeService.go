package service

import (
	"github.com/blrrubik/vk_bot/internal/app/model"
	"github.com/blrrubik/vk_bot/internal/app/utils"
)

type TicTacToeService struct {
	Started bool
	board   [][]string
}

func NewTicTacToeService() TicTacToeInterface {
	return &TicTacToeService{
		Started: false,
	}
}

func (service *TicTacToeService) IsStarted() bool {
	return service.Started
}

func (service *TicTacToeService) StartGame() [][]string {
	board := [][]string{{"▢", "▢", "▢"}, {"▢", "▢", "▢"}, {"▢", "▢", "▢"}}

	service.board = board
	service.Started = true

	return service.board
}

func (service *TicTacToeService) EndGame() {
	service.Started = false
}

func (service *TicTacToeService) PlayerTurn(row, col int) ([][]string, string, *model.Error) {

	if service.board[row][col] == "▢" {
		service.board[row][col] = "X"

		status := service.isWin("X")
		if status != "" {
			return service.board, status, nil
		}

		return service.board, service.ComputerTurn(), nil
	}

	return nil, "", &model.Error{
		Message: "Illegal move",
	}
}

func (service *TicTacToeService) ComputerTurn() string {
	row, col := utils.FindBestMove(service.board)

	service.board[row][col] = "O"

	return service.isWin("O")
}

func (service *TicTacToeService) isWin(player string) string {
	if service.checkTie() {
		return model.TIE
	}

	if service.checkWinner(player) {
		if player == "X" {
			return model.PLAYER_WIN
		}
		return model.COMPUTER_WIN
	}

	return ""
}

func (service *TicTacToeService) checkWinner(player string) bool {
	board := service.board

	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	for i := 0; i < 3; i++ {
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}

	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func (service *TicTacToeService) checkTie() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if service.board[i][j] == "▢" {
				return false
			}
		}
	}

	return true
}
