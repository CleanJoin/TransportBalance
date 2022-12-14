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
        "termsOfService": "https://github.com/CleanJoin/transportBalance/",
        "contact": {
            "name": "Github.com",
            "url": "https://github.com/CleanJoin/transportBalance/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/add": {
            "post": {
                "description": "Зачислить средства",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "addMoneyHandler",
                "parameters": [
                    {
                        "description": "RequestUser",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestMoveMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/docs": {
            "post": {
                "description": "Получить отчет для Бухгалтерии",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reports"
                ],
                "summary": "createGenDocHandler",
                "parameters": [
                    {
                        "description": "RequestUser",
                        "name": "Date",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.GenDoc"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/health": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Info"
                ],
                "summary": "Show the status of server.",
                "responses": {}
            }
        },
        "/api/money": {
            "post": {
                "description": "Получить данные о балансе",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "getMoneyUserHadler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/reduce": {
            "post": {
                "description": "Списать денежные средства",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "reduceMoneyHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestMoveMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/reduceReserve": {
            "post": {
                "description": "списание средств из Резерва",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "reduceReserveHandler",
                "parameters": [
                    {
                        "description": "RequestUser",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.ReserveMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/reserve": {
            "post": {
                "description": "Резерв денежных средств по определенной услуге",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "addMoneyToReserveHandler",
                "parameters": [
                    {
                        "description": "RequestUser",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.ReserveMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/transfer": {
            "post": {
                "description": "Перевести деньги от пользователя к пользователю",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "transferMoneyHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "TransactionsModel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.TransferMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/user": {
            "post": {
                "description": "Регистрация пользователя",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "userHandler",
                "parameters": [
                    {
                        "description": "RequestUser",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestUser"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "balance.GenDoc": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                }
            }
        },
        "balance.RequestMoveMoney": {
            "type": "object",
            "properties": {
                "money": {
                    "type": "number"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "balance.RequestUser": {
            "type": "object",
            "properties": {
                "paswword": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "balance.ReserveMoney": {
            "type": "object",
            "properties": {
                "money": {
                    "type": "number"
                },
                "orderId": {
                    "type": "integer"
                },
                "serviceId": {
                    "type": "integer"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "balance.TransferMoney": {
            "type": "object",
            "properties": {
                "Money": {
                    "type": "number"
                },
                "UserIdFrom": {
                    "type": "integer"
                },
                "UserIdTo": {
                    "type": "integer"
                }
            }
        },
        "balance.User": {
            "type": "object",
            "properties": {
                "userid": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger transportBalance",
	Description:      "This is a sample server transportBalance",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
