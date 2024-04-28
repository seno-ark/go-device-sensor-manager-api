package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-api/internal/entities"
	"go-api/pkg/util"
	"log/slog"
	"time"
)

type Device struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (d *Device) ToEntity() *entities.Device {
	return &entities.Device{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Status:      entities.DeviceStatus(d.Status),
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func (r *repository) CreateDevice(ctx context.Context, payload entities.Device) (string, error) {
	var deviceID string

	nowUTC := time.Now().UTC()
	payload.CreatedAt = nowUTC
	payload.UpdatedAt = nowUTC

	query := `INSERT INTO devices 
	(name, description, status, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		payload.Name,
		payload.Description,
		payload.Status,
		payload.CreatedAt,
		payload.UpdatedAt,
	).Scan(&deviceID)

	if err != nil {
		slog.Error(
			"Failed to CreateDevice",
			slog.Any("err", err),
			slog.Any("payload", payload),
		)
		return deviceID, util.NewErrInternalServer("failed to create device")
	}

	return deviceID, nil
}

func (r *repository) UpdateDevice(ctx context.Context, deviceID string, payload entities.Device) error {
	query := `UPDATE devices 
	SET name = $1, description = $2, status = $3, updated_at = $4 
	WHERE id = $5`

	_, err := r.db.ExecContext(
		ctx,
		query,
		payload.Name,
		payload.Description,
		payload.Status,
		time.Now().UTC(),
		deviceID,
	)
	if err != nil {
		slog.Error(
			"Failed to UpdateDevice",
			slog.Any("err", err),
			slog.Any("deviceID", deviceID),
			slog.Any("payload", payload),
		)
		return util.NewErrInternalServer("failed to update device")
	}

	return nil
}

func (r *repository) DeleteDevice(ctx context.Context, deviceID string) error {
	query := `DELETE FROM devices WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, deviceID)
	if err != nil {
		slog.Error(
			"Failed to DeleteDevice",
			slog.Any("err", err),
			slog.Any("deviceID", deviceID),
		)
		return util.NewErrInternalServer("failed to delete device")
	}

	return err
}

func (r *repository) GetDevice(ctx context.Context, deviceID string) (*entities.Device, error) {
	var model Device

	query := `SELECT id, name, description, status, created_at, updated_at FROM devices WHERE id = $1`
	err := r.db.GetContext(ctx, &model, query, deviceID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.NewErrNotFound("device not found")
		}

		slog.Error(
			"Failed to GetDevice",
			slog.Any("err", err),
			slog.Any("deviceID", deviceID),
		)
		return nil, util.NewErrInternalServer("failed to get device")
	}

	return model.ToEntity(), nil
}

func (r *repository) GetDeviceList(ctx context.Context, params entities.GetDeviceListParams) ([]*entities.Device, int64, error) {
	var (
		total          int64
		availableSorts = []string{"name", "created_at", "updated_at"}
		orderBy        = util.SortValidation(params.Sort, availableSorts)
	)

	queryCount := "SELECT COUNT(id) FROM devices"
	queryData := "SELECT id, name, description, status, created_at, updated_at FROM devices"

	if params.Search != "" {
		params.Search = fmt.Sprintf("%%%s%%", params.Search)
		whereQuery := " WHERE (name LIKE :keyword OR description LIKE :keyword)"
		queryCount += whereQuery
		queryData += whereQuery
	}

	queryData += fmt.Sprintf(" ORDER BY %s LIMIT %d OFFSET %d", orderBy, params.Limit, params.Offset)

	// COUNT ROWS

	stmtCount, err := r.db.PrepareNamed(queryCount)
	if err != nil {
		slog.Error(
			"Failed to GetDeviceList Count PrepareNamed",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get device list")
	}
	defer stmtCount.Close()

	err = stmtCount.GetContext(ctx, &total, map[string]any{
		"keyword": params.Search,
	})
	if err != nil {
		slog.Error(
			"Failed to GetDeviceList Count GetContext",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get device list")
	}

	if total == 0 {
		return nil, total, nil
	}

	// SELECT ROWS

	stmtData, err := r.db.PrepareNamed(queryData)
	if err != nil {
		slog.Error(
			"Failed to GetDeviceList Data PrepareNamed",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get device list")
	}
	defer stmtData.Close()

	var model []Device
	err = stmtData.SelectContext(ctx, &model, map[string]any{
		"keyword": params.Search,
	})
	if err != nil {
		slog.Error(
			"Failed to GetDeviceList Data SelectContext",
			slog.Any("err", err),
			slog.Any("params", params),
		)
		return nil, total, util.NewErrInternalServer("failed to get device list")
	}

	devices := []*entities.Device{}
	for _, v := range model {
		devices = append(devices, v.ToEntity())
	}

	return devices, total, nil
}
