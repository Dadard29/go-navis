package client

import (
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-navis/internal_cli"
	"github.com/urfave/cli"
)

var logger = log.GetLogger("Navis client", log.DEBUG)

func StartClient(c *cli.Context) error {
	logger.Info("starting client...")

	internal_cli.StartCli()

	return nil
}
