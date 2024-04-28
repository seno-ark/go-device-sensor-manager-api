package repositories

import (
	"context"
	"go-api/internal/entities"
)

type IRepository interface {
	CreateDevice(ctx context.Context, payload entities.Device) (string, error)
	UpdateDevice(ctx context.Context, deviceID string, payload entities.Device) error
	DeleteDevice(ctx context.Context, deviceID string) error
	GetDevice(ctx context.Context, deviceID string) (*entities.Device, error)
	GetDeviceList(ctx context.Context, params entities.GetDeviceListParams) ([]*entities.Device, int64, error)

	CreateSensor(ctx context.Context, payload entities.Sensor) (string, error)
	UpdateSensor(ctx context.Context, deviceID string, payload entities.Sensor) error
	DeleteSensor(ctx context.Context, deviceID string) error
	GetSensor(ctx context.Context, deviceID string) (*entities.Sensor, error)
	GetSensorList(ctx context.Context, params entities.GetSensorListParams) ([]*entities.Sensor, int64, error)
}
