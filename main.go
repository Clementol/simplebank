package main

import (
	"database/sql"
	"log"

	"github.com/Clementol/simplebank/api"
	db "github.com/Clementol/simplebank/db/sqlc"
	"github.com/Clementol/simplebank/util"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	config, err := util.LoadConfig(".")
	viper.WatchConfig()

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
