package service

import (
	"context"
	"helios/internal/api/repository/DAL"
	"helios/internal/api/repository/DAL/SQLite"
	service "helios/internal/api/service/data"
	"log"
)

type DataServiceType int

const (
	SQLiteDataService DataServiceType = iota
)

type ServiceFactory struct {
	db     DAL.SQLDatabase
	logger *log.Logger
	ctx    context.Context
}

// * Factory for creating data service *
func NewServiceFactory(db DAL.SQLDatabase, logger *log.Logger, ctx context.Context) *ServiceFactory {
	return &ServiceFactory{
		db:     db,
		logger: logger,
		ctx:    ctx,
	}
}

func (sf *ServiceFactory) CreateDataService(serviceType DataServiceType) (*service.DataServiceSQLite, error) {

	switch serviceType {

	case SQLiteDataService:
		repo, err := SQLite.InitializeSensorRepository(sf.db, sf.ctx)
		if err != nil {
			return nil, err
		}
		ds := service.NewDataServiceSQLite(repo)
		return ds, nil
	default:
		return nil, service.DataError{Message: "Invalid data service type."}
	}
}
