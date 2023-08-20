package main

import (
	"context"
	"kmf-proxy/internal/server"
	"kmf-proxy/pkg/handler"
	"kmf-proxy/pkg/repository/httpclient"
	"kmf-proxy/pkg/repository/storage"
	"kmf-proxy/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := new(server.Server)

	httpCliRepo := httpclient.NewRepository()
	storageRepo := storage.NewRepository()
	proxyService := service.NewProxyService(httpCliRepo, storageRepo)
	handlers := handler.NewHandler(proxyService)

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}

	go func() {
		if err := srv.Run(port, handlers.Init()); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Printf("Server started at port: %s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("cound't shutdown the server: %s", err.Error())
	}

	log.Println("Successful shutdown")
}
