package entities

import "time"

type SensorType string

var (
	SENSOR_TYPE_TEMPERATURE SensorType = "temperature"
	SENSOR_TYPE_AIR         SensorType = "air"
	SENSOR_TYPE_WATER       SensorType = "water"

	SensorTypesName = []map[string]string{
		{
			"slug": string(SENSOR_TYPE_TEMPERATURE),
			"name": "Temperature",
		},
		{
			"slug": string(SENSOR_TYPE_AIR),
			"name": "Air",
		},
		{
			"slug": string(SENSOR_TYPE_WATER),
			"name": "Water",
		},
	}
)

type Sensor struct {
	ID          string     `json:"id"`
	DeviceID    string     `json:"device_id"`
	Type        SensorType `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateSensorPayload struct {
	DeviceID    string     `json:"device_id" validate:"uuid" example:"d2431891-c5e4-462d-bf9b-7a194d5bebda"`
	Type        SensorType `json:"type" validate:"sensorType" example:"temperature"`
	Name        string     `json:"name" validate:"required" example:"Sensor #1"`
	Description string     `json:"description" example:"First Sensor"`
}

type UpdateSensorPayload struct {
	Name        string `json:"name" validate:"required" example:"Sensor #1.2"`
	Description string `json:"description" example:"First Sensor v2"`
}

type GetSensorListParams struct {
	DeviceID string
	Search   string
	Sort     string
	Limit    int
	Offset   int
}
