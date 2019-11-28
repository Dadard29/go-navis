package client

import (
	"encoding/json"
	"fmt"
	"github.com/Dadard29/go-common/config"
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-common/utils"
	"github.com/Dadard29/go-navis/common"
	"io/ioutil"
	"net/http"
	"strconv"
)

var logger = log.GetLogger("Navis client", log.DEBUG, 0)

type Connector struct {
	httpClient *http.Client
	host string
	port int
	token string
	username string
}

func ConnectorNewAsClient() *Connector {
	conf := config.Config.CnfFile.GetMap("apiServer")

	host := utils.ReadlineNew("Enter IP address of the host:").GetInput()

	portDefault := conf["port"]
	portPs1 := fmt.Sprintf("Enter the port (%s):", portDefault)
	portInput := utils.ReadlineNew(portPs1).GetInput()
	pPortInput := &portInput
	if *pPortInput == "" {
		*pPortInput = portDefault
	}

	port, err := strconv.Atoi(*pPortInput)
	utils.CheckErr(err, logger)

	token := utils.ReadlineNew("Enter the access token:").GetInput()

	username, err := common.GetUsername()
	utils.CheckErr(err, logger)

	return ConnectorNew(host, port, token, username)
}

func ConnectorNewAsHost(token string) *Connector {
	host := config.Config.CnfFile.GetValue("apiServer", "host")
	portConfig := config.Config.CnfFile.GetValue("apiServer", "port")
	port, err := strconv.Atoi(portConfig)
	utils.CheckErr(err, logger)
	username, err := common.GetUsername()
	utils.CheckErr(err, logger)

	return ConnectorNew(host, port, token, username)
}

func ConnectorNew(host string, port int, token string, username string) *Connector {

	return &Connector{
		httpClient: &http.Client{},
		host:       host,
		port:       port,
		token:      token,
		username: 	username,
	}
}

func (c Connector) getUrl(route string) string {
	return fmt.Sprintf("http://%s:%d%s", c.host, c.port, route)
}

func (c Connector) TestConnection() bool {
	route := "/health"

	resp, err := c.httpClient.Get(c.getUrl(route))
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	utils.CheckErr(err, logger)

	var health HealthResponse
	utils.CheckErr(json.Unmarshal(body, &health), logger)

	return health.Status
}

func (c Connector) Register() error {
	return nil
}
