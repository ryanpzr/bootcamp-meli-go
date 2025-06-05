package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/repository"
)

var (
	earlyMorningTrips int = 0
	morningTrips      int = 0
	afternoonTrips    int = 0
	nightTrips        int = 0
)

type Travel struct {
	Id         string
	Name       string
	Email      string
	Country    string
	TravelTime string
	Price      string
}

type Service struct {
	repo    repository.Repository
	travels []Travel
}

func NewService(repo repository.Repository, travels []Travel) *Service {
	return &Service{repo: repo, travels: travels}
}

func (t *Service) GetListWithTravels(path string) ([]Travel, error) {
	records, err := t.repo.GetInfoCsv(path)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		trip := Travel{
			Id:         record[0],
			Name:       record[1],
			Email:      record[2],
			Country:    record[3],
			TravelTime: record[4],
			Price:      record[5],
		}

		t.travels = append(t.travels, trip)
	}

	return t.travels, nil
}

func (t *Service) GetTotalTickets(destination string) (int, error) {
	var totalTravelOfDest int
	for _, travel := range t.travels {
		if travel.Country == destination {
			totalTravelOfDest++
		}
	}

	if totalTravelOfDest == 0 {
		return 0, fmt.Errorf("Não foi encontrado viagens realizadas de acordo com o destino: %s", destination)
	}

	return totalTravelOfDest, nil
}

func (t *Service) GetCountByPeriod() (map[string]int, error) {
	for _, travel := range t.travels {
		timeParts := strings.Split(travel.TravelTime, ":")
		time, err := strconv.Atoi(timeParts[0])
		if err != nil {
			fmt.Printf("Erro ao converter string para int. %s", travel.TravelTime)
			continue
		}
		countPeriod(time)
	}

	statistic := make(map[string]int)
	statistic["IM"] = earlyMorningTrips
	statistic["M"] = morningTrips
	statistic["T"] = afternoonTrips
	statistic["N"] = nightTrips

	return statistic, nil
}

func countPeriod(time int) {
	switch {
	case time >= 0 && time <= 6:
		earlyMorningTrips++
	case time >= 7 && time <= 12:
		morningTrips++
	case time >= 13 && time <= 19:
		afternoonTrips++
	case time >= 20 && time <= 23:
		nightTrips++
	default:
		break
	}
}

func (t *Service) AverageDestination(destination string, time int) (float64, error) {
	var traveledToDest float64
	var arrivedInSameTime float64
	for _, travel := range t.travels {
		timeParts := strings.Split(travel.TravelTime, ":")
		timeOfTravel, err := strconv.Atoi(timeParts[0])
		if err != nil {
			fmt.Printf("Erro ao converter string para int. %s", travel.TravelTime)
			continue
		}

		if timeOfTravel == time {
			arrivedInSameTime++

			if travel.Country == destination {
				traveledToDest++
			}
		}
	}

	porcentagem := (traveledToDest / arrivedInSameTime) * 100
	return porcentagem, nil
}
