// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "text/plain"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This API is dedicated to store objects to a bucket.",
    "title": "S3",
    "contact": {
      "name": "Houssem Aloulou"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/objects/{bucket}": {
      "put": {
        "tags": [
          "Object Management"
        ],
        "summary": "Add an object in a bucket.",
        "operationId": "putObject",
        "parameters": [
          {
            "type": "string",
            "description": "Bucket name.",
            "name": "bucket",
            "in": "path",
            "required": true
          },
          {
            "description": "object content.",
            "name": "http-text",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Ok.",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "int64"
                }
              }
            }
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Not authenticated."
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Object not found."
          },
          "500": {
            "description": "Failed to put object."
          },
          "503": {
            "description": "Not available."
          }
        }
      }
    },
    "/objects/{bucket}/{objectID}": {
      "get": {
        "tags": [
          "Object Management"
        ],
        "summary": "Get object by bucket name and object id.",
        "operationId": "getObject",
        "parameters": [
          {
            "type": "string",
            "description": "Bucket name.",
            "name": "bucket",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "object id.",
            "name": "objectID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok.",
            "schema": {
              "type": "object",
              "required": [
                "object"
              ],
              "properties": {
                "object": {
                  "$ref": "#/definitions/object"
                }
              }
            }
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Not authenticated."
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Object not found."
          },
          "500": {
            "description": "Failed to list groups."
          },
          "503": {
            "description": "Not available."
          }
        }
      },
      "delete": {
        "tags": [
          "Object Management"
        ],
        "summary": "delete an object.",
        "operationId": "deleteObject",
        "parameters": [
          {
            "type": "string",
            "description": "Bucket name.",
            "name": "bucket",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "object id.",
            "name": "objectID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Not authenticated."
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Object not found."
          },
          "500": {
            "description": "Failed to delete object."
          },
          "503": {
            "description": "Not available."
          }
        }
      }
    }
  },
  "definitions": {
    "object": {
      "required": [
        "id",
        "data"
      ],
      "properties": {
        "data": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "text/plain"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This API is dedicated to store objects to a bucket.",
    "title": "S3",
    "contact": {
      "name": "Houssem Aloulou"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/objects/{bucket}": {
      "put": {
        "tags": [
          "Object Management"
        ],
        "summary": "Add an object in a bucket.",
        "operationId": "putObject",
        "parameters": [
          {
            "type": "string",
            "description": "Bucket name.",
            "name": "bucket",
            "in": "path",
            "required": true
          },
          {
            "description": "object content.",
            "name": "http-text",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Ok.",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "int64"
                }
              }
            }
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Not authenticated."
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Object not found."
          },
          "500": {
            "description": "Failed to put object."
          },
          "503": {
            "description": "Not available."
          }
        }
      }
    },
    "/objects/{bucket}/{objectID}": {
      "get": {
        "tags": [
          "Object Management"
        ],
        "summary": "Get object by bucket name and object id.",
        "operationId": "getObject",
        "parameters": [
          {
            "type": "string",
            "description": "Bucket name.",
            "name": "bucket",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "object id.",
            "name": "objectID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok.",
            "schema": {
              "type": "object",
              "required": [
                "object"
              ],
              "properties": {
                "object": {
                  "$ref": "#/definitions/object"
                }
              }
            }
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Not authenticated."
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Object not found."
          },
          "500": {
            "description": "Failed to list groups."
          },
          "503": {
            "description": "Not available."
          }
        }
      },
      "delete": {
        "tags": [
          "Object Management"
        ],
        "summary": "delete an object.",
        "operationId": "deleteObject",
        "parameters": [
          {
            "type": "string",
            "description": "Bucket name.",
            "name": "bucket",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "object id.",
            "name": "objectID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Not authenticated."
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Object not found."
          },
          "500": {
            "description": "Failed to delete object."
          },
          "503": {
            "description": "Not available."
          }
        }
      }
    }
  },
  "definitions": {
    "object": {
      "required": [
        "id",
        "data"
      ],
      "properties": {
        "data": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}`))
}