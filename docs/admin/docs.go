// Package admin GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package admin

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Dev team",
            "email": "tuanngo.472000@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "All",
                "operationId": "category-all",
                "parameters": [
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseCategoryAll"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create",
                "operationId": "category-create",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.CategoryBodyCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseCreate"
                        }
                    }
                }
            }
        },
        "/categories/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Detail",
                "operationId": "category-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseCategoryAdmin"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update",
                "operationId": "category-update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.CategoryBodyUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseUpdate"
                        }
                    }
                }
            }
        },
        "/categories/{id}/status": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "ChangeStatus",
                "operationId": "category-change-status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.CategoryChangeStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseChangeStatus"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Common"
                ],
                "summary": "Ping",
                "operationId": "ping",
                "responses": {}
            }
        },
        "/staffs/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Staff"
                ],
                "summary": "Login",
                "operationId": "staff-login",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.StaffBodyLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/staffs/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Staff"
                ],
                "summary": "GetMe",
                "operationId": "staff-get-me",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Staff"
                ],
                "summary": "Update",
                "operationId": "staff-update-info",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.StaffBodyUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/staffs/me/password": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Staff"
                ],
                "summary": "ChangePassword",
                "operationId": "staff-change-password",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.StaffBodyChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "ptime.TimeResponse": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                }
            }
        },
        "requestmodel.CategoryBodyCreate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "requestmodel.CategoryBodyUpdate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "requestmodel.CategoryChangeStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "requestmodel.StaffBodyChangePassword": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "requestmodel.StaffBodyLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "requestmodel.StaffBodyUpdate": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "responsemodel.ResponseCategoryAdmin": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "createdAt": {
                    "$ref": "#/definitions/ptime.TimeResponse"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "$ref": "#/definitions/ptime.TimeResponse"
                }
            }
        },
        "responsemodel.ResponseCategoryAll": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responsemodel.ResponseCategoryAdmin"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "responsemodel.ResponseChangeStatus": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "responsemodel.ResponseCreate": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                }
            }
        },
        "responsemodel.ResponseUpdate": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/admin/expense",
	Schemes:          []string{},
	Title:            "Expense Tracker - Admin API",
	Description:      "All APIs for Expense admin.\n\n******************************\n- Add description\n******************************\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
