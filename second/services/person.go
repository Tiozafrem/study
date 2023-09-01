package services

import (
	"context"
	"fmt"
	"os"

	"github.com/tiozafrem/study/second/model"
	"github.com/tiozafrem/study/second/repository"
)

type PersonServices struct {
	repository repository.Person
}

func NewPersonServices(repository repository.Person) *PersonServices {
	return &PersonServices{repository: repository}
}

func (s *PersonServices) GetAll(ctx context.Context) ([]model.Person, error) {
	peoples, err := s.repository.GetAll(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}
	// for _, person := range peoples {
	// 	fmt.Printf("%d %s %d\n", person.Id, person.Surname, person.Age)
	// }
	return peoples, err
}

func (s *PersonServices) Create(ctx context.Context, person model.Person) error {
	_, err := s.repository.Create(ctx, person)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}
	return err
}

func (s *PersonServices) GetChampionsByAge(ctx context.Context, age model.Age) ([]model.Person, error) {
	var peoples []model.Person

	count, err := s.repository.GetCountByAge(ctx, age)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return peoples, err
	}
	if count < 3 {
		err = fmt.Errorf("для подсчёта необходимо минимум 3 учасника")
		return peoples, err
	}

	peoples, err = s.repository.GetChampionsByAge(ctx, age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}

	// for _, person := range peoples {
	// 	fmt.Printf("%d %s %d %s %s %f\n", person.Id, person.Surname,
	// 		person.Age, person.TimeStart.Format("15:04:05"), person.TimeEnd.Format("15:04:05"), person.Interval.Minutes())
	// }
	return peoples, err
}

func (s *PersonServices) Update(ctx context.Context, person model.Person) error {
	err := s.repository.Update(ctx, person)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}

	//fmt.Printf("%d %s %d\n", age.Id, age.Name, age.AgeStart)

	return err
}

func (s *PersonServices) Delete(ctx context.Context, person model.Person) error {
	err := s.repository.Delete(ctx, person)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}

	//fmt.Printf("%d %s %d\n", age.Id, age.Name, age.AgeStart)

	return err
}
