// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "",
        "contact": {
            "name": "Weather API",
            "url": "",
            "email": "anton.feldmann@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/now/{lat}/{lon}": {
            "get": {
                "description": "get string by ID",
                "summary": "Show an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "51.873960",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "8.156710",
                        "name": "lon",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.WeatherResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.WeatherResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.WeatherResponse"
                        }
                    }
                }
            }
        },
        "/{date}/{lat}/{lon}": {
            "get": {
                "description": "get a weather description of dwd",
                "summary": "get weather information by date",
                "parameters": [
                    {
                        "type": "string",
                        "description": "2020-10-25",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "51.873960",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "8.156710",
                        "name": "lon",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.WeatherResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.WeatherResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.WeatherResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.WeatherResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "cannot load data"
                },
                "status": {
                    "type": "integer",
                    "example": 404
                },
                "weather": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/wetter.WeatherPoint"
                    }
                }
            }
        },
        "wetter.FallBackSourceId": {
            "type": "object",
            "properties": {
                "cloud_cover": {
                    "type": "integer",
                    "example": 2
                },
                "icon": {
                    "type": "string",
                    "example": "cloud"
                },
                "pressure_msl": {
                    "type": "integer",
                    "example": 4
                },
                "visibility": {
                    "type": "integer",
                    "example": 0
                },
                "wind_direction": {
                    "type": "integer",
                    "example": 1
                },
                "wind_gust_speed": {
                    "type": "integer",
                    "example": 5
                },
                "wind_speed": {
                    "type": "integer",
                    "example": 3
                }
            }
        },
        "wetter.WeatherPoint": {
            "type": "object",
            "properties": {
                "cloud_cover": {
                    "type": "number",
                    "example": 90
                },
                "condition": {
                    "type": "string",
                    "example": "dry"
                },
                "dew_point": {
                    "type": "number",
                    "example": 10.6
                },
                "fallback_source_ids": {
                    "$ref": "#/definitions/wetter.FallBackSourceId"
                },
                "precipitation": {
                    "type": "number",
                    "example": 1007.6
                },
                "pressure_msl": {
                    "type": "number",
                    "example": 14.1
                },
                "relative_humidity": {
                    "type": "number",
                    "example": 31.2
                },
                "source_id": {
                    "type": "number",
                    "example": 6946
                },
                "sunshine": {
                    "type": "number",
                    "example": 200
                },
                "temperature": {
                    "type": "number",
                    "example": 19.4
                },
                "timestamp": {
                    "type": "string",
                    "example": "2020-10-25T00:00:00+00:00"
                },
                "visibility": {
                    "type": "number",
                    "example": 41.2
                },
                "wind_direction": {
                    "type": "number",
                    "example": 100
                },
                "wind_gust_direction": {
                    "type": "number",
                    "example": 40.2
                },
                "wind_gust_speed": {
                    "type": "number",
                    "example": 12.2
                },
                "wind_speed": {
                    "type": "number",
                    "example": 13.5
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "https://weather",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Weather information download",
	Description:      "request for weather informations",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
