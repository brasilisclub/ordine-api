// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
                "description": "Login for a User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.AuthRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginUserSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register for a User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.AuthRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/auth.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            }
        },
        "/ordines": {
            "get": {
                "description": "Retrieve a list of all ordines",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordine"
                ],
                "summary": "Get all ordines",
                "responses": {
                    "200": {
                        "description": "List of ordines",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ordine.Ordine"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new ordine with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordine"
                ],
                "summary": "Create a new ordine",
                "parameters": [
                    {
                        "description": "Ordine data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ordine.OrdineRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created ordine",
                        "schema": {
                            "$ref": "#/definitions/ordine.Ordine"
                        }
                    },
                    "400": {
                        "description": "Invalid input or error creating ordine",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            }
        },
        "/ordines/{id}": {
            "get": {
                "description": "Retrieve an ordine by its unique ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordine"
                ],
                "summary": "Get ordine by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ordine ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ordine data",
                        "schema": {
                            "$ref": "#/definitions/ordine.Ordine"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing ordine by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordine"
                ],
                "summary": "Update an ordine",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ordine ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Ordine data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ordine.OrdineRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated ordine\"  // Retorna o ordine atualizado",
                        "schema": {
                            "$ref": "#/definitions/ordine.Ordine"
                        }
                    },
                    "400": {
                        "description": "Invalid input or error updating ordine",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an ordine based on the given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordine"
                ],
                "summary": "Delete ordine",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ordine id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Resource deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Error trying to delete ordine",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            }
        },
        "/ordines/{id}/products": {
            "post": {
                "description": "Add products to an existing ordine by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ordine"
                ],
                "summary": "Add products to an ordine",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ordine ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Products to add to the ordine",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ordine.OrderProductBody"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated ordine",
                        "schema": {
                            "$ref": "#/definitions/ordine.Ordine"
                        }
                    },
                    "400": {
                        "description": "Invalid input or error updating ordine",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Retrieve a list of all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/product.Product"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product based on the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid body or bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Retrieve a product based on the given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid ID or error retrieving product",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a product based on the given ID and provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update an existing product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid body or error updating product",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a product by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Error trying to delete the product",
                        "schema": {
                            "$ref": "#/definitions/utils.GenericResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AuthRequestBody": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.LoginUserSuccessResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "ordine.OrderProductBody": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "ordine.OrderProducts": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ordine_id": {
                    "type": "integer"
                },
                "product": {
                    "$ref": "#/definitions/product.Product"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "ordine.Ordine": {
            "type": "object",
            "properties": {
                "client_name": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ordine_products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ordine.OrderProducts"
                    }
                },
                "status": {
                    "type": "boolean"
                },
                "table": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "ordine.OrdineRequestBody": {
            "type": "object",
            "properties": {
                "client_name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "table": {
                    "type": "integer"
                }
            }
        },
        "product.Product": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "product.ProductRequestBody": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "utils.GenericResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
