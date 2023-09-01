package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/tiozafrem/study/second/model"
)

type AgePostgres struct {
	db *sql.DB
}

func NewAgePostgres(db *sql.DB) *AgePostgres {
	return &AgePostgres{db: db}
}

func (r *AgePostgres) GetAll(ctx context.Context) ([]model.Age, error) {
	var ages []model.Age

	rows, err := r.db.QueryContext(ctx,
		fmt.Sprintf(
			`SELECT id, name, age_start 
			FROM %s
			ORDER BY age_start ASC`, ageTable))
	if err != nil {
		return ages, err
	}

	for rows.Next() {
		var age model.Age
		err := rows.Scan(&age.Id, &age.Name, &age.AgeStart)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		} else {
			ages = append(ages, age)
		}
	}

	return ages, err
}

func (r *AgePostgres) Update(ctx context.Context, age model.Age) error {

	result, err := r.db.ExecContext(ctx,
		fmt.Sprintf(`
		UPDATE %s
		SET age_start = $2
		WHERE id = $1
		`, ageTable), age.Id, age.AgeStart)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return fmt.Errorf("not affected rows")
	}

	return nil
}
