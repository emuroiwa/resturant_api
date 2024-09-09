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
        "/dishes": {
            "get": {
                "description": "Retrieve a list of all available dishes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes"
                ],
                "summary": "List all dishes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Dish"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the details of an existing dish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes"
                ],
                "summary": "Update a dish",
                "parameters": [
                    {
                        "description": "Dish data",
                        "name": "dish",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Dish"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.Dish"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Submit a new dish that will be processed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes"
                ],
                "summary": "Create a new dish",
                "parameters": [
                    {
                        "description": "Dish data",
                        "name": "dish",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Dish"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Dish submitted successfully, it will be processed shortly.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dishes/search": {
            "get": {
                "description": "Search for dishes by name or other criteria",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes"
                ],
                "summary": "Search for dishes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Dish"
                            }
                        }
                    },
                    "400": {
                        "description": "Query parameter is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dishes/{dish_id}/ratings": {
            "get": {
                "description": "Retrieve all ratings for a specific dish by its UUID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ratings"
                ],
                "summary": "Get ratings for a dish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dish ID",
                        "name": "dish_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Rating"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid dish ID format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Submit a rating for a specific dish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ratings"
                ],
                "summary": "Create a new rating for a dish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dish ID",
                        "name": "dish_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Rating data",
                        "name": "rating",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Rating"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Rating submitted successfully, it will be processed shortly.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input or dish ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dishes/{id}": {
            "get": {
                "description": "Retrieve the details of a specific dish by its UUID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes"
                ],
                "summary": "Get a dish by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dish ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dish"
                        }
                    },
                    "400": {
                        "description": "Invalid UUID format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dish not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a specific dish by its UUID",
                "tags": [
                    "Dishes"
                ],
                "summary": "Delete a dish by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dish ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid UUID format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dish not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Dish": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "restaurantID": {
                    "type": "string"
                }
            }
        },
        "models.Rating": {
            "type": "object",
            "required": [
                "dishID",
                "score",
                "userID"
            ],
            "properties": {
                "dishID": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "review": {
                    "type": "string",
                    "maxLength": 500
                },
                "score": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                },
                "userID": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Restaurant API",
	Description:      "This is a sample server for a restaurant management system.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
