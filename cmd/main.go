package main

import (
	"github.com/blrrubik/vk_bot/internal/app/bot"
	"github.com/blrrubik/vk_bot/internal/app/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	cfg := config.NewConfig()

	b, readError := os.ReadFile("./config/config.yml")
	if readError != nil {
		log.Fatal("Error on reading config")
	}

	if err := yaml.Unmarshal(b, &cfg); err != nil {
		log.Fatal(err)
	}

	vkBot := bot.NewBot(cfg)

	vkBot.Start()
}
