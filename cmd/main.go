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

// запуск логов; загрузка конфигов; запуск бд и сервера.

func main() {
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)
	log.Debug("log is loaded")

	var cfg config.MainConfig
	cfg.ConfigMustLoad("local")
	log.Debug("config is loaded")

	var PGDB database.PostgresDatabase
	PGDB.Run(cfg)
	var MNGDB database.MongoDatabase
	MNGDB.Run(cfg)
	var SQLTEDB database.SqliteDatabase
	SQLTEDB.Run(cfg)
	log.Debug("databases is loaded")

	//go metrics.MetricListen(":8082")

	var application app.App
	go application.NewServer(cfg.Server.Port)
	log.Debug("server is loaded")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop()
}
