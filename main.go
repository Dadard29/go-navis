package main

import (
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-common/utils"
	"github.com/Dadard29/go-navis/host"
	"github.com/Dadard29/go-navis/internal_cli"
	"github.com/urfave/cli"
	"os"
)
var logger = log.GetLogger("Navis", log.DEBUG, 0)

func main() {

	app := cli.NewApp()
	app.Name = "go-navis"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:        "host",
			Aliases:     []string{"H"},
			Usage:       "use it to start hosting a game",
			Description: "Host the game on your computer",
			Action: host.StartHost,
		}, {
			Name:        "client",
			Aliases:     []string{"c"},
			Usage:       "use it to connect to a game",
			Description: "Connect to a game hosted elsewhere",
			Action: internal_cli.StartClientCli,
		},
	}

	utils.CheckErr(app.Run(os.Args), logger)
}
