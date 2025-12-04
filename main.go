package main

import (
	"application"
	"context"
	"fmt"
)

func main(){
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("Error starting the application:", err)
	}

}