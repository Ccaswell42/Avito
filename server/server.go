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

	d := handlers.Data{DB: db, Port: conf.Port}
	mux.HandleFunc("/balance", d.Balance)
	mux.HandleFunc("/replenish", d.Replenish)
	mux.HandleFunc("/reserve", d.Reserve)
	mux.HandleFunc("/unreserve", d.UnReserve)
	mux.HandleFunc("/revenue", d.Revenue)
	mux.HandleFunc("/report", d.Report)
	mux.Handle("/", http.FileServer(http.Dir(".")))

	srv.Addr = conf.Port
	srv.Handler = mux
	srv.ReadTimeout = 10 * time.Second
	srv.WriteTimeout = 10 * time.Second

	log.Println("starting server at", conf.Port)
	return srv.ListenAndServe()
}
