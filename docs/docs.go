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
        "/menu": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ListMenu, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
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
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call UpdateMenu, return ok by json.",
                "parameters": [
                    {
                        "description": "User input 바꿀 메뉴 이름 toUpdate로 추가, 바꿀내용만 작성",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateMenuInput"
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
        "/order": {
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
                    },
                    {
                        "type": "string",
                        "description": "1은 현재 주문, 그 외 과거 주문",
                        "name": "cur",
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
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call UpdateOrder, return ok by json.",
                "parameters": [
                    {
                        "description": "수정할 주문 번호, 변경한 주문 메뉴 [{메뉴이름, 수량}]",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateOrderInput"
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
        "/order/admin": {
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
                        "description": "사업체 id",
                        "name": "id",
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
            },
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
                            "$ref": "#/definitions/controller.UpdateStateInput"
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
        "/order/make": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateOrder, return ok by json.",
                "parameters": [
                    {
                        "description": "주문자 이름,  메뉴 배열형태로 메뉴ID, 주문 수량 입력",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateOrderInput"
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
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ReadReviewControl, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "메뉴 id",
                        "name": "id",
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
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateReview, return ok by json.",
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
        "controller.CreateOrderInput": {
            "type": "object",
            "properties": {
                "bid": {
                    "type": "string"
                },
                "menu": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "menuID": {
                                "type": "string"
                            },
                            "number": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "orderer": {
                    "type": "string"
                }
            }
        },
        "controller.ReviewInput": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "menu_id": {
                    "type": "string"
                },
                "order_id": {
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
        "controller.UpdateMenuInput": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "controller.UpdateOrderInput": {
            "type": "object",
            "properties": {
                "menu": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "menuID": {
                                "type": "string"
                            },
                            "number": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "orderID": {
                    "type": "string"
                }
            }
        },
        "controller.UpdateStateInput": {
            "type": "object",
            "required": [
                "orderId",
                "state"
            ],
            "properties": {
                "orderId": {
                    "type": "string"
                },
                "state": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 1
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
