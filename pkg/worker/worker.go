package worker

import (
	"fmt"

	utils "github.com/agniBit/youtube-search/utl/common"
	"github.com/agniBit/youtube-search/utl/config"
)

func Start(configPath string) error {
	cfg, err := config.Load(configPath)
	utils.CheckErr(err)
	fmt.Println(cfg)
	return nil
}
