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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "apiのヘルスチェックを行う。ルーティングが正常に登録されているかを確かめる。",
                "tags": [
                    "HealthCheck"
                ],
                "summary": "apiのヘルスチェックを行う",
                "responses": {
                    "200": {
                        "description": "Health check message",
                        "schema": {
                            "$ref": "#/definitions/presenter.SuccessResponse-health_healthResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "新しいユーザーを登録する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "ユーザーの登録",
                "parameters": [
                    {
                        "description": "ユーザー登録のための情報",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.PostUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "登録されたユーザーの情報",
                        "schema": {
                            "$ref": "#/definitions/presenter.SuccessResponse-user_PostUserResponse"
                        }
                    },
                    "400": {
                        "description": "不正なリクエスト",
                        "schema": {
                            "$ref": "#/definitions/presenter.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "内部サーバーエラー",
                        "schema": {
                            "$ref": "#/definitions/presenter.FailureResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "description": "ユーザーを退会させ、ユーザー情報を削除する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "ユーザーの退会",
                "parameters": [
                    {
                        "type": "string",
                        "description": "削除するユーザーのid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "不正なリクエスト",
                        "schema": {
                            "$ref": "#/definitions/presenter.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "内部サーバーエラー",
                        "schema": {
                            "$ref": "#/definitions/presenter.FailureResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "全てのユーザーのID・名前をリストで取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "全ユーザーを取得する",
                "responses": {
                    "200": {
                        "description": "登録されたユーザーの情報",
                        "schema": {
                            "$ref": "#/definitions/presenter.SuccessResponse-array_user_GetUsersResponse"
                        }
                    },
                    "400": {
                        "description": "不正なリクエスト",
                        "schema": {
                            "$ref": "#/definitions/presenter.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "内部サーバーエラー",
                        "schema": {
                            "$ref": "#/definitions/presenter.FailureResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "health.healthResponse": {
            "type": "object",
            "properties": {
                "health_check": {
                    "type": "string"
                }
            }
        },
        "presenter.FailureResponse": {
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
        "presenter.SuccessResponse-array_user_GetUsersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.GetUsersResponse"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "presenter.SuccessResponse-health_healthResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/health.healthResponse"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "presenter.SuccessResponse-user_PostUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/user.PostUserResponse"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "user.GetUsersResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.PostUserRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.PostUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
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
	Title:            "TodoRestAPI",
	Description:      "This is TodoARestPI by golang.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
