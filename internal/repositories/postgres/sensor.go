package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-api/internal/entities"
	"go-api/pkg/util"
	"log/slog"
	"strings"
	"time"
)

type Sensor struct {
	ID          string    `db:"id"`
	DeviceID    string    `db:"device_id"`
	Type        string    `db:"type"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (s *Sensor) ToEntity() *entities.Sensor {
	return &entities.Sensor{
		ID:          s.ID,
		DeviceID:    s.DeviceID,
		Type:        entities.SensorType(s.Type),
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func (r *repository) CreateSensor(ctx context.Context, payload entities.Sensor) (string, error) {
	var sensorID string

	nowUTC := time.Now().UTC()
	payload.CreatedAt = nowUTC
	payload.UpdatedAt = nowUTC

	query := `INSERT INTO sensors 
		(device_id, type, name, description, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		payload.DeviceID,
		payload.Type,
		payload.Name,
		payload.Description,
		payload.CreatedAt,
		payload.UpdatedAt,
	).Scan(&sensorID)
	if err != nil {
		slog.Error(
			"Failed to CreateSensor",
			slog.Any("err", err),
			slog.Any("payload", payload),
		)
		return sensorID, util.NewErrInternalServer("failed to create sensor")
	}

	return sensorID, nil
}

func (r *repository) UpdateSensor(ctx context.Context, sensorID string, payload entities.Sensor) error {
	query := `UPDATE sensors 
		SET name = $1, description = $2, updated_at = $3 
		WHERE id = $4`

	_, err := r.db.ExecContext(
		ctx,
		query,
		payload.Name,
		payload.Description,
		time.Now().UTC(),
		sensorID,
	)
	if err != nil {
		slog.Error(
			"Failed to UpdateSensor",
			slog.Any("err", err),
			slog.Any("sensorID", sensorID),
			slog.Any("payload", payload),
		)
		return util.NewErrInternalServer("failed to update sensor")
	}

	return nil
}

func (r *repository) DeleteSensor(ctx context.Context, sensorID string) error {
	query := `DELETE FROM sensors WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, sensorID)
	if err != nil {
		slog.Error(
			"Failed to DeleteSensor",
			slog.Any("err", err),
			slog.Any("sensorID", sensorID),
		)
		return util.NewErrInternalServer("failed to delete sensor")
	}

	return err
}

func (r *repository) GetSensor(ctx context.Context, sensorID string) (*entities.Sensor, error) {
	var model Sensor

	query := `SELECT id, device_id, type, name, description, created_at, updated_at FROM sensors WHERE id = $1`
	err := r.db.GetContext(ctx, &model, query, sensorID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.NewErrNotFound("sensor not found")
		}

		slog.Error(
			"Failed to GetSensor",
			slog.Any("err", err),
			slog.Any("sensorID", sensorID),
		)
		return nil, util.NewErrInternalServer("failed to get sensor")
	}

	return model.ToEntity(), nil
}

func (r *repository) GetSensorList(ctx context.Context, params entities.GetSensorListParams) ([]*entities.Sensor, int64, error) {
	var (
		total          int64
		availableSorts = []string{"name", "created_at", "updated_at"}
		orderBy        = util.SortValidation(params.Sort, availableSorts)
	)

	queryCount := "SELECT COUNT(id) FROM sensors"
	queryData := "SELECT id, device_id, type, name, description, created_at, updated_at FROM sensors"

	whereQueries := []string{}
	if params.Search != "" {
		params.Search = fmt.Sprintf("%%%s%%", params.Search)
		whereQueries = append(whereQueries, "(name LIKE :keyword OR description LIKE :keyword)")
	}
	if params.DeviceID != "" {
		whereQueries = append(whereQueries, "device_id = :device_id")
	}
	if len(whereQueries) > 0 {
		whereQuery := fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND "))

		queryCount += whereQuery
		queryData += whereQuery
	}

	queryData += fmt.Sprintf(" ORDER BY %s LIMIT %d OFFSET %d", orderBy, params.Limit, params.Offset)

	// COUNT ROWS

	stmtCount, err := r.db.PrepareNamed(queryCount)
	if err != nil {
		slog.Error(
			"Failed to GetSensorList Count PrepareNamed",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get sensor list")
	}
	defer stmtCount.Close()

	err = stmtCount.GetContext(ctx, &total, map[string]any{
		"device_id": params.DeviceID,
		"keyword":   params.Search,
	})
	if err != nil {
		slog.Error(
			"Failed to GetSensorList Count GetContext",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get sensor list")
	}

	if total == 0 {
		return nil, total, nil
	}

	// SELECT ROWS

	stmtData, err := r.db.PrepareNamed(queryData)
	if err != nil {
		slog.Error(
			"Failed to GetSensorList Data PrepareNamed",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get sensor list")
	}
	defer stmtData.Close()

	var model []Sensor
	err = stmtData.SelectContext(ctx, &model, map[string]any{
		"device_id": params.DeviceID,
		"keyword":   params.Search,
	})
	if err != nil {
		slog.Error(
			"Failed to GetSensorList Data SelectContext",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get sensor list")
	}

	sensors := []*entities.Sensor{}
	for _, v := range model {
		sensors = append(sensors, v.ToEntity())
	}

	return sensors, total, nil
}
