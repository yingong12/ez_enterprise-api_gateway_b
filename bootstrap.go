package main

import (
	"account_service/http"
	"account_service/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//bootstrap providers,以及routines
func bootStrap() (err error) {
	err, shutdownLogger := logger.Start()
	if err != nil {
		return
	}
	log.Println("Logger Started ")

	err, shutdownHttpServer := http.Start()
	if err != nil {
		return
	}
	log.Println("Httpserver started ")

	//wait for sys signals
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case sig := <-exitChan:
		log.Println("Doing cleaning works before shutdown...")
		shutdownLogger()
		shutdownHttpServer()
		log.Println("You abandoned me, bye bye", sig)
	}
	return
}
