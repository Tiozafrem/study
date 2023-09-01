package repository

import (
	"context"
	"database/sql"

	"github.com/tiozafrem/study/second/model"
	"github.com/tiozafrem/study/second/repository/postgres"
)

type Person interface {
	GetAll(ctx context.Context) ([]model.Person, error)
	GetChampionsByAge(ctx context.Context, age model.Age) ([]model.Person, error)
	GetCountByAge(ctx context.Context, age model.Age) (int, error)
	Create(ctx context.Context, person model.Person) (int, error)
	Update(ctx context.Context, person model.Person) error
	Delete(ctx context.Context, person model.Person) error
}

type Age interface {
	GetAll(ctx context.Context) ([]model.Age, error)
	Update(ctx context.Context, age model.Age) error
}

type Repository struct {
	Person
	Age
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Person: postgres.NewPersonPostgres(db),
		Age:    postgres.NewAgePostgres(db),
	}
}
