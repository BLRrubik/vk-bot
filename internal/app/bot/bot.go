package bot

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/blrrubik/vk_bot/internal/app/config"
	"github.com/blrrubik/vk_bot/internal/app/service"
	"log"
)

type Bot struct {
	vk *api.VK
}

type User struct {
	inMenu      bool
	inOtherMenu bool
	inRules     bool
	GameService *service.GamesService
	// другие переменные
}

func NewBot(config *config.Config) *Bot {
	return &Bot{
		vk: api.NewVK(config.Token),
	}
}

func (bot *Bot) Start() {
	vk := bot.vk

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	users := make(map[int]*User)

	// Канал для передачи сообщений от пользователей в главную горутину
	msgCh := make(chan events.MessageNewObject)

	// Главная горутина обрабатывает сообщения из канала и запускает новые горутины для каждого пользователя
	go func() {
		for obj := range msgCh {
			// Получаем информацию о пользователе
			userID := obj.Message.FromID
			user, ok := users[userID]
			if !ok {
				// Если пользователь еще не существует, создаем новую запись
				user = &User{
					inMenu:      true,
					GameService: service.NewGamesService(),
				}
				users[userID] = user
			}

			go func(obj events.MessageNewObject) {
				isStarted, game := user.GameService.IsGameStarted()

				if obj.Message.Text == "Начать" {
					messageToSend := params.NewMessagesSendBuilder()
					messageToSend.RandomID(0)
					messageToSend.PeerID(obj.Message.PeerID)
					messageToSend.Message("Привет. Я мини игровой бот, созданный для отбора на стажировку VK. " +
						"\nЧто я умею? Не много, но и не мало. Я могу сыграть с тобой в крестики-нолики или " +
						"в камень-ножницы-бумага. Ты можешь также посмотреть правила этих игры по кнопке в меню. \n" +
						"Что придумать для 4 кнопки мне не хватило фантазии. Поэтому там ты сможешь посмотреть время и " +
						"сделать ping-pong со мной. Считай тестовая вкладка. \n" +
						"Надеюсь ты со мной не заскучаешь))")

					bot.SendMessage(messageToSend.Params)

					bot.sendMenu(obj, user)
				}

				if isStarted {
					bot.receiveGameCommands(obj, game, user)
					return
				}
				if user.inRules {
					bot.receiveRulesCommands(obj, user)
					return
				}
				if user.inOtherMenu {
					bot.receiveOtherMenuCommands(obj, user)
					return
				}
				if user.inMenu {
					bot.receiveMenuCommands(obj, user)
				}
			}(obj)
		}
	}()

	// Обработчик новых сообщений добавляет их в канал
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		msgCh <- obj
	})

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}

func (bot *Bot) SendMessage(params api.Params) {
	_, err := bot.vk.MessagesSend(params)
	if err != nil {
		log.Fatal(err)
	}
}
