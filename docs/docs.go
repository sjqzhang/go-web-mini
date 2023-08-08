// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://localhost",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/access/list/user/:userId": {
            "get": {
                "tags": [
                    "menu"
                ],
                "responses": {}
            }
        },
        "/api/create": {
            "post": {
                "tags": [
                    "api"
                ],
                "responses": {}
            }
        },
        "/api/delete/batch": {
            "delete": {
                "tags": [
                    "api"
                ],
                "responses": {}
            }
        },
        "/api/list": {
            "get": {
                "tags": [
                    "api"
                ],
                "responses": {}
            }
        },
        "/api/tree": {
            "get": {
                "tags": [
                    "api"
                ],
                "responses": {}
            }
        },
        "/api/update/:apiId": {
            "patch": {
                "tags": [
                    "api"
                ],
                "responses": {}
            }
        },
        "/branch": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "查询单个branch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "branch_name",
                        "name": "branch_name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "branch_type",
                        "name": "branch_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "commit_id",
                        "name": "commit_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "commit_time",
                        "name": "commit_time",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "commit_title",
                        "name": "commit_title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "committer",
                        "name": "committer",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "creator",
                        "name": "creator",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "is_dev",
                        "name": "is_dev",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "jira_key",
                        "name": "jira_key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "第几页",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页多少条",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "repo",
                        "name": "repo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "sync_time",
                        "name": "sync_time",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ListBranchResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "创建branch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "branch_name",
                        "name": "branch_name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "branch_type",
                        "name": "branch_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "commit_id",
                        "name": "commit_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "commit_time",
                        "name": "commit_time",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "commit_title",
                        "name": "commit_title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "committer",
                        "name": "committer",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "creator",
                        "name": "creator",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "is_dev",
                        "name": "is_dev",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "jira_key",
                        "name": "jira_key",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "repo",
                        "name": "repo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "sync_time",
                        "name": "sync_time",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.CreateBranchResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "批量删除branch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branch"
                ],
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "待编号",
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.DeleteBranchResponse"
                        }
                    }
                }
            }
        },
        "/branch/:id": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "查询branch列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branch"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "待编号",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.GetBranchResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "更新branch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "branch_name",
                        "name": "branch_name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "branch_type",
                        "name": "branch_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "commit_id",
                        "name": "commit_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "commit_time",
                        "name": "commit_time",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "commit_title",
                        "name": "commit_title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "committer",
                        "name": "committer",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "creator",
                        "name": "creator",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "is_dev",
                        "name": "is_dev",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "jira_key",
                        "name": "jira_key",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "repo",
                        "name": "repo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "sync_time",
                        "name": "sync_time",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.UpdateBranchResponse"
                        }
                    }
                }
            }
        },
        "/log/operation/delete/batch": {
            "delete": {
                "tags": [
                    "log"
                ],
                "responses": {}
            }
        },
        "/log/operation/list": {
            "get": {
                "tags": [
                    "log"
                ],
                "responses": {}
            }
        },
        "/menu/access/tree/:userId": {
            "get": {
                "tags": [
                    "menu"
                ],
                "responses": {}
            }
        },
        "/menu/create": {
            "post": {
                "tags": [
                    "menu"
                ],
                "responses": {}
            }
        },
        "/menu/delete/batch": {
            "delete": {
                "tags": [
                    "menu"
                ],
                "responses": {}
            }
        },
        "/menu/list": {
            "get": {
                "tags": [
                    "menu"
                ],
                "responses": {}
            }
        },
        "/menu/tree": {
            "get": {
                "tags": [
                    "menu"
                ],
                "responses": {}
            }
        },
        "/menu/update/:menuId": {
            "patch": {
                "tags": [
                    "menu"
                ],
                "responses": {}
            }
        },
        "/news": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "查询单个新闻",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻内容",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "新闻创建者",
                        "name": "creator",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "第几页",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页多少条",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "新闻标题",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ListNewsResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "创建新闻",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻内容",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "新闻创建者",
                        "name": "creator",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "新闻标题",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.CreateNewsResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "批量删除新闻",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "待编号",
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.DeleteNewsResponse"
                        }
                    }
                }
            }
        },
        "/news/:id": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "查询新闻列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "待编号",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.GetNewsResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "更新新闻",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻内容",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "新闻创建者",
                        "name": "creator",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "新闻标题",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.UpdateNewsResponse"
                        }
                    }
                }
            }
        },
        "/role/apis/get/:roleId": {
            "get": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/role/apis/update/:roleId": {
            "patch": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/role/create": {
            "post": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/role/delete/batch": {
            "delete": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/role/list": {
            "get": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/role/menus/get/:roleId": {
            "get": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/role/menus/update/:roleId": {
            "patch": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/role/update/:roleId": {
            "patch": {
                "tags": [
                    "role"
                ],
                "responses": {}
            }
        },
        "/user": {
            "post": {
                "tags": [
                    "user"
                ],
                "responses": {}
            }
        },
        "/user/changePwd": {
            "put": {
                "tags": [
                    "user"
                ],
                "responses": {}
            }
        },
        "/user/delete/batch": {
            "delete": {
                "tags": [
                    "user"
                ],
                "responses": {}
            }
        },
        "/user/info": {
            "post": {
                "tags": [
                    "user"
                ],
                "responses": {}
            }
        },
        "/user/list": {
            "get": {
                "tags": [
                    "user"
                ],
                "responses": {}
            }
        },
        "/user/update/:userId": {
            "patch": {
                "tags": [
                    "user"
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "vo.Branch": {
            "type": "object",
            "properties": {
                "branch_name": {
                    "description": "branch_name",
                    "type": "string"
                },
                "branch_type": {
                    "description": "branch_type",
                    "type": "integer"
                },
                "commit_id": {
                    "description": "commit_id",
                    "type": "string"
                },
                "commit_time": {
                    "description": "commit_time",
                    "type": "integer"
                },
                "commit_title": {
                    "description": "commit_title",
                    "type": "string"
                },
                "committer": {
                    "description": "committer",
                    "type": "string"
                },
                "created_at": {
                    "description": "created_at",
                    "type": "string"
                },
                "creator": {
                    "description": "creator",
                    "type": "string"
                },
                "deleted_at": {
                    "description": "deleted_at",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "is_dev": {
                    "description": "is_dev",
                    "type": "integer"
                },
                "jira_key": {
                    "description": "jira_key",
                    "type": "string"
                },
                "repo": {
                    "description": "repo",
                    "type": "string"
                },
                "sync_time": {
                    "description": "sync_time",
                    "type": "integer"
                },
                "updated_at": {
                    "description": "updated_at",
                    "type": "string"
                }
            }
        },
        "vo.CreateBranchResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.Branch"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.CreateNewsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.News"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.DeleteBranchResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "type": "integer"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.DeleteNewsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "type": "integer"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.GetBranchResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.Branch"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.GetNewsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.News"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.ListBranchResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.PagerBranch"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.ListNewsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.PagerNews"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.News": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "新闻内容",
                    "type": "string"
                },
                "created_at": {
                    "description": "created_at",
                    "type": "string"
                },
                "creator": {
                    "description": "新闻创建者",
                    "type": "string"
                },
                "deleted_at": {
                    "description": "deleted_at",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "title": {
                    "description": "新闻标题",
                    "type": "string"
                },
                "updated_at": {
                    "description": "updated_at",
                    "type": "string"
                }
            }
        },
        "vo.PagerBranch": {
            "type": "object",
            "properties": {
                "extra": {
                    "type": "object",
                    "additionalProperties": true
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.Branch"
                    }
                },
                "pageNum": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "vo.PagerNews": {
            "type": "object",
            "properties": {
                "extra": {
                    "type": "object",
                    "additionalProperties": true
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.News"
                    }
                },
                "pageNum": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "vo.UpdateBranchResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.Branch"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "vo.UpdateNewsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，非0:失败",
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/vo.News"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
