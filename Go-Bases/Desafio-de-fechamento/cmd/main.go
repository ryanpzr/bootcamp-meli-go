package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/repository"
	"github.com/bootcamp-go/desafio-go-bases/internal/service"
)

func main() {
	repository := repository.NewRepository()
	var travelslist []service.Travel
	service := service.NewService(repository, travelslist)

	travels, err := service.GetListWithTravels("../tickets.csv")
	if err != nil {
		fmt.Println(err)
	}

	for travel := range travels {
		fmt.Printf("%+v\n", travels[travel])
	}

	totalTickets, err := service.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nForam encontradas %d viagens de acordo com o destino solicitado.\n", totalTickets)

	mapStatis, err := service.GetCountByPeriod()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nViagens feitas durante períodos: \nInicio da manhã: %d\nManhã: %d\nTarde: %d\nNoite: %d\n", mapStatis["IM"], mapStatis["M"], mapStatis["T"], mapStatis["N"])

	averageTravels, err := service.AverageDestination("Brazil", 23)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%.2f%% das pessoas viajaram para o país Brazil no dia 23, com relação ao restante dos viajantes.\n", averageTravels)
}
