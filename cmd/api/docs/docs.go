// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/devices": {
            "get": {
                "description": "Get list of Device.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Get list of Device.",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Pagination page number (default 1, max 500)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "Pagination data limit  (default 10, max 100)",
                        "name": "count",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "-created_at",
                        "description": "Data sorting (value: name/created_at/updated_at). For desc order, use prefix '-'",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "raspi",
                        "description": "Keyword for searching device by title or content",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Device"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new Device.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Create Device.",
                "parameters": [
                    {
                        "description": "Device data",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.CreateUpdateDevicePayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Device"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/v1/devices/{device_id}": {
            "get": {
                "description": "Get device by device ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Get device by device ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "device_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Device"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing Device.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Update Device.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "01HQSH92SNYQVCBDSD38XNBRYM",
                        "description": "Device ID",
                        "name": "device_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Device data",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.CreateUpdateDevicePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Device"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Device.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Delete Device.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "01HQSH92SNYQVCBDSD38XNBRYM",
                        "description": "Device ID",
                        "name": "device_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/v1/sensors": {
            "get": {
                "description": "Get list of Sensor.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensors"
                ],
                "summary": "Get list of Sensor.",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Pagination page number (default 1, max 500)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "Pagination data limit  (default 10, max 100)",
                        "name": "count",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "-created_at",
                        "description": "Data sorting (value: name/created_at/updated_at). For desc order, use prefix '-'",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "96a5ec77-9012-4bf3-b08e-39ef4c07fcce",
                        "description": "Filter sensors by device ID",
                        "name": "device_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "soil",
                        "description": "Keyword for searching sensors by name or description",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Sensor"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new Sensor.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensors"
                ],
                "summary": "Create Sensor.",
                "parameters": [
                    {
                        "description": "Sensor data",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.CreateSensorPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Sensor"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/v1/sensors/types": {
            "get": {
                "description": "Get Sensor Types.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensors"
                ],
                "summary": "Get Sensor Types.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/v1/sensors/{sensor_id}": {
            "get": {
                "description": "Get sensor by sensor ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensors"
                ],
                "summary": "Get sensor by sensor ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sensor ID",
                        "name": "sensor_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Sensor"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing Sensor.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensors"
                ],
                "summary": "Update Sensor.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "96a5ec77-9012-4bf3-b08e-39ef4c07fcce",
                        "description": "Sensor ID",
                        "name": "sensor_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Sensor data",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UpdateSensorPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Sensor"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Sensor.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensors"
                ],
                "summary": "Delete Sensor.",
                "parameters": [
                    {
                        "type": "string",
                        "example": "96a5ec77-9012-4bf3-b08e-39ef4c07fcce",
                        "description": "Sensor ID",
                        "name": "sensor_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.CreateSensorPayload": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "First Sensor"
                },
                "device_id": {
                    "type": "string",
                    "example": "d2431891-c5e4-462d-bf9b-7a194d5bebda"
                },
                "name": {
                    "type": "string",
                    "example": "Sensor #1"
                },
                "type": {
                    "type": "string",
                    "example": "temperature"
                }
            }
        },
        "entities.CreateUpdateDevicePayload": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "First device"
                },
                "name": {
                    "type": "string",
                    "example": "Device #1"
                },
                "status": {
                    "type": "string",
                    "example": "active"
                }
            }
        },
        "entities.Device": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entities.Sensor": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "device_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entities.UpdateSensorPayload": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "First Sensor v2"
                },
                "name": {
                    "type": "string",
                    "example": "Sensor #1.2"
                }
            }
        },
        "util.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/util.ResponseMeta"
                },
                "validation_errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "util.ResponseMeta": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Device-Sensor API",
	Description:      "This is a simple api server to manage devices and sensors",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
