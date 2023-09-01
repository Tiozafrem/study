package services

import (
	"context"

	"github.com/tiozafrem/study/second/model"
	"github.com/tiozafrem/study/second/repository"
)

type Person interface {
	GetAll(ctx context.Context) ([]model.Person, error)
	GetChampionsByAge(ctx context.Context, age model.Age) ([]model.Person, error)
	Create(ctx context.Context, person model.Person) error
	Update(ctx context.Context, person model.Person) error
	Delete(ctx context.Context, person model.Person) error
}

type Age interface {
	GetAll(ctx context.Context) ([]model.Age, error)
	Update(ctx context.Context, age model.Age) error
}

type Services struct {
	Person
	Age
}

func NewServices(repository *repository.Repository) *Services {
	return &Services{
		Person: NewPersonServices(repository.Person),
		Age:    NewAgeServices(repository.Age),
	}
}
