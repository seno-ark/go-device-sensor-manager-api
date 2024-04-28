package entities

import "time"

type DeviceStatus string

var (
	DEVICE_STATUS_ACTIVE   DeviceStatus = "active"
	DEVICE_STATUS_INACTIVE DeviceStatus = "inactive"
)

type Device struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      DeviceStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type CreateUpdateDevicePayload struct {
	Name        string       `json:"name" validate:"required" example:"Device #1"`
	Description string       `json:"description" example:"First device"`
	Status      DeviceStatus `json:"status" validate:"deviceStatus" example:"active"`
}

type GetDeviceListParams struct {
	Search string
	Sort   string
	Limit  int
	Offset int
}
