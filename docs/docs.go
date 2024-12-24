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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/quizzes": {
            "post": {
                "description": "Creates a new quiz along with its questions and answers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Create a new quiz",
                "parameters": [
                    {
                        "description": "Quiz details",
                        "name": "quiz",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateQuizRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Quiz"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates the details of an existing quiz.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Update an existing quiz",
                "parameters": [
                    {
                        "description": "Updated quiz details",
                        "name": "quiz",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateQuizRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Quiz"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizzes/answer": {
            "post": {
                "description": "Submits an answer for the current question and updates the progression.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Answer a quiz question",
                "parameters": [
                    {
                        "description": "Answer details",
                        "name": "answerQuizQuestion",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AnswerQuizQuestionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AnswerQuizQuestionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizzes/begin": {
            "post": {
                "description": "Starts a quiz session for a user, initializing the progression with the first question.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Begin a quiz",
                "parameters": [
                    {
                        "description": "Quiz start details",
                        "name": "beginQuiz",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BeginQuizRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.BeginQuizResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizzes/questions/{id}": {
            "get": {
                "description": "Retrieves a question by its ID along with its answer options.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Get a quiz question by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.QuestionWithOptionsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizzes/submit": {
            "post": {
                "description": "Marks a quiz as finished and calculates the score based on correct answers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Finalize a quiz",
                "parameters": [
                    {
                        "description": "Quiz finalization details",
                        "name": "finalizeQuiz",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FinalizeQuizRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FinalizeQuizResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizzes/{id}": {
            "get": {
                "description": "Retrieves a quiz by its ID, including its questions and answers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Get a quiz by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quiz ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Quiz"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a quiz by its ID, including its associated questions and answers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quizzes"
                ],
                "summary": "Delete a quiz",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quiz ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieves a paginated list of users based on optional filters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get a list of users",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "List of user IDs",
                        "name": "id_list",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name to search for",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "size",
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
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new user in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates user details based on the provided request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "description": "Updated user details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Retrieves a user by their ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a user by their ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{userId}/quiz/{quizId}": {
            "get": {
                "description": "Retrieves the user's score for a specific quiz.",
                "tags": [
                    "Users"
                ],
                "summary": "Get user's score for a specific quiz",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Quiz ID",
                        "name": "quizId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Score"
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
        "/users/{userId}/quiz/{quizId}/ranking": {
            "get": {
                "description": "Retrieves the user's ranking, score, and percentage of quizzers they outperformed in a specific quiz.",
                "tags": [
                    "Users"
                ],
                "summary": "Get user's ranking by score in a specific quiz",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Quiz ID",
                        "name": "quizId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetUserRankingByScoreResponse"
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
        }
    },
    "definitions": {
        "models.Answer": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "optionId": {
                    "type": "integer"
                },
                "quizId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.AnswerQuizQuestionRequest": {
            "type": "object",
            "required": [
                "optionId",
                "progressionId"
            ],
            "properties": {
                "optionId": {
                    "type": "integer"
                },
                "progressionId": {
                    "type": "integer"
                }
            }
        },
        "models.AnswerQuizQuestionResponse": {
            "type": "object",
            "properties": {
                "progression": {
                    "$ref": "#/definitions/models.Progression"
                }
            }
        },
        "models.BeginQuizRequest": {
            "type": "object",
            "required": [
                "quizId",
                "userId"
            ],
            "properties": {
                "quizId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.BeginQuizResponse": {
            "type": "object",
            "properties": {
                "progression": {
                    "$ref": "#/definitions/models.Progression"
                }
            }
        },
        "models.CreateAnswerRequest": {
            "type": "object",
            "properties": {
                "isCorrect": {
                    "type": "boolean"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.CreateQuestionRequest": {
            "type": "object",
            "required": [
                "answers",
                "question"
            ],
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CreateAnswerRequest"
                    }
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "models.CreateQuizRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CreateQuestionRequest"
                    }
                }
            }
        },
        "models.CreateUserRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.FinalizeQuizRequest": {
            "type": "object",
            "required": [
                "progressionId"
            ],
            "properties": {
                "progressionId": {
                    "type": "integer"
                }
            }
        },
        "models.FinalizeQuizResponse": {
            "type": "object",
            "properties": {
                "score": {
                    "$ref": "#/definitions/models.Score"
                }
            }
        },
        "models.GetUserRankingByScoreResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "percent": {
                    "type": "number"
                },
                "rank": {
                    "type": "integer"
                },
                "userScore": {
                    "$ref": "#/definitions/models.Score"
                }
            }
        },
        "models.Option": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isCorrect": {
                    "type": "boolean"
                },
                "questionId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.OptionBase": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "questionId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.Progression": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "currentQuestionId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "isFinished": {
                    "type": "boolean"
                },
                "questionNumber": {
                    "type": "integer"
                },
                "quizId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Question": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Option"
                    }
                },
                "progressions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Progression"
                    }
                },
                "question": {
                    "type": "string"
                },
                "quizId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.QuestionWithOptionsResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OptionBase"
                    }
                },
                "question": {
                    "type": "string"
                },
                "quizId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Quiz": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Question"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Score": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "quizId": {
                    "type": "integer"
                },
                "score": {
                    "type": "number"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.UpdateQuizRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.UpdateUserRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Quiz Maker API",
	Description:      "Quiz Maker API for Fast Track Interview",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
