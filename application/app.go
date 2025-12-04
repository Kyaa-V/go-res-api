package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: setupRouter(),
	}
	return app
}

func (a *App) Start(ctx context.Context){
	server := &http.Server{
		Addr : ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port :3000")

	return nil
}