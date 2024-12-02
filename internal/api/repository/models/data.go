package models

import "context"

type SensorData struct {
	ID        int     `json:"id"`
	Type      int     `json:"type"`
	Value     float64 `json:"value"`
	Timestamp string  `json:"timestamp"`
}

type DataRepository interface {
	Create(Data *SensorData, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*SensorData, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*SensorData, error)
	Update(data *SensorData, ctx context.Context) (int64, error)
	Delete(data *SensorData, ctx context.Context) (int64, error)
}
