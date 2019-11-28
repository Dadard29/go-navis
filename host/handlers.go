package host

import (
	"github.com/Dadard29/go-common/api"
	"github.com/Dadard29/go-common/config"
	"github.com/Dadard29/go-common/utils"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	apiInfos := api.GetApiInfos(config.Config.CnfFile.GetMap("apiInfos")).ToMap()

	content := api.ResponseContent{
		Content:map[string]map[string]interface{} {
			"ApiInfos": apiInfos,
		},
	}

	err := api.EncodeAndWriteResponse(w, true, "health check performed", content)
	utils.CheckErr(err, api.Logger)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

}
