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
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/querymodel.CategoryAll"
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
        "/incomes": {
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
                    "Income"
                ],
                "summary": "Create",
                "operationId": "create-income-money",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.IncomeBodyCreate"
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
        "/incomes/{id}": {
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
                    "Income"
                ],
                "summary": "Create",
                "operationId": "update-income-money",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.IncomeBodyUpdate"
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
        "querymodel.CategoryAll": {
            "type": "object",
            "properties": {
                "pageToken": {
                    "type": "string"
                },
                "sort": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "requestmodel.IncomeBodyCreate": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "money": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                }
            }
        },
        "requestmodel.IncomeBodyUpdate": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "money": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
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
