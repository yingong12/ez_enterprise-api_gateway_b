package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	http.Server
	ctx context.Context
}

func (s *HttpServer) run() (err error) {
	go func() {
		if err = s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			err = fmt.Errorf("http.Server.ListenAndServe: %w", err)
			return
		}
		err = nil
	}()
	return
}

// Start 启动http服务器
func Start(port int) (err error, shutdown func()) {
	ctx, cancel := context.WithCancel(context.TODO())
	shutdown = func() {
		log.Println("Shuting down http server")
		cancel()
	}
	server := &HttpServer{
		ctx: ctx,
	}
	//TODO:先写死端口
	server.Addr = fmt.Sprintf(":%d", port)
	server.Handler = loadRouter()

	if err = server.run(); err != nil {
		return
	}
	return
}
