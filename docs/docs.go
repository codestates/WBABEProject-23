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
        "/menu/admin/modify": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ModifyMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "사업체 ID",
                        "name": "business_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User input 바꿀 메뉴 이름 toUpdate로 추가, 바꿀내용만 작성",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ModifyMenuInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/menu/admin/new": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call NewMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "사업체 ID",
                        "name": "business_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User input",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.NewMenuInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/menu/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call MenuList, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sort할 컬럼이름",
                        "name": "sort",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "order= 1은 오름찬순 그 외 내림차순 ",
                        "name": "order",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/menu/list/review": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call MenuReadReview, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "가게 사업체 id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "메뉴 이름",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/order/admin/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call AdminListOrderController, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "사업체 이름",
                        "name": "businessname",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/order/admin/update": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call UpdateState, return ok by json.",
                "parameters": [
                    {
                        "description": "주문 번호, 상태 ",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateOrderStateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/order/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ListOrder, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "유저이름",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/order/make": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call MakeOrder, return ok by json.",
                "parameters": [
                    {
                        "description": "주문자 이름, 주문 가게 이름, 메뉴 배열형태만 입력 ]",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/order/modify": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ModifyOrder, return ok by json.",
                "parameters": [
                    {
                        "description": "수정할 주문 번호, 변경한 주문 메뉴 [{메뉴이름, 수량}]",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ModifyOrderInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/review": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call MakeReview, return ok by json.",
                "parameters": [
                    {
                        "description": "리뷰",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReviewInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Controller": {
            "type": "object"
        },
        "controller.ModifyMenuInput": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "origin": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                },
                "toUpdate": {
                    "type": "string"
                }
            }
        },
        "controller.ModifyOrderInput": {
            "type": "object",
            "properties": {
                "menu": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MenuNum"
                    }
                },
                "orderID": {
                    "type": "string"
                }
            }
        },
        "controller.NewMenuInput": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "controller.ReviewInput": {
            "type": "object",
            "properties": {
                "businessID": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "menuName": {
                    "type": "string"
                },
                "orderID": {
                    "type": "string"
                },
                "orderer": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                }
            }
        },
        "controller.UpdateOrderStateInput": {
            "type": "object",
            "properties": {
                "orderId": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "model.MenuNum": {
            "type": "object",
            "properties": {
                "isReviewed": {
                    "type": "boolean"
                },
                "menuName": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                }
            }
        },
        "model.Order": {
            "type": "object",
            "properties": {
                "businessName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "menu": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MenuNum"
                    }
                },
                "orderID": {
                    "type": "integer"
                },
                "orderer": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
