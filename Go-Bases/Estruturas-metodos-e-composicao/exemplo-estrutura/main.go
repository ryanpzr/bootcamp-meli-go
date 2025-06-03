package main

import (
	"metodo/repository"
	"metodo/service"
)

func main() {
	repo := repository.Repository{}
	newService := service.NewService(repo)
	newService.Save("oi")
}
