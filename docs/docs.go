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
        "/login": {
            "get": {
                "tags": [
                    "Example"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorException.BadRequestResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errorException.NotFoundResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorException.InternalServerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errorException.BadRequestResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "token_invalid"
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "errorException.InternalServerResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "internal_server_error"
                },
                "status_code": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "errorException.NotFoundResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "user_not_found"
                },
                "status_code": {
                    "type": "integer",
                    "example": 404
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Example: \"Bearer {token}\".",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
