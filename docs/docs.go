// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-03 23:22:40.693767842 +0800 CST m=+0.040913370

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/chatfuel/fried/{regional}": {
            "get": {
                "description": "以地區取得炸物資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chatfuel"
                ],
                "summary": "以地區取得炸物資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/chatfuel/pics/type/{type}/regional/{regional}": {
            "get": {
                "description": "以分類及地區取得圖片資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chatfuel"
                ],
                "summary": "以分類及地區取得圖片資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/chatfuel/restaurants/type/{type}/{regional}": {
            "get": {
                "description": "以分類及地區取得資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chatfuel"
                ],
                "summary": "以分類及地區取得資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/pics": {
            "get": {
                "description": "取得圖片列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "圖片"
                ],
                "summary": "取得圖片列表",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/pics-type/{type}/regional/{regional}": {
            "get": {
                "description": "以分類及地區取得圖片資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "圖片"
                ],
                "summary": "以分類及地區取得圖片資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/pics/{id}": {
            "get": {
                "description": "取得圖片資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "圖片"
                ],
                "summary": "取得圖片資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/restaurants": {
            "get": {
                "description": "取得餐廳列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "餐廳"
                ],
                "summary": "取得餐廳列表",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/restaurants-type/{type}/{regional}": {
            "get": {
                "description": "以分類及地區取得餐廳資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "餐廳"
                ],
                "summary": "以分類及地區取得餐廳資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/restaurants/{id}": {
            "get": {
                "description": "取得餐廳資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "餐廳"
                ],
                "summary": "取得餐廳資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "description": "取得用戶列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用戶"
                ],
                "summary": "取得用戶列表",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            },
            "post": {
                "description": "新增用戶",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用戶"
                ],
                "summary": "新增用戶",
                "parameters": [
                    {
                        "description": "用戶名稱",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "會員密碼",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "會員角色",
                        "name": "roles",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "用戶真實名稱",
                        "name": "real_name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "description": "取得用戶資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用戶"
                ],
                "summary": "取得用戶資料",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            },
            "put": {
                "description": "更新用戶資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用戶"
                ],
                "summary": "更新用戶資料",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用戶名稱",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "會員密碼",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "會員角色",
                        "name": "roles",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "用戶真實名稱",
                        "name": "real_name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            },
            "delete": {
                "description": "刪除用戶",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用戶"
                ],
                "summary": "刪除用戶",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        },
        "/auth": {
            "get": {
                "description": "登入驗證",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "驗證"
                ],
                "summary": "登入驗證",
                "parameters": [
                    {
                        "description": "username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
