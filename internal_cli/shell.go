package internal_cli

import (
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-common/utils"
	"github.com/Dadard29/go-navis/client"
	"github.com/Dadard29/go-navis/common"
	initConfig "github.com/Dadard29/go-navis/config"
	"github.com/chzyer/readline"
	"github.com/urfave/cli"
)

var logger = log.GetLogger("Navis CLI", log.DEBUG, 0)

func StartCli(isHost bool, token string) {
	r, err := readline.New("navis> ")
	utils.CheckErr(err, logger)
	defer r.Close()

	connector := client.ConnectorNew(isHost, token)

	if connector == nil {
		common.PrintError(common.FAILED_TO_CONNECT)
	} else {
		common.PrintResponse(CONNECTED)
	}

	input, err := r.Readline()
	utils.CheckErr(err, logger)
	for input != QUIT {
		output, err := ProcessInput(input, &connector)
		if err != nil {
			common.PrintError(err.Error())
		} else {
			common.PrintResponse(output)
		}

		input, err = r.Readline()
		utils.CheckErr(err, logger)
	}
}

func StartClientCli(c *cli.Context) error {
	initConfig.InitConfigLogger()

	StartCli(false, "")
	return nil
}
