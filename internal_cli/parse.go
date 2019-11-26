package internal_cli

import (
	"fmt"
	"github.com/Dadard29/go-navis/client"
)


func ping(c *client.Connector) (string, error) {
	if c == nil {
		return CONNECTION_INFOS_NOT_SET, nil
	}

	if c.TestConnection() {
		return CONNECTED, nil
	} else {
		return NOT_CONNECTED, nil
	}
}

func connect(c **client.Connector) (string, error) {
	var err error
	*c, err = client.ConnectorNew(false, "")
	if err != nil {
		return FAILED_TO_CONNECT, err
	}
	return ping(*c)
}

func help() (string, error) {
	return fmt.Sprintf("Commands availables:\n" +
		"\t%s: exit the CLI\n" +
		"\t%s: ping the game server\n" +
		"\t%s: connect to a game server\n" +
		"\t%s: show this help",
		QUIT, PING, CONNECT, HELP,
	), nil
}

func ProcessInput(input string, c *client.Connector) (string, error) {
	if input == PING {
		return ping(c)
	} else if input == CONNECT {
		return connect(&c)
	} else if input == HELP {
		return help()
	}

	return UNKNOWN_COMMAND, nil
}
