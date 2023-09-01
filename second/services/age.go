package services

import (
	"context"
	"fmt"
	"os"

	"github.com/tiozafrem/study/second/model"
	"github.com/tiozafrem/study/second/repository"
)

type AgeServices struct {
	repository repository.Age
}

func NewAgeServices(repository repository.Age) *AgeServices {
	return &AgeServices{repository: repository}
}

func (s *AgeServices) GetAll(ctx context.Context) ([]model.Age, error) {
	ages, err := s.repository.GetAll(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}
	// for _, age := range ages {
	// 	fmt.Printf("%d %s %d\n", age.Id, age.Name, age.AgeStart)
	// }
	return ages, err
}

func (s *AgeServices) Update(ctx context.Context, age model.Age) error {
	err := s.repository.Update(ctx, age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}

	//fmt.Printf("%d %s %d\n", age.Id, age.Name, age.AgeStart)

	return err
}
