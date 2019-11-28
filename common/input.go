package common

import (
	"fmt"
	"github.com/Dadard29/go-common/utils"
	"os"
)

func GetUsername() (string, error) {
	def, err := os.Hostname()
	if err != nil {
		return "", err
	}
	ps1 := fmt.Sprintf("Enter your username (%s):", def)
	u := utils.ReadlineNew(ps1).GetInput()
	pu := &u
	if *pu == "" {
		*pu = def
	}
	return *pu, nil
}
