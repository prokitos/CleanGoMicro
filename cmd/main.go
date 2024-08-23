package main

import (
	"modules/internal/app"
	"modules/internal/config"
	"modules/internal/database"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debug("log is loaded")

	var cfg config.MainConfig
	cfg.ConfigMustLoad("local")
	log.Debug("config is loaded")

	// одна функция которая и мигририует и другое и се
	var userDB database.UserDatabase
	userDB.OpenConnection(cfg)
	userDB.StartMigration()
	userDB.GlobalSet()

	var computerDB database.ComputerDatabase
	computerDB.OpenConnection(cfg)
	computerDB.StartMigration()
	computerDB.GlobalSet()

	var application app.App
	go application.NewServer(cfg.Server.Port)
	log.Debug("server is loaded")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop()
}
