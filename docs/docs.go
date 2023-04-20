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
        "/admin/admin_info": {
            "get": {
                "description": "管理员信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员信息",
                "operationId": "/admin/admin_info",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminInfoOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/change_pwd": {
            "post": {
                "description": "修改密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "修改密码",
                "operationId": "/admin/change_pwd",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ChangePwdInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/login": {
            "post": {
                "description": "管理员登陆",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员登陆",
                "operationId": "/admin_login/login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AdminLoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminLoginOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/logout": {
            "get": {
                "description": "管理员退出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员退出",
                "operationId": "/admin_login/logout",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_add_grpc": {
            "post": {
                "description": "grpc服务添加",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "grpc服务添加",
                "operationId": "/service/service_add_grpc",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceAddGrpcInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_add_http": {
            "post": {
                "description": "添加HTTP服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "添加HTTP服务",
                "operationId": "/service/service_add_http",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceAddHTTPInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_add_tcp": {
            "post": {
                "description": "tcp服务添加",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "tcp服务添加",
                "operationId": "/service/service_add_tcp",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceAddTcpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_delete": {
            "get": {
                "description": "服务删除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "服务删除",
                "operationId": "/service/service_delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "服务ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_list": {
            "get": {
                "description": "服务列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "服务列表",
                "operationId": "/service/service_list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "info",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页个数",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "当前页数",
                        "name": "page_no",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ServiceListOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_update_grpc": {
            "post": {
                "description": "grpc服务更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "grpc服务更新",
                "operationId": "/service/service_update_grpc",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceUpdateGrpcInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_update_http": {
            "post": {
                "description": "修改HTTP服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "修改HTTP服务",
                "operationId": "/service/service_update_http",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceUpdateHTTPInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_update_tcp": {
            "post": {
                "description": "tcp服务更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "tcp服务更新",
                "operationId": "/service/service_update_tcp",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceUpdateTcpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AdminInfoOutput": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "login_time": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.AdminLoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "description": "管理员用户名",
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dto.AdminLoginOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "token",
                    "type": "string",
                    "example": "token"
                }
            }
        },
        "dto.ChangePwdInput": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "dto.ServiceAddGrpcInput": {
            "type": "object",
            "required": [
                "ip_list",
                "port",
                "service_desc",
                "service_name",
                "weight_list"
            ],
            "properties": {
                "black_list": {
                    "type": "string"
                },
                "clientip_flow_limit": {
                    "type": "integer"
                },
                "forbid_list": {
                    "type": "string"
                },
                "header_transfor": {
                    "type": "string"
                },
                "ip_list": {
                    "type": "string"
                },
                "open_auth": {
                    "type": "integer"
                },
                "port": {
                    "type": "integer",
                    "maximum": 8999,
                    "minimum": 8001
                },
                "round_type": {
                    "type": "integer"
                },
                "service_desc": {
                    "type": "string"
                },
                "service_flow_limit": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "weight_list": {
                    "type": "string"
                },
                "white_host_name": {
                    "type": "string"
                },
                "white_list": {
                    "type": "string"
                }
            }
        },
        "dto.ServiceAddHTTPInput": {
            "type": "object"
        },
        "dto.ServiceAddTcpInput": {
            "type": "object",
            "required": [
                "ip_list",
                "port",
                "service_desc",
                "service_name",
                "weight_list"
            ],
            "properties": {
                "black_list": {
                    "type": "string"
                },
                "clientip_flow_limit": {
                    "type": "integer"
                },
                "forbid_list": {
                    "type": "string"
                },
                "header_transfor": {
                    "type": "string"
                },
                "ip_list": {
                    "type": "string"
                },
                "open_auth": {
                    "type": "integer"
                },
                "port": {
                    "type": "integer",
                    "maximum": 8999,
                    "minimum": 8001
                },
                "round_type": {
                    "type": "integer"
                },
                "service_desc": {
                    "type": "string"
                },
                "service_flow_limit": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "weight_list": {
                    "type": "string"
                },
                "white_host_name": {
                    "type": "string"
                },
                "white_list": {
                    "type": "string"
                }
            }
        },
        "dto.ServiceListItemOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "id",
                    "type": "integer",
                    "example": 1
                },
                "load_type": {
                    "description": "类型",
                    "type": "integer",
                    "example": 1
                },
                "qpd": {
                    "description": "qpd",
                    "type": "integer",
                    "example": 1
                },
                "qps": {
                    "description": "qps",
                    "type": "integer",
                    "example": 1
                },
                "service_addr": {
                    "description": "服务地址",
                    "type": "string",
                    "example": ""
                },
                "service_desc": {
                    "description": "服务描述",
                    "type": "string",
                    "example": ""
                },
                "service_name": {
                    "description": "服务名称",
                    "type": "string",
                    "example": ""
                },
                "total_node": {
                    "description": "节点数",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.ServiceListOutput": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ServiceListItemOutput"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "dto.ServiceUpdateGrpcInput": {
            "type": "object",
            "required": [
                "id",
                "ip_list",
                "port",
                "service_desc",
                "service_name",
                "weight_list"
            ],
            "properties": {
                "black_list": {
                    "type": "string"
                },
                "clientip_flow_limit": {
                    "type": "integer"
                },
                "forbid_list": {
                    "type": "string"
                },
                "header_transfor": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ip_list": {
                    "type": "string"
                },
                "open_auth": {
                    "type": "integer"
                },
                "port": {
                    "type": "integer",
                    "maximum": 8999,
                    "minimum": 8001
                },
                "round_type": {
                    "type": "integer"
                },
                "service_desc": {
                    "type": "string"
                },
                "service_flow_limit": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "weight_list": {
                    "type": "string"
                },
                "white_host_name": {
                    "type": "string"
                },
                "white_list": {
                    "type": "string"
                }
            }
        },
        "dto.ServiceUpdateHTTPInput": {
            "type": "object"
        },
        "dto.ServiceUpdateTcpInput": {
            "type": "object",
            "required": [
                "id",
                "ip_list",
                "port",
                "service_desc",
                "service_name",
                "weight_list"
            ],
            "properties": {
                "black_list": {
                    "type": "string"
                },
                "clientip_flow_limit": {
                    "type": "integer"
                },
                "forbid_list": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ip_list": {
                    "type": "string"
                },
                "open_auth": {
                    "type": "integer"
                },
                "port": {
                    "type": "integer",
                    "maximum": 8999,
                    "minimum": 8001
                },
                "round_type": {
                    "type": "integer"
                },
                "service_desc": {
                    "type": "string"
                },
                "service_flow_limit": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "weight_list": {
                    "type": "string"
                },
                "white_host_name": {
                    "type": "string"
                },
                "white_list": {
                    "type": "string"
                }
            }
        },
        "middleware.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/middleware.ResponseCode"
                },
                "data": {},
                "msg": {
                    "type": "string"
                },
                "stack": {},
                "trace_id": {}
            }
        },
        "middleware.ResponseCode": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                401,
                1000,
                2001
            ],
            "x-enum-varnames": [
                "SuccessCode",
                "UndefErrorCode",
                "ValidErrorCode",
                "InternalErrorCode",
                "InvalidRequestErrorCode",
                "CustomizeCode",
                "GROUPALL_SAVE_FLOWERROR"
            ]
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
