package main

import (
	"flag"
	"math/rand"
	"net/http"
	"time"

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
	util.InitColorUtil()
	rand.Seed(time.Now().UnixNano())

	configs := NewConfigs()
	util.ReadConfig(&configs)

	var (
		optInputPort = flag.String("r", configs.ZmqTarget, "Specify IP and port of server zeromq SUB running.")
	)

	flag.Parse()

	endpoint := "tcp://" + *optInputPort

	sender := content.NewContentSender(endpoint)
	contents := []content.CylinderContent{
		//		content.NewContentExLine(),
		content.NewContentSinWideLine(),
		content.NewContentSinLine(),
		content.NewContentCirWave(),
	}
	go sender.SetContentToPlay(contents, 4*time.Second)

	webapi.SetUpWebAPIforCommon(sender)

	http.ListenAndServe(":5004", nil)
}
