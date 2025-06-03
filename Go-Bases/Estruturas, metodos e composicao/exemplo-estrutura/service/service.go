package service

import "metodo/repository"

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{repository: r}
}

func (s Service) Save(st string) {
	s.repository.SaveValue(st)
}
