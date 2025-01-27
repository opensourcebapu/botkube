package main

import (
	"fmt"

	"github.com/infracloudio/botkube/pkg/bot"
	"github.com/infracloudio/botkube/pkg/config"
	"github.com/infracloudio/botkube/pkg/controller"
	log "github.com/infracloudio/botkube/pkg/logging"
	"github.com/infracloudio/botkube/pkg/utils"
)

func main() {
	log.Logger.Info("Starting controller")
	Config, err := config.New()
	if err != nil {
		log.Logger.Fatal(fmt.Sprintf("Error in loading configuration. Error:%s", err.Error()))
	}

	if Config.Communications.Slack.Enabled {
		log.Logger.Info("Starting slack bot")
		sb := bot.NewSlackBot()
		go sb.Start()
	}

	if Config.Communications.Mattermost.Enabled {
		log.Logger.Info("Starting mattermost bot")
		mb := bot.NewMattermostBot()
		go mb.Start()
	}

	if Config.Settings.UpgradeNotifier {
		log.Logger.Info("Starting upgrade notifier")
		go controller.UpgradeNotifier(Config)

	}

	// Init KubeClient, InformerMap and start controller
	utils.InitKubeClient()
	utils.InitInformerMap()
	controller.RegisterInformers(Config)
}
