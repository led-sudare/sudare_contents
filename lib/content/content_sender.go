package content

import (
	"reflect"
	"time"

	log "github.com/cihub/seelog"
	zmq "github.com/zeromq/goczmq"
)

type ContentSender interface {
	SetContentToPlay(contents []CylinderContent, interval time.Duration)
	Enable(enable bool)
	IsEnable() bool
}

type conentSenderImpl struct {
	con      chan CylinderContent
	enable   chan bool
	isEnable bool
}

func NewContentSender(endpoint string) ContentSender {

	sender := new(conentSenderImpl)

	sender.con = make(chan CylinderContent)
	sender.enable = make(chan bool)
	sender.isEnable = true

	log.Info("New Pub: ", endpoint)
	zmqsock := zmq.NewSock(zmq.Pub)
	err := zmqsock.Connect(endpoint)
	if err != nil {
		panic(err)
	}
	go worker(zmqsock, sender)
	return sender
}

func (s *conentSenderImpl) SetContentToPlay(contents []CylinderContent, interval time.Duration) {
	for {
		for _, c := range contents {
			s.con <- c
			time.Sleep(interval)
		}
	}
}

func (s *conentSenderImpl) Enable(enable bool) {
	s.isEnable = enable
	s.enable <- enable
}

func (s *conentSenderImpl) IsEnable() bool {
	return s.isEnable
}

func worker(zmqsock *zmq.Sock,
	sender *conentSenderImpl) {

	defer zmqsock.Destroy()

	var c CylinderContent
	enable := sender.IsEnable()
	t := time.NewTicker(50 * time.Millisecond) // 3秒おきに通知
	defer t.Stop()                             // タイマを止める。

	for {
		select {
		case c = <-sender.con:
			log.Info("change content: ", reflect.TypeOf(c))
		case enable = <-sender.enable:
			log.Info("enable:", enable)
		case <-t.C:
			if c != nil && enable {
				zmqsock.SendFrame(c.GetFrame(), zmq.FlagNone)
				log.Info("Send Frame.")
			}
		}
	}
}
