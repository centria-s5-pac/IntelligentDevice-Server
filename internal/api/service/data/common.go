package data

import (
	"context"
	"helios/internal/api/repository/models"
)

type DataService interface {
	Create(data *models.SensorData, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*models.SensorData, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.SensorData, error)
	Update(data *models.SensorData, ctx context.Context) (int64, error)
	Delete(data *models.SensorData, ctx context.Context) (int64, error)
	ValidateData(data *models.SensorData) error
}

type DataError struct {
	Message string
}

func (de DataError) Error() string {
	return de.Message
}
