// Code generated by github.com/reearth/reearth-backend/tools/cmd/embed, DO NOT EDIT.

package manifest

const SchemaJSON string = `{
  "$id": "https://app.reearth.io/schemas/plugin-manifest",
  "$schema": "http://json-schema.org/draft-04/schema",
  "description": "Re:Earth plugin manifest schema",
  "definitions": {
    "id": {
      "$id": "#id",
      "type": "string",
      "pattern": "^[A-Za-z]{1}[\\w-:.]{0,}$"
    },
    "id?": {
      "$id": "#id?",
      "type": [
        "string",
        "null"
      ],
      "pattern": "^[A-Za-z]{1}[\\w-:.]{0,}$"
    },
    "valuetype": {
      "$id": "#valuetype",
      "type": "string",
      "enum": [
        "bool",
        "number",
        "string",
        "url",
        "latlng",
        "latlngheight",
        "camera",
        "typography",
        "coordinates",
        "polygon",
        "rect",
        "ref"
      ]
    },
    "propertyPointer": {
      "$id": "#propertyPointer",
      "type": [
        "object",
        "null"
      ],
      "properties": {
        "schemaGroupId": {
          "type": "string"
        },
        "fieldId": {
          "type": "string"
        }
      },
      "required": [
        "schemaGroupId",
        "fieldId"
      ],
      "additionalProperties": false
    },
    "propertyLinkableFields": {
      "$id": "#propertyLinkableFields",
      "type": [
        "object",
        "null"
      ],
      "properties": {
        "latlng": {
          "$ref": "#/definitions/propertyPointer"
        },
        "url": {
          "$ref": "#/definitions/propertyPointer"
        }
      },
      "additionalProperties": false
    },
    "propertyCondition": {
      "$id": "#propertyCondition",
      "type": [
        "object",
        "null"
      ],
      "properties": {
        "field": {
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/valuetype"
        },
        "value": {}
      },
      "required": [
        "field",
        "type",
        "value"
      ],
      "additionalProperties": false
    },
    "propertySchemaField": {
      "$id": "#propertySchemaField",
      "type": "object",
      "properties": {
        "id": {
          "$ref": "#/definitions/id"
        },
        "title": {
          "type": [
            "string",
            "null"
          ]
        },
        "description": {
          "type": [
            "string",
            "null"
          ]
        },
        "type": {
          "$ref": "#/definitions/valuetype"
        },
        "prefix": {
          "type": [
            "string",
            "null"
          ]
        },
        "suffix": {
          "type": [
            "string",
            "null"
          ]
        },
        "defaultValue": {},
        "ui": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "layer",
            "color",
            "multiline",
            "selection",
            "buttons",
            "range",
            "image",
            "video",
            "file",
            "camera_pose"
          ]
        },
        "min": {
          "type": [
            "number",
            "null"
          ]
        },
        "max": {
          "type": [
            "number",
            "null"
          ]
        },
        "choices": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "key": {
                "type": "string"
              },
              "label": {
                "type": "string"
              },
              "icon": {
                "type": "string"
              }
            },
            "required": [
              "key"
            ],
            "additionalProperties": false
          }
        },
        "availableIf": {
          "$ref": "#/definitions/propertyCondition"
        }
      },
      "required": [
        "id",
        "type",
        "title"
      ],
      "additionalProperties": false
    },
    "propertySchemaGroup": {
      "$id": "#propertySchemaGroup",
      "type": "object",
      "properties": {
        "id": {
          "$ref": "#/definitions/id"
        },
        "title": {
          "type": "string"
        },
        "description": {
          "type": [
            "string",
            "null"
          ]
        },
        "list": {
          "type": "boolean"
        },
        "availableIf": {
          "$ref": "#/definitions/propertyCondition"
        },
        "representativeField": {
          "$ref": "#/definitions/id?"
        },
        "fields": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/propertySchemaField"
          }
        }
      },
      "required": [
        "id",
        "title"
      ],
      "additionalProperties": false
    },
    "propertySchema": {
      "$id": "#propertySchema",
      "type": [
        "object",
        "null"
      ],
      "properties": {
        "version": {
          "type": "number"
        },
        "linkable": {
          "$ref": "#/definitions/propertyLinkableFields"
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/propertySchemaGroup"
          }
        }
      },
      "additionalProperties": false
    },
    "location": {
      "$id": "#location",
      "type": [
        "object",
        "null"
      ],
      "properties": {
        "zone": {
          "type": "string",
          "enum": [
            "inner",
            "outer"
          ]
        },
        "section": {
          "type": "string",
          "enum": [
            "left",
            "center",
            "right"
          ]
        },
        "area": {
          "type": "string",
          "enum": [
            "top",
            "middle",
            "bottom"
          ]
        }
      }
    },
    "extension": {
      "$id": "#extension",
      "type": "object",
      "properties": {
        "id": {
          "$ref": "#/definitions/id"
        },
        "title": {
          "type": "string"
        },
        "description": {
          "type": [
            "string",
            "null"
          ]
        },
        "icon": {
          "type": [
            "string",
            "null"
          ]
        },
        "visualizer": {
          "type": "string",
          "enum": [
            "cesium"
          ]
        },
        "type": {
          "type": "string",
          "enum": [
            "primitive",
            "widget",
            "block",
            "visualizer",
            "infobox"
          ]
        },
        "widgetLayout": {
          "type": [
            "object",
            "null"
          ],
          "properties": {
            "extendable": {
              "type": [
                "boolean",
                "null"
              ]
            },
            "extended": {
              "type": "boolean"
            },
            "defaultLocation": {
              "$ref": "#/definitions/location"
            }
          }
        },
        "schema": {
          "$ref": "#/definitions/propertySchema"
        }
      },
      "required": [
        "id",
        "title",
        "visualizer",
        "type"
      ],
      "additionalProperties": false
    },
    "root": {
      "$id": "#root",
      "type": "object",
      "properties": {
        "id": {
          "$ref": "#/definitions/id"
        },
        "title": {
          "type": "string"
        },
        "system": {
          "type": "boolean"
        },
        "version": {
          "type": "string"
        },
        "description": {
          "type": [
            "string",
            "null"
          ]
        },
        "repository": {
          "type": [
            "string",
            "null"
          ]
        },
        "author": {
          "type": [
            "string",
            "null"
          ]
        },
        "main": {
          "type": [
            "string",
            "null"
          ]
        },
        "extensions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/extension"
          }
        },
        "schema": {
          "$ref": "#/definitions/propertySchema"
        }
      },
      "required": [
        "id",
        "title"
      ],
      "additionalProperties": false
    }
  },
  "$ref": "#/definitions/root"
}`
