package host

import (
	"context"
	"github.com/Dadard29/go-common/api"
	"github.com/Dadard29/go-common/config"
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-navis/internal_cli"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"time"
)

var logger = log.GetLogger("Navis Host", log.DEBUG)

func StartHost(c *cli.Context) error {
	logger.Info("starting host...")

	// init logger
	api.SetLogger("Navis Host")

	// init conf
	homePath := os.Getenv("GOPATH")
	packagePath := "/src/github.com/Dadard29/go-navis"
	configPath := "/config/config.json"

	config.SetConfig(homePath + packagePath + configPath)

	apiService := GetHost()
	go apiService.RunServerSynchronous()

	internal_cli.StartCli()

	logger.Debug("Closing the host...")
	err := apiService.Srv.Shutdown(apiService.Context)
	if err != nil {
		return err
	}

	logger.Debug("Server closed")
	return nil
}

func GetHost() *api.Api {
	api.Logger.Info("Registering routes...")
	routes := map[string]func (w http.ResponseWriter, r *http.Request){
		"/health": healthHandler,
	}

	for i, _ := range routes {
		api.Logger.Debug(i)
	}

	var router = api.GetRouter(routes)

	api.Logger.Info("Configuring web server...")
	var server = api.GetServer(router, config.Config.CnfFile.GetMap("apiServer"))

	api.Logger.Info("Retrieving infos from config...")
	infos := api.GetApiInfos(config.Config.CnfFile.GetMap("apiInfos"))

	// create a context
	var wait time.Duration
	ctx, _ := context.WithTimeout(context.Background(), wait)

	return &api.Api{
		Srv:   server,
		Infos: infos,
		Context: ctx,
	}
}
