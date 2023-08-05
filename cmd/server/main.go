package main

import (
	"github.com/t67y110v/web/internal/app/config"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/server"
)

//@title eCRF API
//@version 0.1
//@description API Server for eCRF Application

// @host 127.0.0.1:4000
// @basePath /

// @securityDefinitions.apikey  BasicAuth
// @in Cookie
// @name Cookie
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
