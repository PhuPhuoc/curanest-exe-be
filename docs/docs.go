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
        "/nurses": {
            "get": {
                "description": "get list nurses (card)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nurses"
                ],
                "summary": "get list nurses (card)",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "create new nurse",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nurses"
                ],
                "summary": "create new nurse",
                "parameters": [
                    {
                        "description": "Nurse creation data",
                        "name": "creation_form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/nursemodel.NurseCreationModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message success",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/nurses/{nurse_id}": {
            "get": {
                "description": "get list nurses (card)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nurses"
                ],
                "summary": "get list nurses (card)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nurse ID",
                        "name": "nurse_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "user",
                            "admin"
                        ],
                        "type": "string",
                        "description": "Role",
                        "name": "role",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/techniques": {
            "get": {
                "description": "get techniques",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "techniques"
                ],
                "summary": "get techniques",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "create new technique",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "techniques"
                ],
                "summary": "create new technique",
                "parameters": [
                    {
                        "description": "new technique information",
                        "name": "technique_form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/techniquemodel.TechniqueCreationModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message success",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/users/authentication": {
            "post": {
                "description": "User login sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User login sign in",
                "parameters": [
                    {
                        "description": "User log in info",
                        "name": "user-login-data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usermodel.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/users/sign-up": {
            "post": {
                "description": "register new user's account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "register new user's account",
                "parameters": [
                    {
                        "description": "User register data",
                        "name": "register_form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usermodel.RegisterForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message success",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/users/{user_id}/create-customer-profile": {
            "post": {
                "description": "create new customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "create new customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer register data",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customermodel.CustomerModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message success",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "customermodel.CustomerModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "citizen_id": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "ward": {
                    "type": "string"
                }
            }
        },
        "nursemodel.NurseCreationModel": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "certificate": {
                    "type": "string"
                },
                "citizen_id": {
                    "type": "string"
                },
                "current_workplace": {
                    "type": "string"
                },
                "education_level": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expertise": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "slogan": {
                    "type": "string"
                },
                "techniques": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "work_experience": {
                    "type": "string"
                }
            }
        },
        "techniquemodel.TechniqueCreationModel": {
            "type": "object",
            "properties": {
                "estimated_time": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usermodel.LoginForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "usermodel.RegisterForm": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Curanest API for EXE201",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
