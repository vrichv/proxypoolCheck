package main

import (
	"flag"
	"github.com/vrichv/proxypoolCheck/api"
	"github.com/vrichv/proxypoolCheck/config"
	"github.com/vrichv/proxypoolCheck/internal/app"
	"github.com/vrichv/proxypoolCheck/internal/cron"
	"log"
	"net/http"
)

var configFilePath = ""

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6061", nil)
	}()

	//Slog.SetLevel(Slog.DEBUG) // Print original pack log

	// fetch configuration
	flag.StringVar(&configFilePath, "c", "", "path to config file: config.yaml")
	flag.Parse()
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	err := config.Parse(configFilePath)
	if err != nil {
		log.Fatal(err, "\n\"Config file err. Exit\"")
		return
	}

	go app.InitApp()
	log.Printf("The program will run every %v minutes\n", config.Config.CronInterval)
	go cron.Cron()
	// Run
	api.Run()

}
