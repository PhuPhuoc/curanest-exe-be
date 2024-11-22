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
        "/appointments": {
            "post": {
                "description": "create new appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointments"
                ],
                "summary": "create new appointment",
                "parameters": [
                    {
                        "description": "Nurse creation data",
                        "name": "creation_form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appointmentmodel.RegisterAppoinment"
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
        "/appointments/card/{account_id}": {
            "get": {
                "description": "get appointment (card)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointments"
                ],
                "summary": "get appointment (card)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "patient",
                            "nurse"
                        ],
                        "type": "string",
                        "description": "Role",
                        "name": "role",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date from",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date to",
                        "name": "to",
                        "in": "query",
                        "required": true
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
        "/appointments/{appointment_id}": {
            "get": {
                "description": "get appointment detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointments"
                ],
                "summary": "get appointment detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Appointment ID",
                        "name": "appointment_id",
                        "in": "path",
                        "required": true
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
        "/appointments/{appointment_id}/confirm": {
            "post": {
                "description": "nurse confirm appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointments"
                ],
                "summary": "nurse confirm appointment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Appointment ID",
                        "name": "appointment_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "confirmed",
                            "cancel"
                        ],
                        "type": "string",
                        "description": "confirm",
                        "name": "confirm",
                        "in": "query",
                        "required": true
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
        "/customers/{customer_id}": {
            "get": {
                "description": "get customer information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "get customer information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "customer_id",
                        "in": "path",
                        "required": true
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
        "/customers/{customer_id}/create-patient-profile": {
            "post": {
                "description": "create new customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patients"
                ],
                "summary": "create new customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "customer_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Patient register data",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/patientmodel.PatientCreationModel"
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
        "/customers/{customer_id}/patients": {
            "get": {
                "description": "get list patients of customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patients"
                ],
                "summary": "get list patients of customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "customer_id",
                        "in": "path",
                        "required": true
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
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by nurse's full name",
                        "name": "full_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by nurse's phone number",
                        "name": "phone_number",
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
                "description": "get nurse detail (card)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nurses"
                ],
                "summary": "get nurse detail (card)",
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
        "/nurses/{nurse_id}/get-suitable-time-frames": {
            "get": {
                "description": "get suitable time frames for booking services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nurses"
                ],
                "summary": "get suitable time frames for booking services",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nurse ID",
                        "name": "nurse_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date from",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date to",
                        "name": "to",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "total minute",
                        "name": "total_minute",
                        "in": "query",
                        "required": true
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
        "/nurses/{nurse_id}/get-weekly-work-schedule": {
            "get": {
                "description": "get weeky shedule of nurses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nurses"
                ],
                "summary": "get weeky shedule of nurses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nurse ID",
                        "name": "nurse_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date from",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date to",
                        "name": "to",
                        "in": "query",
                        "required": true
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
        "/nurses/{nurse_id}/register-weekly-work-schedule": {
            "post": {
                "description": "register for nurse's work schedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nurses"
                ],
                "summary": "register for nurse's work schedule",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nurse ID",
                        "name": "nurse_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Nurse shift information",
                        "name": "creation_form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/nursemodel.WeeklyWorkSchedule"
                        }
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
                    "customers"
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
        "appointmentmodel.RegisterAppoinment": {
            "type": "object",
            "properties": {
                "appointment_date": {
                    "type": "string"
                },
                "listNurseWorkSchedules": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "nurse_id": {
                    "type": "string"
                },
                "patient_id": {
                    "type": "string"
                },
                "techniques": {
                    "type": "string"
                },
                "time_from_to": {
                    "type": "string"
                },
                "total_fee": {
                    "type": "integer"
                }
            }
        },
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
        "nursemodel.Shift": {
            "type": "object",
            "properties": {
                "shift_date": {
                    "type": "string",
                    "example": "2024-11-20"
                },
                "shift_from": {
                    "type": "string",
                    "example": "00:00:00"
                },
                "shift_to": {
                    "type": "string",
                    "example": "00:00:00"
                }
            }
        },
        "nursemodel.WeeklyWorkSchedule": {
            "type": "object",
            "properties": {
                "shifts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/nursemodel.Shift"
                    }
                },
                "week_from": {
                    "type": "string",
                    "example": "2024-11-17"
                },
                "week_to": {
                    "type": "string",
                    "example": "2024-11-23"
                }
            }
        },
        "patientmodel.PatientCreationModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatar": {
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
                "medical_description": {
                    "type": "string"
                },
                "note_for_nurses": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "techniques": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ward": {
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
