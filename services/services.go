package services

import "go-boilerplate/repositories"

type IServices interface {
}

type service struct {
	r repositories.IRepository
}

func NewService(r repositories.IRepository) IServices {
	return &service{r: r}
}

// func (s *service) {

// }
