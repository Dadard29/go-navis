package init_config

import (
	"github.com/Dadard29/go-common/api"
	"github.com/Dadard29/go-common/config"
	"github.com/Dadard29/go-common/log"
	"os"
)

func InitConfigLogger() {
	// init conf
	homePath := os.Getenv("GOPATH")
	packagePath := "/src/github.com/Dadard29/go-navis"
	configPath := "/config/config.json"

	config.SetConfig(homePath + packagePath + configPath)

	// init logger
	api.SetLogger("Navis Host", log.ERROR)
}
