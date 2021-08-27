package telegramclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

const (
	TelegramApiHost    = "https://api.telegram.org/bot"
	SetWebHookEndpoint = "setWebhook"
)

type Bot struct {
	queue           *queueMessages
	webhookEndpoint string
	token           string
	receiver        QueueMessagesReceiver
}

func NewBot(token string, messageReceiver QueueMessagesReceiver) *Bot {
	this := &Bot{
		token:    token,
		queue:    newQueueMessages(),
		receiver: messageReceiver,
	}
	go this.messageBroadcaster()
	return this
}

func (this *Bot) InitWebhook(host string, engine *gin.Engine) (*http.Response, error) {
	webhookEndpoint := fmt.Sprintf(
		"bot/update/%s",
		this.token,
	)
	data := &jsonModelSubscribeWebhook{
		URL: fmt.Sprintf(
			"%s/%s",
			host,
			webhookEndpoint,
		),
	}
	bts, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(bts)
	url := fmt.Sprintf(
		"%s%s/%s",
		TelegramApiHost,
		this.token,
		SetWebHookEndpoint,
	)
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		return response, err
	}
	if response.StatusCode != http.StatusOK {
		return response, errors.New("Webhook unregistered")
	}
	this.webhookEndpoint = webhookEndpoint
	this.getWebhookUpdate(engine)
	log.Println(
		fmt.Sprintf(
			"Application was subscribed on webhook [%s] and will receive updates by URL: %s",
			url,
			data.URL,
		),
	)
	return nil, nil
}

func (this *Bot) GetFilePath(fileUniqueId string) (string, error) {
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/getFile?file_id=%s",
		this.token,
		fileUniqueId,
	)
	response, err := http.Get(
		url,
	)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	var (
		getFileResult = new(jsonGetFile)
	)
	err = json.NewDecoder(response.Body).Decode(getFileResult)
	log.Println(*getFileResult)
	if err != nil {
		return "", err
	}
	if !getFileResult.OK {
		return "", errors.New(fmt.Sprintf(
			"Get file with ID : [%s] has non result",
			fileUniqueId,
		))
	}
	if getFileResult.Result == nil || getFileResult.Result.FilePath == "" {
		return "", errors.New(fmt.Sprintf(
			"Get file with ID : [%s] has non result",
			fileUniqueId,
		))
	}
	return getFileResult.Result.FilePath, nil
}

func (this *Bot) DownloadFile(filePath string) (io.Reader, error) {
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/%s",
		this.token,
		filePath,
	)
	response, err := http.Get(
		url,
	)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	err = response.Write(&buffer)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buffer.Bytes()), nil
}

func (this *Bot) getWebhookUpdate(engine *gin.Engine) {
	engine.POST(this.webhookEndpoint, func(c *gin.Context) {
		var (
			message ClientMessage
		)
		if err := c.BindJSON(&message); err != nil {
			log.Println(
				fmt.Sprintf(
					"An error occurred while receiving an update from bot [%s], error: [%s]",
					this.token,
					err.Error(),
				),
			)
			return
		} else {
			go func(message ClientMessage) {
				this.queue.Push(message)
			}(message)
		}
	})
}

func (this *Bot) messageBroadcaster() {
	for {
		runtime.Gosched()
		if this.queue.Size() == 0 {
			continue
		}
		this.queue.GetNext(this.receiver)
	}
}
