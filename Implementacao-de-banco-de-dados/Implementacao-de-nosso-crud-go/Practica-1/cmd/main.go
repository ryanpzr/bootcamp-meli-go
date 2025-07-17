package main

import (
	"app/internal/application"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// env
	// ...

	// app
	// - config
	app := application.NewApplicationDefault("", "./docs/db/json/products.json")
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env: " + err.Error())
	}
	// - tear down
	defer app.TearDown()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
