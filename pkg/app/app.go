package app

import (
	"context"
	"gomoku/pkg/http/handler"
	"gomoku/pkg/http/server"
	"os"
)

func Run(ctx context.Context) error {
	// add config
	hlr := handler.NewHandler()
	router := hlr.InitRoutes()

	// todo pass addr from config
	srv := server.NewHttpServer(ctx, "127.0.0.1:8080", router)
	go srv.Run()
	srv.Shutdown(os.Interrupt)

	return nil
}
