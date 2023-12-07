package repositories

import (
	"lib/internal/models"

	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (brm *BookRepositoryMock) ReadAllBooks(page int, limit int) ([]models.BookResponseModel, error) {
	args := brm.Mock.Called(page, limit)
	return args.Get(0).([]models.BookResponseModel), args.Error(1)
}

func (brm *BookRepositoryMock) CreateNewBook(body *models.BookModel) error {
	args := brm.Mock.Called(body)
	return args.Error(0)
}
