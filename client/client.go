package client

import (
	"encoding/json"
	"fmt"
	"github.com/Dadard29/go-common/config"
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-common/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

var logger = log.GetLogger("Navis client", log.DEBUG)

func ConnectorNew(isHost bool, token string) (*Connector, error) {
	var host string
	var portInput string

	c := config.Config.CnfFile.GetMap("apiServer")

	if isHost {
		host = c["host"]
		portInput = c["port"]
	} else {
		host = utils.ReadlineNew("Enter IP address of the host:").GetInput()

		portDefault := c["port"]
		portPs1 := fmt.Sprintf("Enter the port (%s):", portDefault)
		portInput = utils.ReadlineNew(portPs1).GetInput()
		if portInput == "" {
			portInput = portDefault
		}

		token = utils.ReadlineNew("Enter the access token:").GetInput()
	}

	port, err := strconv.Atoi(portInput)

	return &Connector{
		httpClient: &http.Client{},
		host:       host,
		port:       port,
		token:      token,
	}, err
}

type Connector struct {
	httpClient *http.Client
	host string
	port int
	token string
}

func (c Connector) getUrl(route string) string {
	return fmt.Sprintf("http://%s:%d%s", c.host, c.port, route)
}

func (c Connector) TestConnection() bool {
	route := "/health"

	resp, err := c.httpClient.Get(c.getUrl(route))
	utils.CheckErr(err, logger)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	utils.CheckErr(err, logger)

	var health HealthResponse
	utils.CheckErr(json.Unmarshal(body, &health), logger)

	return health.Status
}
