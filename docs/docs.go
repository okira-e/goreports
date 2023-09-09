// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/report/delete": {
            "delete": {
                "description": "Delete a report",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "Delete a report",
                "parameters": [
                    {
                        "description": "The name of the report to be deleted",
                        "name": "reportName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/report/list": {
            "get": {
                "description": "List all stored reports",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "List all reports",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/report/render": {
            "post": {
                "description": "Render a report",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "Render a report",
                "parameters": [
                    {
                        "description": "The name of the report",
                        "name": "reportName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The parameters injected inside the report body to be passed at runtime",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "The printing options to be used in the report",
                        "name": "printingOptions",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.PrintingOptions"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/report/save": {
            "post": {
                "description": "Save a report",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "Save a report",
                "parameters": [
                    {
                        "description": "The name of the report",
                        "name": "reportName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The title of the report",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The description of the report",
                        "name": "description",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The body of the report",
                        "name": "` + "`" + `body` + "`" + `",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The header of the report",
                        "name": "header",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The footer of the report",
                        "name": "footer",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        }
    },
    "definitions": {
        "types.PageNumbersOptions": {
            "type": "object",
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "position": {
                    "type": "string"
                }
            }
        },
        "types.PrintingOptions": {
            "type": "object",
            "properties": {
                "landscape": {
                    "type": "boolean"
                },
                "marginBottom": {
                    "type": "integer"
                },
                "marginLeft": {
                    "type": "integer"
                },
                "marginRight": {
                    "type": "integer"
                },
                "marginTop": {
                    "type": "integer"
                },
                "pageNumbers": {
                    "$ref": "#/definitions/types.PageNumbersOptions"
                },
                "paperSize": {
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
