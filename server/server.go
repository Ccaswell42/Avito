package server

import (
	"avito/config"
	"avito/server/handlers"
	"database/sql"
	"log"
	"net/http"
	"time"
)

func LaunchServer(conf *config.Config, srv *http.Server, db *sql.DB) error {
	mux := http.NewServeMux()

	d := handlers.Data{db}
	mux.HandleFunc("/balance", d.Balance)
	mux.HandleFunc("/replenish", d.Replenish)
	mux.HandleFunc("/reserve", d.Reserve)
	mux.HandleFunc("/unreserve", d.UnReserve)
	mux.HandleFunc("/revenue", d.Revenue)

	srv.Addr = conf.Port
	srv.Handler = mux
	srv.ReadTimeout = 10 * time.Second
	srv.WriteTimeout = 10 * time.Second

	log.Println("starting server at", conf.Port)
	return srv.ListenAndServe()
}
