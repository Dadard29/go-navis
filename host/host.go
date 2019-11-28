package host

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/Dadard29/go-common/api"
	"github.com/Dadard29/go-common/config"
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-common/utils"
	"github.com/Dadard29/go-navis/client"
	init_config "github.com/Dadard29/go-navis/config"
	"github.com/Dadard29/go-navis/internal_cli"
	"github.com/Dadard29/go-navis/player"
	"github.com/urfave/cli"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var logger = log.GetLogger("Navis Host", log.DEBUG, 0)

var playerList player.PlayerList

func getLocalIpAddress() string {
	ifaces, err := net.Interfaces()
	utils.CheckErr(err, logger)
	//var ipAddrList []net.IPNet
	//var ipNetList []net.IPNet

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		utils.CheckErr(err, logger)
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func setAccessToken() string {
	rand.Seed(time.Now().UnixNano())
	key := string(rand.Intn(100))
	hash := sha256.New()
	hash.Write([]byte(key))

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func StartHost(c *cli.Context) error {
	var err error

	init_config.InitConfigLogger()

	logger.Info("starting host...")

	apiService := GetHost()
	apiService.RunServerAsynchronous()

	// init the access token
	logger.Info("generating access token...")
	token := setAccessToken()

	port := config.Config.CnfFile.GetValue("apiServer", "port")
	ip := getLocalIpAddress()

	connectInfosLog := fmt.Sprintf("host: %s, port: %s, access token: %s", ip, port, token)
	logger.Info(connectInfosLog)
	logger.Info("Send these information to the others players so they can connect")

	// start the CLI with a connector already connected to the started server
	connector := client.ConnectorNewAsHost(token)
	internal_cli.StartCli(connector)

	logger.Debug("Closing the host...")
	err = apiService.Srv.Shutdown(apiService.Context)
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
		"/register": registerHandler,
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
