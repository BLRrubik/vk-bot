package service

type Services map[string]Game

type GamesService struct {
	Services
}

func NewGamesService() *GamesService {
	var services = make(map[string]Game)

	services["TTT"] = NewTicTacToeService()
	services["RPS"] = NewRocPaperScissorsService()

	return &GamesService{
		Services: services,
	}
}

func (service *GamesService) IsGameStarted() (bool, Game) {
	for _, v := range service.Services {
		if v.IsStarted() {
			return true, v
		}
	}

	return false, nil
}
