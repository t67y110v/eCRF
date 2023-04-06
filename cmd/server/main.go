package main

import (
	"github.com/t67y110v/web/internal/app/config"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/server"
)

func main() {
	logging.Init()
	l := logging.GetLogger()
	l.Infoln("Config initialization")
	config, err := config.LoadConfig()
	if err != nil {

		l.Fatal(err)
	}
	l.Infoln("Starting api server on addr :%s", config.Port)
	if err := server.Start(&config); err != nil {
		l.Fatal(err)
	}
}
