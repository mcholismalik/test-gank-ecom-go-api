// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthLoginResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthRegisterResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "abstraction.PaginationInfo": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "more_records": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                },
                "sort_by": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginResponse": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone",
                "role"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderEntityModel"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.AuthLoginResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.AuthRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone",
                "role"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.AuthRegisterResponse": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone",
                "role"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderEntityModel"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.AuthRegisterResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.AuthRegisterResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "model.OrderEntityModel": {
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "order_items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderItemEntityModel"
                    }
                },
                "user_id": {
                    "description": "relations",
                    "type": "integer"
                }
            }
        },
        "model.OrderItemEntityModel": {
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "order_id": {
                    "description": "relations",
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                }
            }
        },
        "response.Meta": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/abstraction.PaginationInfo"
                },
                "message": {
                    "type": "string",
                    "default": "true"
                },
                "success": {
                    "type": "boolean",
                    "default": true
                }
            }
        },
        "response.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/response.Meta"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "localhost:3030",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "gank-ecommerce-auth",
	Description: "This is a doc for gank-ecommerce-auth.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
