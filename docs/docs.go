// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Auth login",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Auth register",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.AuthRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/resp.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/transactions/count/{userId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Count History Transaction",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/transactions/topUp": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "parameters": [
                    {
                        "description": "TopUp",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.TopUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Transactions"
                        }
                    }
                }
            }
        },
        "/transactions/transfer": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "parameters": [
                    {
                        "description": "Transfer",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.TransferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Transactions"
                        }
                    }
                }
            }
        },
        "/transactions/{userId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get History Transaction",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.GetTransactionsResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Users"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "Update Personal Information",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UpdateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.UpdateAccountRespone"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "Change Password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UpdatePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.UpdatePasswordResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Disable Account",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.DisableAccountResponse"
                        }
                    }
                }
            }
        },
        "/users/{phoneNumber}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User PhoneNumber",
                        "name": "phoneNumber",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Users"
                        }
                    }
                }
            }
        },
        "/wallets/{userId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get Wallet",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Wallet"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Transactions": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "createAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "paymentMethodID": {
                    "type": "string"
                },
                "sourceWalletID": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "model.Users": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "disableAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "passwordConfirm": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "model.Wallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "rekeningUser": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "req.AuthLoginRequest": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "loginOption": {
                    "$ref": "#/definitions/req.loginOption"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "req.AuthRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "fullName",
                "password",
                "passwordConfirm",
                "phoneNumber",
                "userName"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                },
                "password": {
                    "type": "string"
                },
                "passwordConfirm": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 10
                },
                "userName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "req.TopUpRequest": {
            "type": "object",
            "required": [
                "amount",
                "paymentMethodId",
                "userId",
                "walletID"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "minimum": 10000
                },
                "paymentMethodId": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "walletID": {
                    "type": "string"
                }
            }
        },
        "req.TransferRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "destination_wallet_id": {
                    "type": "string"
                },
                "payment_method_id": {
                    "type": "string"
                },
                "source_user_id": {
                    "type": "string"
                },
                "source_wallet_id": {
                    "type": "string"
                }
            }
        },
        "req.UpdateAccountRequest": {
            "type": "object",
            "required": [
                "email",
                "fullName",
                "id",
                "phoneNumber",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                },
                "id": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 10
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "req.UpdatePasswordRequest": {
            "type": "object",
            "required": [
                "currentPassword",
                "newPassword",
                "newPasswordConfirm",
                "userName"
            ],
            "properties": {
                "currentPassword": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string"
                },
                "newPasswordConfirm": {
                    "type": "string"
                },
                "userName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "req.loginOption": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "resp.DisableAccountResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "resp.GetTransactionsResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "destination_wallet_id": {
                    "type": "string"
                },
                "id_transaction": {
                    "type": "string"
                },
                "payment_method": {
                    "$ref": "#/definitions/resp.paymentMethod"
                },
                "time_of_transaction": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/resp.user"
                },
                "wallet": {
                    "$ref": "#/definitions/resp.wallet"
                }
            }
        },
        "resp.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "resp.RegisterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "resp.UpdateAccountRespone": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "resp.UpdatePasswordResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "resp.paymentMethod": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "resp.user": {
            "type": "object",
            "properties": {
                "user_name": {
                    "type": "string"
                }
            }
        },
        "resp.wallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "rekening_user": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "dompet-online",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
