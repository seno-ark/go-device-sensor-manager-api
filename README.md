# go-device-sensor-manager-api

Backend Test Submission for Mertani.

This API facilitates CRUD operations for managing devices and sensors. 
It is created with Golang and PostgreSQL. The routing is handled using Chi, and database operations are managed through Sqlx.

## How to

Clone this repository
```
git clone git@github.com:seno-ark/go-device-sensor-manager-api.git
```

Copy .env-example to .env
```
cd go-device-sensor-manager-api
cp .env-example .env
```

Run local postgresql DB
```
make postgres
```

Migrate database schema
```
make migrate-up
```

Start the application in dev mode
```
make dev
```

If you prefer running the build version:
```
make build
./bin/api
```

## Endpoints

All endpoints are available in swagger and postman collection:
- Swagger: http://localhost:9000/swagger/index.html
- Postman: https://github.com/seno-ark/go-device-sensor-manager-api/blob/main/mertani.postman_collection.json

#### Create Device
```
POST /v1/devices
json body:
{
  "name": "Device1",
  "description": "Device 1",
  "status": "active"
}
```

#### Update Device
```
PUT /v1/devices/:device_id
json body:
{
  "name": "Device#1",
  "description": "Device 2",
  "status": "active"
}
```

#### Delete Device
```
DELETE /v1/devices/:device_id
```

#### Get Device
```
GET /v1/devices/:device_id
```

#### Get Device List
```
GET /v1/devices
query params:
- page (int) 
- count (int) 
- sort (string) : name, -name, created_at, -created_at, updated_at, -updated_at
- search (string) 
```

#### Create Sensor
```
POST /v1/sensors
json body:
{
  "device_id": "d2431891-c5e4-462d-bf9b-7a194d5bebda",    
  "description": "sensor1.1",
  "name": "sensor #1.1",
  "type": "air"
}
```

#### Update Sensor
```
PUT /v1/sensors/:sensor_id
json body:
{
  "description": "sensor1.2",
  "name": "sensor #1.2"
}
```

#### Delete Sensor
```
DELETE /v1/sensors/:sensor_id
```

#### Get Sensor
```
GET /v1/sensors/:sensor_id
```

#### Get Sensor List
```
GET /v1/sensors
query params:
- page (int) 
- count (int) 
- sort (string) : name, -name, created_at, -created_at, updated_at, -updated_at
- device_id (string) 
- search (string) 
```

#### Get Sensor Type List
```
GET /v1/sensors/types
```

## Commands

### make dev
To run the application in dev mode

### make build
To build the executable binary

### make swagger
To build swagger documentation

need install https://github.com/swaggo/swag

### make postgres
To start local postgresql database using docker

### make migrate-file
To generate database migration file

need install https://github.com/golang-migrate/migrate

### make migrate-up
To execute database migration

need install https://github.com/golang-migrate/migrate

### make migrate-down
To undo database migration

need install https://github.com/golang-migrate/migrate
