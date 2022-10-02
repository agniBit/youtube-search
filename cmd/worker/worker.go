package main

import (
	"flag"

	"github.com/agniBit/youtube-search/pkg/worker"
)

func main() {
	confPath := flag.String("p", "./cmd/conf.local.yaml", "Path to config file")
	flag.Parse()
	_ = worker.Start(*confPath)
}
