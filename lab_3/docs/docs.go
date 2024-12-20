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
        "/records": {
            "get": {
                "description": "Возвращает массив элементов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Получить все записи",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/data.Record"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет новый элемент",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Добавить новую запись",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/data.Record"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "data.Record": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "My Project API",
	Description:      "Это API для работы с записями.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
