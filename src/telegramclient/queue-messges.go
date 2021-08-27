package telegramclient

import (
	"errors"
	"sync"
)

type QueueMessagesDoPop bool
type QueueMessagesReceiver func(message ClientMessage, err error) QueueMessagesDoPop

var (
	QueueMessagesErrorEmpty     = errors.New("Queue messages empty")
	QueueMessagesErrorIsntExist = errors.New("Message isn't exist in Queue")
)

type queueMessages struct {
	mx   *sync.RWMutex
	queue []ClientMessage
}

func newQueueMessages() *queueMessages {
	return &queueMessages{
		mx:   new(sync.RWMutex),
		queue: make([]ClientMessage, 0),
	}
}

func (this *queueMessages) Push(message ClientMessage) {
	defer this.mx.Unlock()
	this.mx.Lock()
	this.queue = append(this.queue, message)
}

func (this *queueMessages) pop(index int) {
	var (
		queue = make([]ClientMessage, 0)
	)
	for i, m := range this.queue {
		if i == index {
			continue
		}
		queue = append(queue, m)
	}
	this.queue = queue
}

func (this *queueMessages) Size() int {
	return len(this.queue)
}

func (this *queueMessages) GetNext(receiver QueueMessagesReceiver) {
	defer this.mx.Unlock()
	this.mx.Lock()
	if this.Size() == 0 {
		receiver(ClientMessage{}, QueueMessagesErrorEmpty)
		return
	}
	doPop := receiver(this.queue[0], nil)
	if doPop {
		this.pop(0)
	}
}
