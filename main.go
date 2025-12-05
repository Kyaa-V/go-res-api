package main

import (
	"context"
	"fmt"

	"github.com/Kyaa-V/go-res-api/application"
)

func main() {
	fmt.Println("application setup")
	app := application.New()

	fmt.Println("starting aplication")
	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("Error starting the application:", err)
	}

}
