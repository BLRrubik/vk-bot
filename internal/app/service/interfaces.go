package service

import "github.com/blrrubik/vk_bot/internal/app/model"

type TicTacToeInterface interface {
	Game
	PlayerTurn(row, col int) ([][]string, string, *model.Error)
	ComputerTurn() string
	StartGame() [][]string
	EndGame()
}

type RockPaperScissorsInterface interface {
	Game
	IsWin(playerGuess int) string
	StartGame()
	EndGame()
}

type Game interface {
	IsStarted() bool
}
