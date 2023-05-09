package service

import (
	"fmt"
	"github.com/blrrubik/vk_bot/internal/app/model"
	"math/rand"
)

type RockPaperScissorsService struct {
	started bool
}

func NewRocPaperScissorsService() RockPaperScissorsInterface {
	return &RockPaperScissorsService{
		started: false,
	}
}

// IsWin
//
//	1 - rock
//	2 - paper
//	3 - scissors
//
// /**
func (service *RockPaperScissorsService) IsWin(playerChoice int) string {
	computerChoice := rand.Intn(3) + 1

	response := fmt.Sprintf("Ты выбрал: %s, компьютер выбрал: %s.",
		getStringFromChoice(playerChoice),
		getStringFromChoice(computerChoice))

	if playerChoice == computerChoice {
		return fmt.Sprintf("%s %s", response, model.TIE)
	} else if (playerChoice == 1 && computerChoice == 2) ||
		(playerChoice == 2 && computerChoice == 3) ||
		(playerChoice == 3 && computerChoice == 1) {
		return fmt.Sprintf("%s %s", response, model.PLAYER_WIN)
	} else {
		return fmt.Sprintf("%s %s", response, model.COMPUTER_WIN)
	}

}

func (service *RockPaperScissorsService) EndGame() {
	service.started = false
}

func (service *RockPaperScissorsService) IsStarted() bool {
	return service.started
}

func (service *RockPaperScissorsService) StartGame() {
	service.started = true
}

func getStringFromChoice(choice int) string {
	switch choice {
	case 1:
		return "Камень"
	case 2:
		return "Ножницы"
	case 3:
		return "Бумага"
	default:
		return ""
	}
}
