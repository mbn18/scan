package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mbn18/scan/config"
	"github.com/mbn18/scan/mapper/postgres"
	"github.com/mbn18/scan/server"
	"log"
)

func main() {
	log.Print("Starting HTTP server")

	conf := config.Get()
	dbConn, err := pgxpool.New(context.Background(), conf.GetConnString())
	if err != nil {
		log.Panic(err)
	}
	defer dbConn.Close()

	mapper := postgres.NewMapper(dbConn)
	server.Start(conf.Address, mapper)
}
