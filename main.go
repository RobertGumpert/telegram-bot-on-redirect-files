package main

import (
	"application/src/config"
	"application/src/dropboxclient"
	"application/src/telegramclient"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

const (
	TelegramApiHost    = "https://api.telegram.org/bot"
	SetWebHookEndpoint = "setWebhook"
)

var (
	bot     *telegramclient.Bot
	dropbox *dropboxclient.DropboxClientWrapper
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	env, err := config.Reader(config.Production, config.Env)
	if err != nil {
		log.Fatal(err)
	}
	dropbox = dropboxclient.NewDropboxClient(env.DropboxToken)
	engine, run := getServerEngine(env.ApplicationPort)
	bot = telegramclient.NewBot(env.TelegramToken, messageReceiver)
	_, err = bot.InitWebhook(env.ApplicationHost, engine)
	if err != nil {
		log.Fatal(err)
	}
	run()
}

func messageReceiver(message telegramclient.ClientMessage, err error) telegramclient.QueueMessagesDoPop {
	if err != nil {
		log.Println(err)
		return telegramclient.QueueMessagesDoPop(false)
	}
	log.Println(*message.Message.Document)
	filePath, err := bot.GetFilePath(message.Message.Document.FileID)
	if err != nil {
		panic(err)
	}
	ioReader, err := bot.DownloadFile(filePath)
	if err != nil {
		panic(err)
	}
	meta, err := dropbox.Upload(
		message.Message.Document.FileName,
		ioReader,
	)
	if err != nil {
		panic(err)
	}
	log.Println(meta)
	return telegramclient.QueueMessagesDoPop(true)
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
