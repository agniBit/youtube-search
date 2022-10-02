package main

import (
	"flag"

	"github.com/agniBit/youtube-search/pkg/job"
)

func main() {
	//flag lib is used to get the config path from command line too
	confPath := flag.String("p", "./cmd/config.local.yaml", "Path to config file")
	flag.Parse()

	_ = job.Start(*confPath)
}
