package main

import (
	"flag"
	"net/http"

	"sudare_contents/lib/content"
	"sudare_contents/lib/util"
	"sudare_contents/lib/webapi"
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

	sender := content.NewContentSender(endpoint)
	sender.SetContentToPlay(content.NewContentSinLine())

	webapi.SetUpWebAPIforCommon(sender)

	http.ListenAndServe(":5001", nil)
}
