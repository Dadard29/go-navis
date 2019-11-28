package internal_cli

import (
	"errors"
	"fmt"
	"github.com/Dadard29/go-navis/client"
	"github.com/Dadard29/go-navis/common"
)


func ping(c **client.Connector) (string, error) {
	if *c == nil {
		return "", errors.New(CONNECTION_INFOS_NOT_SET)
	}

	if (*c).TestConnection() {
		return CONNECTED, nil
	} else {
		return NOT_CONNECTED, nil
	}
}

func connect(c **client.Connector) (string, error) {
	*c = client.ConnectorNewAsClient()
	if c == nil {
		return "", errors.New(common.FAILED_TO_CONNECT)
	}
	return ping(c)
}

func disconnect(c **client.Connector) (string, error) {
	*c = nil
	return DISCONNECTED, nil
}

func register(c **client.Connector) (string, error) {
	if *c == nil {
		return "", errors.New(CONNECTION_INFOS_NOT_SET)
	}

	err := (*c).Register()
	if err != nil {
		return "", errors.New(common.FAILED_TO_REGISTER)
	}

	return "ok", nil
}

func help() (string, error) {
	return fmt.Sprintf("Commands availables:\n" +
		"\t%s: exit the CLI\n" +
		"\t%s: ping the game server\n" +
		"\t%s: connect to a game server\n" +
		"\t%s: disconnect from the game server\n" +
		"\t%s: register to the game server with the access token\n" +
		"\t%s: show this help",
		QUIT, PING, CONNECT, DISCONNECT, REGISTER, HELP,
	), nil
}

func ProcessInput(input string, c **client.Connector) (string, error) {
	if input == PING {
		return ping(c)
	} else if input == CONNECT {
		return connect(c)
	} else if input == DISCONNECT {
		return disconnect(c)
	} else if input == REGISTER {
		return register(c)
	} else if input == HELP {
		return help()
	}

	return "", errors.New(UNKNOWN_COMMAND)
}
