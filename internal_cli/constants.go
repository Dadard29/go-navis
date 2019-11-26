package internal_cli

// commands
const (
	QUIT = "exit"
	PING = "ping"
	CONNECT = "connect"
	HELP = "help"
)

// responses
const (
	UNKNOWN_COMMAND = "Unknown command, see help"
	CONNECTED       = "Connected to server"
	NOT_CONNECTED   = "Not connected"
	CONNECTION_INFOS_NOT_SET = "You didn't provided connections information, use the <connect> command"
)

// error messages
const  (
	FAILED_TO_CONNECT = "Failed to connect !"
)
