// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.0",
  "info": {
    "description": "This is API documentation on DataBx Sumbing1 Microservices",
    "version": "1.0.0",
    "title": "DataBx Sumbing1 Microservices",
    "contact": {
      "email": "akbar.muhammadakbarmaulana@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    }
  },
  "tags": [
    {
      "name": "utils",
      "description": "Dead simple utils that can be useful to interact with microservices"
    }
  ],
  "paths": {
    "/convert/json-to-bson": {
      "post": {
        "tags": [
          "utils"
        ],
        "summary": "Convert body json to bson encoded in byte",
        "description": "Only accept json and only return bson",
        "operationId": "ConvertJSONtoBSON",
        "security": [
          {"basicAuth":  []}
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/xClientId"
          }
        ],
        "requestBody": {
          "$ref": "#/components/requestBodies/JSON"
        },
        "responses": {
          "200": {
            "description": "ConvertJSONtoBSON success, with given request",
            "content": {
              "application/bson": {}
            }
          },
          "400": {
            "description": "Check request, body not comply with request rules",
            "content": {
              "application/bson": {}
            }
          },
          "401": {
            "description": "You are not permitted!"
          },
          "500": {
            "description": "General error, system cannot tell you what wrongs",
            "content": {
              "application/bson": {}
            }
          }
        }
      }
    }
  },
  "servers": [
    {
      "url": "http://localhost:7778/sumbing1/api/v1"
    }
  ],
  "components": {
    "requestBodies": {
      "JSON": {
        "required": true,
        "content": {
          "application/json": {
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "securitySchemes": {
      "basicAuth": {
        "type": "http",
        "scheme": "basic"
      }
    },
    "parameters": {
      "xClientId": {
        "name": "X-Client-ID",
        "in": "header",
        "description": "Custom Id that defines you as caller",
        "required": false,
        "schema": {
          "type": "string",
          "maxLength": 32,
          "pattern": "[A-Za-z0-9]{32}"
        }
      }
    }
  },
  "security": [
    {
      "basicAuth": []
    }
  ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
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