package repository

import (
	"context"
	"testing/go/grcp/models"
)

type Rspository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var implementation Repository

func SetRepository(respository Repository) {
	implementation = respository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implemetnacion.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implemetnacion.GetStudent(ctx, id)
}
