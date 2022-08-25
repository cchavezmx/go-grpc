package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"platzi.com/go/grpc/models"
)

type PostgresRespository struct {
	db *sql.DB
}

func NewPostgresRespository(url string) (*PostgresRespository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRespository{db}, nil
}

func (repo *PostgresRespository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name) VALUES ($1, $2)", student.Id, student.Name, student.Age)
	return err
}

func (repo *PostgresRespository) GetStudent(ctx, context.Context, id string) (*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	
	defer func ()  {
		err: = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var student models.Student{}
	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
	}

	return &student, nil
}