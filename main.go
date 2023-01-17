package main

import (
	"database/sql"
	"log"

	"github.com/LeGion013/banktemplate/api"
	db "github.com/LeGion013/banktemplate/db/sqlc"
	"github.com/LeGion013/banktemplate/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".") // "." - means current folder

	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSsource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
