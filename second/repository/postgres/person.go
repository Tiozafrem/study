package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/tiozafrem/study/second/model"
)

type PersonPostgres struct {
	db *sql.DB
}

func NewPersonPostgres(db *sql.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}

func (r *PersonPostgres) GetAll(ctx context.Context) ([]model.Person, error) {
	var peoples []model.Person

	rows, err := r.db.QueryContext(ctx,
		fmt.Sprintf(
			`
			SELECT peoples.id, peoples.surname, peoples.age,
			peoples.time_start, peoples.time_end, 
			(CASE 
				WHEN (peoples.time_end < peoples.time_start) 
				THEN
					(extract(epoch from (peoples.time_end - peoples.time_start + '1 day'::interval))* 1000000000)::bigint
				ELSE
					(extract(epoch from (peoples.time_end - peoples.time_start))* 1000000000)::bigint
				END) as interval
			FROM %s peoples
			ORDER BY age ASC`,
			personTable))
	if err != nil {
		return peoples, err
	}

	for rows.Next() {
		var person model.Person
		err := rows.Scan(&person.Id, &person.Surname, &person.Age,
			&person.TimeStart, &person.TimeEnd, &person.Interval)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		} else {
			peoples = append(peoples, person)
		}
	}

	return peoples, err
}

func (r *PersonPostgres) GetChampionsByAge(ctx context.Context, age model.Age) ([]model.Person, error) {
	var peoples []model.Person

	rows, err := r.db.QueryContext(ctx,
		fmt.Sprintf(
			`WITH 
				age_end AS (
					SELECT COALESCE(a2.age_start, 122) as age_end 
					FROM %s a1
						LEFT JOIN  %s a2
							ON a2.age_start > a1.age_start
					WHERE a1.id = %d
					ORDER BY age_end ASC 
					LIMIT 1),
				people_tournament AS (
					SELECT peoples.id, peoples.surname, peoples.age,
					peoples.time_start, peoples.time_end, 
					interval, RANK() OVER (ORDER BY interval) as rank
					FROM (
						SELECT peoples.id, peoples.surname, peoples.age,
						peoples.time_start, peoples.time_end, 
						(CASE 
							WHEN (peoples.time_end < peoples.time_start) 
							THEN
								(extract(epoch from (peoples.time_end - peoples.time_start + '1 day'::interval))* 1000000000)::bigint
							ELSE
								(extract(epoch from (peoples.time_end - peoples.time_start))* 1000000000)::bigint
							END) as interval
						FROM %s peoples
					) as peoples
					INNER JOIN age_end
						ON peoples.age BETWEEN %d AND age_end.age_end
					ORDER BY interval ASC)
			SELECT people_tournament.id, people_tournament.surname, people_tournament.age,
				people_tournament.time_start, people_tournament.time_end, 
				people_tournament.interval
			FROM people_tournament
			WHERE rank = 1`,
			ageTable, ageTable, age.Id,
			personTable, age.AgeStart), ageTable)
	if err != nil {
		return peoples, err
	}

	for rows.Next() {
		var person model.Person
		err := rows.Scan(&person.Id, &person.Surname, &person.Age,
			&person.TimeStart, &person.TimeEnd, &person.Interval)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		} else {
			peoples = append(peoples, person)
		}
	}

	return peoples, err
}

func (r *PersonPostgres) GetCountByAge(ctx context.Context, age model.Age) (int, error) {

	rows := r.db.QueryRowContext(ctx,
		fmt.Sprintf(
			`WITH age_end AS (
				SELECT age_start as age_end 
				FROM %s 
				WHERE age_start > %d 
				ORDER BY age_start ASC 
				LIMIT 1)
			SELECT COUNT(*)
			FROM %s peoples
			WHERE
				peoples.age < COALESCE(
					(SELECT age_end.age_end 
						FROM age_end), 122)
				AND peoples.age > %d`,
			ageTable, age.AgeStart,
			personTable, age.AgeStart))

	var count int
	err := rows.Scan(&count)

	return count, err
}

func (r *PersonPostgres) Create(ctx context.Context, person model.Person) (int, error) {
	var id int
	row := r.db.QueryRowContext(ctx,
		fmt.Sprintf(
			`INSERT INTO %s (surname, age, time_start, time_end)
			VALUES ($1, $2, $3, $4) RETURNING id`, personTable,
		), person.Surname, person.Age, person.TimeStart, person.TimeEnd)

	err := row.Scan(&id)
	return id, err
}

func (r *PersonPostgres) Update(ctx context.Context, person model.Person) error {

	result, err := r.db.ExecContext(ctx,
		fmt.Sprintf(`
		UPDATE %s
		SET surname = $2, age = $3,
			time_start = $4, time_end = $5
		WHERE id = $1
		`, personTable), person.Id, person.Surname, person.Age,
		person.TimeStart, person.TimeEnd)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return fmt.Errorf("not affected rows")
	}

	return nil
}

func (r *PersonPostgres) Delete(ctx context.Context, person model.Person) error {

	result, err := r.db.ExecContext(ctx,
		fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
		`, personTable), person.Id)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return fmt.Errorf("not affected rows")
	}

	return nil
}
