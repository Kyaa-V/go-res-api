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
	fmt.Println("setup App")
	app := &App{
		router: setupRouter(),
	}
	fmt.Println("done setup App")
	return app
}

func (a *App) Start(ctx context.Context) error{
	fmt.Println("setup server")
	server := &http.Server{
		Addr : ":3000",
		Handler: a.router,
	}

	fmt.Println("starting Server")

	fmt.Println("Server is running on port :3000")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("error starting server:", err)
		panic(err)
	}

	return nil
}