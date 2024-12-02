package SQLite

import (
	"context"
	"database/sql"
	"helios/internal/api/repository/DAL"
	"helios/internal/api/repository/models"
)

type DataRepository struct {
	sqlDB *sql.DB
	createStmt,
	readStmt,
	readManyStmt,
	updateStmt,
	deleteStmt *sql.Stmt
	ctx context.Context
}

func InitializeSensorRepository(sqlDB DAL.SQLDatabase, ctx context.Context) (models.DataRepository, error) {

	repo := &DataRepository{
		sqlDB: sqlDB.Connection(),
		ctx:   ctx,
	}

	// Create the data table if it doesn't exist
	if _, err := repo.sqlDB.Exec(`CREATE TABLE  IF NOT EXISTS sensor (
		id INTEGER PRIMARY KEY,
		type INTEGER,
		value FLOAT,
		timestamp TIMESTAMP
	);`); err != nil {
		repo.sqlDB.Close()
		return nil, err
	}

	// * Create needed Prepared SQL statements, this is more efficient than running each query individually
	createStmt, err := repo.sqlDB.Prepare(`INSERT INTO sensor (id, type, value, timestamp) VALUES (?, ?, ?, ?)`)
	if err != nil {
		repo.sqlDB.Close() // Close the database connection if statement preparation fails
		return nil, err
	}
	repo.createStmt = createStmt

	readStmt, err := repo.sqlDB.Prepare("SELECT id, type, value, timestamp FROM sensor WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readStmt = readStmt

	readManyStmt, err := repo.sqlDB.Prepare("SELECT id, type, value, timestamp FROM sensor LIMIT ? OFFSET ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readManyStmt = readManyStmt

	updateStmt, err := repo.sqlDB.Prepare("UPDATE sensor SET id = ?, type = ?, value = ?, timestamp = ? WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateStmt = updateStmt

	deleteStmt, err := repo.sqlDB.Prepare("DELETE FROM sensor WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.deleteStmt = deleteStmt

	go Close(ctx, repo)

	return repo, nil
}

func Close(ctx context.Context, r *DataRepository) {

	<-ctx.Done()
	r.createStmt.Close()
	r.readStmt.Close()
	r.updateStmt.Close()
	r.deleteStmt.Close()
	r.readManyStmt.Close()
	r.sqlDB.Close()
}

func (r *DataRepository) Create(data *models.SensorData, ctx context.Context) error {

	res, err := r.createStmt.ExecContext(ctx, data.ID, data.Type, data.Value, data.Timestamp)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	data.ID = int(id)
	return nil
}

func (r *DataRepository) ReadOne(id int, ctx context.Context) (*models.SensorData, error) {
	row := r.readStmt.QueryRowContext(ctx, id)
	var data models.SensorData
	err := row.Scan(&data.ID, &data.Type, &data.Value, &data.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (r *DataRepository) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.SensorData, error) {

	if page < 1 {
		return r.ReadAll()
	}

	offset := rowsPerPage * (page - 1)
	rows, err := r.readManyStmt.QueryContext(ctx, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.SensorData
	for rows.Next() {
		var d models.SensorData
		err := rows.Scan(&d.ID, &d.Type, &d.Value, &d.Timestamp)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}
	return data, nil
}

func (r *DataRepository) ReadAll() ([]*models.SensorData, error) {
	rows, err := r.sqlDB.QueryContext(context.Background(), "SELECT id, type, value, timestamp FROM sensor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.SensorData
	for rows.Next() {
		var d models.SensorData
		err := rows.Scan(&d.ID, &d.Type, &d.Value, &d.Timestamp)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}
	return data, nil
}

func (r *DataRepository) Update(data *models.SensorData, ctx context.Context) (int64, error) {
	res, err := r.updateStmt.ExecContext(ctx, data.ID, data.Type, data.Value, data.Timestamp, data.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *DataRepository) Delete(data *models.SensorData, ctx context.Context) (int64, error) {
	res, err := r.deleteStmt.ExecContext(ctx, data.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
