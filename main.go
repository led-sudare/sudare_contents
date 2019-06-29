package main

import (
	"flag"
	"time"

	"sudare_contents/lib/content"
	"sudare_contents/lib/util"

	log "github.com/cihub/seelog"
	zmq "github.com/zeromq/goczmq"
)

type Configs struct {
	ZmqTarget string `json:"zmqTarget"`
}

func NewConfigs() Configs {
	return Configs{
		ZmqTarget: "0.0.0.0:5510",
	}
}

func main() {
	configs := NewConfigs()
	util.ReadConfig(&configs)

	var (
		optInputPort = flag.String("r", configs.ZmqTarget, "Specify IP and port of server zeromq SUB running.")
	)

	flag.Parse()

	endpoint := "tcp://" + *optInputPort
	log.Info("New Pub: ", endpoint)
	zmqsock := zmq.NewSock(zmq.Pub)
	err := zmqsock.Connect(endpoint)
	if err != nil {
		panic(err)
	}
	defer zmqsock.Destroy()

	c := content.NewContentSinLine()
	t := time.NewTicker(50 * time.Millisecond) // 3秒おきに通知
	defer t.Stop()                             // タイマを止める。

	for {
		select {
		case <-t.C:
			zmqsock.SendFrame(c.GetFrame(), zmq.FlagNone)
			log.Info("Send Frame.")
		}
	}
}
