package internal_cli

import (
	"fmt"
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-common/utils"
	"github.com/Dadard29/go-navis/client"
	initConfig "github.com/Dadard29/go-navis/config"
	"github.com/urfave/cli"
)

var logger = log.GetLogger("Navis CLI", log.DEBUG)

func StartCli(isHost bool, token string) {
	r := utils.ReadlineNew("navis>")
	connector, err := client.ConnectorNew(isHost, token)

	if err != nil || !connector.TestConnection() {
		logger.Error(err.Error())
		logger.Warning(FAILED_TO_CONNECT)
		connector = nil
	}

	input := r.GetInput()
	for input != QUIT {
		output, err := ProcessInput(input, connector)
		utils.CheckErr(err, logger)
		fmt.Println(output)

		input = r.GetInput()
	}
}

func StartClientCli(c *cli.Context) error {
	initConfig.InitConfigLogger()

	StartCli(false, "")
	return nil
}
