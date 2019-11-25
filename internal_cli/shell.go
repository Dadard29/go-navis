package internal_cli

import (
	"fmt"
	"github.com/Dadard29/go-common/log"
	"github.com/Dadard29/go-common/utils"
)

var logger = log.GetLogger("Navis CLI", log.DEBUG)

func StartCli() {
	r := utils.ReadlineNew("navis>")
	input := r.GetInput()
	for input != QUIT {
		output, err := ProcessInput(input)

		utils.CheckErr(err, logger)
		fmt.Println(output)

		input = r.GetInput()
	}
}
