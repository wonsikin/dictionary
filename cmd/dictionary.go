package main

import (
	"flag"

	"github.com/CardInfoLink/log"

	"github.com/wonsikin/dictionary/src"
)

var (
	address  = flag.String("addr", ":7100", `KPass service address to listen on`)
	bindHost = flag.String("bindhost", "", `Bind a host name to KPass client, default to service address`)
	logPath  = flag.String("logpath", "", `KPass log file path - has to be a file, not directory, default to stdout`)
	devMode  = flag.Bool("dev", false, "Development mode, will use memory database as default")
)

func main() {
	flag.Parse()

	srv := src.NewApp(*address)

	log.Infof("Wapiti is served at %s", *address)
	log.Error(srv.ListenAndServe())
}
