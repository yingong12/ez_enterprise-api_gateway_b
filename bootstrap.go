package main

import (
	"api_gateway_b/http"
	"api_gateway_b/library/env"
	"api_gateway_b/logger"
	"api_gateway_b/providers"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"api_gateway_b/library"

	"github.com/joho/godotenv"
)

//bootstrap providers,以及routines
func bootStrap() (err error) {
	//加载环境变量
	filePath := ".env"
	serverPort := 8680
	flag.StringVar(&filePath, "c", ".env", "配置文件")
	flag.IntVar(&serverPort, "p", 8680, "http端口")
	flag.Parse()
	if err = godotenv.Load(filePath); err != nil {
		return
	}
	log.Println("env loadded from file ", filePath)
	//logger
	err, shutdownLogger := logger.Start()
	if err != nil {
		return
	}
	log.Println("Logger Started ")
	//http client
	httpClients := []*library.HttpClientConfig{
		{
			Name:     "static_server",
			BaseURL:  `http://` + env.GetStringVal("LB_COMPANY_SERVICE"),
			Receiver: &providers.HttpClientCompanyService,
		}, {
			Name:     "static_server",
			BaseURL:  `http://` + env.GetStringVal("LB_ACCOUNT_SERVICE"),
			Receiver: &providers.HttpClientAccountService,
		},
	}
	for _, cfg := range httpClients {
		if cfg.Receiver == nil {
			return fmt.Errorf("config receiver cannot be nil")
		}
		*cfg.Receiver = library.NewHttpClient(cfg)
		(**cfg.Receiver).BaseURL = cfg.BaseURL
	}
	//http server
	err, shutdownHttpServer := http.Start(serverPort)
	if err != nil {
		return
	}
	log.Println("Httpserver started on port ", serverPort)
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
