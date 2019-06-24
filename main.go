package main

import (
	"flag"
	"time"

	"sudare_contents/lib/content"

	log "github.com/cihub/seelog"
	zmq "github.com/zeromq/goczmq"
)

var (
	logVerbose   = flag.Bool("v", false, "output detailed log.")
	optInputPort = flag.String("r", "127.0.0.1:5563", "Specify IP and port of server main_realsense_serivce.py running.")
)

func main() {
	flag.Parse()

	endpoint := "tcp://" + *optInputPort
	zmqsock, err := zmq.NewPub(endpoint)
	if err != nil {
		panic(err)
	}
	err = zmqsock.Connect(endpoint)
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
