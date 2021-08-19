package main

import (
	"application/src/config"
	"application/src/telegram"
	"log"
	"runtime"
	"github.com/gin-gonic/gin"
)

const (
	TelegramApiHost    = "https://api.telegram.org/bot"
	SetWebHookEndpoint = "setWebhook"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	env, err := config.ReadENV()
	if err != nil {
		log.Fatal(err)
	}
	engine, run := getServerEngine(env.AppPort)
	bot := telegram.NewBot(env.TgToken, messageReceiver)
	_, err = bot.InitWebhook(env.Host, engine)
	if err != nil {
		log.Fatal(err)
	}
	run()
}

func messageReceiver(message telegram.ClientMessage, err error) telegram.QueueMessagesDoPop {
	if err != nil {
		log.Println(err)
		return telegram.QueueMessagesDoPop(false)
	}
	log.Println(*message.Message.Document)
	return telegram.QueueMessagesDoPop(true)
}

func getServerEngine(port string) (*gin.Engine, func()) {
	engine := gin.Default()
	return engine, func() {
		err := engine.Run(port)
		if err != nil {
			log.Fatal(err)
		}
	}
}
