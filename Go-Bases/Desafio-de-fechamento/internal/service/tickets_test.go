package service_test

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/repository"
	"github.com/bootcamp-go/desafio-go-bases/internal/service"
)

func TestGetTotalTickets(t *testing.T) {
	service := getMock()

	totalTickets, err := service.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}

	expected := 3

	if expected != totalTickets {
		t.Errorf("Resultado retornado não corresponde ao esperado. Esperado: %d Atual: %d", expected, totalTickets)
	}
}

func TestGetCountByPeriod(t *testing.T) {
	service := getMock()

	mapStatis, err := service.GetCountByPeriod()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mapStatis)

	expectedInicioDaManha := 30
	expectedManha := 20
	expectedTarde := 30
	expectedNoite := 20

	if mapStatis["IM"] != expectedInicioDaManha {
		t.Errorf("Resultado retornado não corresponde ao esperado. Esperado: %d Atual: %d", expectedInicioDaManha, mapStatis["IM"])
	}
	if mapStatis["M"] != expectedManha {
		t.Errorf("Resultado retornado não corresponde ao esperado. Esperado: %d Atual: %d", expectedManha, mapStatis["M"])
	}
	if mapStatis["T"] != expectedTarde {
		t.Errorf("Resultado retornado não corresponde ao esperado. Esperado: %d Atual: %d", expectedTarde, mapStatis["T"])
	}
	if mapStatis["N"] != expectedNoite {
		t.Errorf("Resultado retornado não corresponde ao esperado. Esperado: %d Atual: %d", expectedNoite, mapStatis["N"])
	}
}

func TestAverageDestination(t *testing.T) {
	service := getMock()
	averageTravels, err := service.AverageDestination("Poland", 22)
	if err != nil {
		fmt.Println(err)
	}

	expected := 40.00

	if averageTravels != expected {
		t.Errorf("Resultado retornado não corresponde ao esperado. Esperado: %f Atual: %f", expected, averageTravels)
	}
}

func getMock() service.Service {
	repo := repository.NewRepository()
	var travels []service.Travel
	service := service.NewService(repo, travels)

	_, err := service.GetListWithTravels("/Users/ryplima/Developer/Personal/Vscode/desafio-go-bases/test/tickets-test.csv")
	if err != nil {
		fmt.Println(err)
	}

	return *service
}
