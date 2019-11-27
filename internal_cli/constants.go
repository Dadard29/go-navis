package internal_cli

// commands
const (
	QUIT = "exit"
	PING = "ping"
	CONNECT = "connect"
	HELP = "help"
	REGISTER = "register"
	DISCONNECT = "disconnect"
)

// responses
const (
	UNKNOWN_COMMAND = "Unknown command, see help"
	CONNECTED       = "Connected to server"
	NOT_CONNECTED   = "Not connected"
	DISCONNECTED 	= "Disconnected"
	CONNECTION_INFOS_NOT_SET = "You're not connected, use the <connect> command"
)


