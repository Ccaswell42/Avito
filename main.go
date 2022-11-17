package main

import (
	"avito/config"
	"avito/server"
	"avito/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := storage.ConnectToDb(conf)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{}

	go func() {
		err := server.LaunchServer(conf, srv, db)
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
		fmt.Println("server completed successfully")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err = srv.Shutdown(closeCtx)
	if err != nil {
		log.Println("fail shutdown:", err)
	}
	log.Println("Shutdown completed successfully")
}
