package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/agniBit/youtube-search/pkg/api"
)

func main() {
	//flag lib is used to get the config path from command line too
	confPath := flag.String("p", "./cmd/config.local.yaml", "Path to config file")
	runMigration := flag.Bool("migrate", false, "Run migration")
	flag.Parse()

	if *runMigration {
		log.Println("Running migration...")
		if _, err := os.Stat("/usr/bin/migration/main"); err != nil {
			if os.IsNotExist(err) {
				log.Fatal("migration binary not found at /usr/bin/migration/main")
			} else {
				log.Fatal(err)
			}
			return
		}

		cmd := exec.Command("/usr/bin/migration/main")
		if err := cmd.Run(); err != nil {
			log.Fatal("unable to migrate database")
			log.Fatal(err)
		}
	}

	_ = api.Start(*confPath)
}
