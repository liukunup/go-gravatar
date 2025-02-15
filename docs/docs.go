// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/avatar": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "修改头像",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "头像模块"
                ],
                "summary": "修改头像",
                "parameters": [
                    {
                        "type": "file",
                        "description": "avatar",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image form other source",
                        "name": "jsonData",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "删除头像",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "头像模块"
                ],
                "summary": "删除头像",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.Response"
                        }
                    }
                }
            }
        },
        "/avatar/{hash}": {
            "get": {
                "description": "获取头像",
                "produces": [
                    "image/jpeg",
                    " image/png",
                    " image/webp"
                ],
                "tags": [
                    "头像模块"
                ],
                "summary": "获取头像",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "image"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "支持使用` + "`" + `用户名` + "`" + `/` + "`" + `邮箱` + "`" + `+` + "`" + `密码` + "`" + `来登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "账号登录",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.LoginResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "目前只支持邮箱注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.Response"
                        }
                    }
                }
            }
        },
        "/reset": {
            "post": {
                "description": "目前只支持邮箱重置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "重置密码",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.ResetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取用户信息(不包括用户头像)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.GetProfileResponse"
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
                "description": "修改用户信息(不包括用户头像)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.UpdateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "删除用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-gravatar_api_v1.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "go-gravatar_api_v1.GetProfileResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/go-gravatar_api_v1.GetProfileResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "go-gravatar_api_v1.GetProfileResponseData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "username@example.lan"
                },
                "nickname": {
                    "type": "string",
                    "example": "Billy"
                },
                "userId": {
                    "type": "string",
                    "example": "ExWFdl17WS"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "go-gravatar_api_v1.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "go-gravatar_api_v1.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/go-gravatar_api_v1.LoginResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "go-gravatar_api_v1.LoginResponseData": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        },
        "go-gravatar_api_v1.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "username@example.lan"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                }
            }
        },
        "go-gravatar_api_v1.ResetRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "username@example.lan"
                }
            }
        },
        "go-gravatar_api_v1.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "go-gravatar_api_v1.UpdateProfileRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "username@example.lan"
                },
                "nickname": {
                    "type": "string",
                    "example": "Billy"
                },
                "username": {
                    "type": "string",
                    "example": "username"
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
	Version:          "1.0.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Nunu Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
