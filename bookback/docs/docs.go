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
        "/books": {
            "get": {
                "description": "Извлекает список всех книг",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Книги"
                ],
                "summary": "Получить список книг",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Book"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает новую книгу",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Книги"
                ],
                "summary": "Создать книгу",
                "parameters": [
                    {
                        "description": "Book object",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "description": "Извлекает книгу по ее ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Книги"
                ],
                "summary": "Получить книгу по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет книгу по ее ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Книги"
                ],
                "summary": "Обновить книгу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book object",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет книгу по ее ID",
                "tags": [
                    "Книги"
                ],
                "summary": "Удалить книгу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/chapters": {
            "get": {
                "description": "Извлекает список всех глав",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Главы"
                ],
                "summary": "Получить список глав",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Chapter"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает новую главу",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Главы"
                ],
                "summary": "Создать главу",
                "parameters": [
                    {
                        "description": "Chapter object",
                        "name": "chapter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Chapter"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Chapter"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/chapters/{id}": {
            "get": {
                "description": "Извлекает главу по ее ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Главы"
                ],
                "summary": "Получить главу по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID главы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Chapter"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет главу по ее ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Главы"
                ],
                "summary": "Обновить главу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID главы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Chapter object",
                        "name": "chapter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Chapter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Chapter"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет главу по ее ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Главы"
                ],
                "summary": "Удалить главу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID главы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Chapter"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Возвращает статус приложения",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Статус приложения"
                ],
                "summary": "Получить статус приложения",
                "responses": {
                    "200": {
                        "description": "healthy",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/pages": {
            "get": {
                "description": "Извлекает список всех страниц",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Страницы"
                ],
                "summary": "Получить список страниц",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Page"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает новую страницу",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Страницы"
                ],
                "summary": "Создать страницу",
                "parameters": [
                    {
                        "description": "Page object",
                        "name": "page",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Page"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Page"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/pages/{id}": {
            "get": {
                "description": "Извлекает страницу по ее ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Страницы"
                ],
                "summary": "Получить страницу по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID страницы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Page"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/paragraphs": {
            "get": {
                "description": "Извлекает список всех параграфов",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параграфы"
                ],
                "summary": "Получить список параграфов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Paragraph"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает новый параграф",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параграфы"
                ],
                "summary": "Создать параграф",
                "parameters": [
                    {
                        "description": "Paragraph object",
                        "name": "paragraph",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Paragraph"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Paragraph"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/paragraphs/{id}": {
            "get": {
                "description": "Извлекает параграф по его ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параграфы"
                ],
                "summary": "Получить параграф по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID параграфа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Paragraph"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет параграф по его ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Параграфы"
                ],
                "summary": "Обновить параграф",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID параграфа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Paragraph object",
                        "name": "paragraph",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Paragraph"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Paragraph"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет параграф по его ID",
                "tags": [
                    "Параграфы"
                ],
                "summary": "Удалить параграф",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID параграфа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Paragraph"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_public": {
                    "type": "boolean"
                },
                "owner": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Chapter": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_public": {
                    "type": "boolean"
                },
                "number": {
                    "type": "integer"
                },
                "pages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Page"
                    }
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Page": {
            "type": "object",
            "properties": {
                "chapter_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_public": {
                    "type": "boolean"
                },
                "text": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Paragraph": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_public": {
                    "type": "boolean"
                },
                "page_id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:7077",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Book API",
	Description:      "Это API для работы с книгами",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}