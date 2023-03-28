// Package app GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package app

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
                "operationId": "app-category-all",
                "parameters": [
                    {
                        "type": "string",
                        "name": "pageToken",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/income-moneys": {
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
                    "IncomeMoney"
                ],
                "summary": "All",
                "operationId": "app-income-money-all",
                "parameters": [
                    {
                        "type": "string",
                        "name": "fromAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "pageToken",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "toAt",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
                    "IncomeMoney"
                ],
                "summary": "Create",
                "operationId": "app-income-money-create",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.IncomeMoneyBodyCreate"
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
        "/income-moneys/{id}": {
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
                    "IncomeMoney"
                ],
                "summary": "Detail",
                "operationId": "app-income-money-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Income money id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
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
                    "IncomeMoney"
                ],
                "summary": "Update",
                "operationId": "app-income-money-update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Income money id",
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
                            "$ref": "#/definitions/requestmodel.IncomeMoneyBodyUpdate"
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
        "/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login",
                "operationId": "user-login",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.UserBodyLogin"
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
        "/users/me": {
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
                "operationId": "user-get-me",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register",
                "operationId": "user-register",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.UserBodyRegister"
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
        "requestmodel.IncomeMoneyBodyCreate": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "money": {
                    "type": "number"
                },
                "note": {
                    "type": "string"
                }
            }
        },
        "requestmodel.IncomeMoneyBodyUpdate": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "money": {
                    "type": "number"
                },
                "note": {
                    "type": "string"
                }
            }
        },
        "requestmodel.UserBodyLogin": {
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
        "requestmodel.UserBodyRegister": {
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
                "password": {
                    "type": "string"
                },
                "phone": {
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
	BasePath:         "/app/expense",
	Schemes:          []string{},
	Title:            "Expense Tracker - App API",
	Description:      "All APIs for Expense app.\n\n******************************\n- Add description\n******************************\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
