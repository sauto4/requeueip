// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 The RequeueIP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "RequeueIP API",
    "license": {
      "name": "Apache 2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "v1"
  },
  "basePath": "/v1",
  "paths": {
    "/ipam/healthz": {
      "get": {
        "tags": [
          "ipam"
        ],
        "responses": {
          "200": {
            "description": "Success"
          }
        }
      }
    },
    "/ipam/ips": {
      "post": {
        "tags": [
          "ipam"
        ],
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CmdAddArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/CmdAddResult"
            }
          },
          "500": {
            "description": "Failed to allocate IP addresses",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "ipam"
        ],
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CmdDelArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed to release IP addresses",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CmdAddArgs": {
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "podNamespace",
        "podName"
      ],
      "properties": {
        "containerID": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        },
        "ipv4Pools": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ipv6Pools": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        }
      }
    },
    "CmdAddResult": {
      "type": "object",
      "required": [
        "ips"
      ],
      "properties": {
        "dns": {
          "type": "object",
          "$ref": "#/definitions/DNS"
        },
        "ips": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/IPConfig"
          }
        },
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Route"
          }
        }
      }
    },
    "CmdDelArgs": {
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "podNamespace",
        "podName"
      ],
      "properties": {
        "containerID": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        }
      }
    },
    "DNS": {
      "type": "object",
      "required": [
        "nameservers"
      ],
      "properties": {
        "domain": {
          "type": "string"
        },
        "nameservers": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "search": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Error": {
      "type": "string"
    },
    "IPConfig": {
      "type": "object",
      "required": [
        "address"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "gateway": {
          "type": "string"
        }
      }
    },
    "Route": {
      "type": "object",
      "required": [
        "dst"
      ],
      "properties": {
        "dst": {
          "type": "string"
        },
        "gw": {
          "type": "string"
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "RequeueIP API",
    "license": {
      "name": "Apache 2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "v1"
  },
  "basePath": "/v1",
  "paths": {
    "/ipam/healthz": {
      "get": {
        "tags": [
          "ipam"
        ],
        "responses": {
          "200": {
            "description": "Success"
          }
        }
      }
    },
    "/ipam/ips": {
      "post": {
        "tags": [
          "ipam"
        ],
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CmdAddArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/CmdAddResult"
            }
          },
          "500": {
            "description": "Failed to allocate IP addresses",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "ipam"
        ],
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CmdDelArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed to release IP addresses",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CmdAddArgs": {
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "podNamespace",
        "podName"
      ],
      "properties": {
        "containerID": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        },
        "ipv4Pools": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ipv6Pools": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        }
      }
    },
    "CmdAddResult": {
      "type": "object",
      "required": [
        "ips"
      ],
      "properties": {
        "dns": {
          "type": "object",
          "$ref": "#/definitions/DNS"
        },
        "ips": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/IPConfig"
          }
        },
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Route"
          }
        }
      }
    },
    "CmdDelArgs": {
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "podNamespace",
        "podName"
      ],
      "properties": {
        "containerID": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        }
      }
    },
    "DNS": {
      "type": "object",
      "required": [
        "nameservers"
      ],
      "properties": {
        "domain": {
          "type": "string"
        },
        "nameservers": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "search": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Error": {
      "type": "string"
    },
    "IPConfig": {
      "type": "object",
      "required": [
        "address"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "gateway": {
          "type": "string"
        }
      }
    },
    "Route": {
      "type": "object",
      "required": [
        "dst"
      ],
      "properties": {
        "dst": {
          "type": "string"
        },
        "gw": {
          "type": "string"
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
}
