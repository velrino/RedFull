define({ "api": [
  {
    "type": "post",
    "url": "/login",
    "title": "Login",
    "name": "Login",
    "group": "Auth",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Content-Type",
            "description": "<p><code>application/json</code> JSON representation</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>E-mail of user</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>Password default <code>12345</code></p>"
          }
        ]
      }
    },
    "sampleRequest": [
      {
        "url": "/login"
      }
    ],
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Auth"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "C__Users_velrino_go_src_github_com_velrino_RedFull_doc_main_js",
    "groupTitle": "C__Users_velrino_go_src_github_com_velrino_RedFull_doc_main_js",
    "name": ""
  },
  {
    "type": "get",
    "url": "/api/users/:id",
    "title": "Get",
    "name": "Get_Users",
    "group": "Users",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the User.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>E-mail of the User.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "gravatar",
            "description": "<p>Gravatar of the User.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\t\"id\": 1,\n\t\"email\": \"colin@redventures.com\",\n\t\"gravatar\": \"http://www.Gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterID\",\n\t\"name\": \"Colin\"\n}",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/api/users/1"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p><code>Bearer {TOKEN}</code> Bearer token for authentication.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UserNotFound",
            "description": "<p>The Users was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"User Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Users"
  },
  {
    "type": "get",
    "url": "/api/users",
    "title": "List",
    "name": "List_Users",
    "group": "Users",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the User.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>E-mail of the User.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "gravatar",
            "description": "<p>Gravatar of the User.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n[\n\t{\n\t\t\"id\": 1,\n\t\t\"email\": \"colin@redventures.com\",\n\t\t\"gravatar\": \"http://www.Gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterID\",\n\t\t\"name\": \"Colin\"\n\t},\n\t{\n\t\t\"id\": 2,\n\t\t\"email\": \"kyle@redventures.com\",\n\t\t\"gravatar\": \"http://www.Gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterID\",\n\t\t\"name\": \"Kyle\"\n\t},\n]",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/api/users"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p><code>Bearer {TOKEN}</code> Bearer token for authentication.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UserNotFound",
            "description": "<p>The Users was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Users Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Users"
  },
  {
    "type": "post",
    "url": "/api/widgets",
    "title": "Create",
    "name": "Create_Widget",
    "group": "Widgets",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "Name",
            "description": "<p>Name of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "Inventory",
            "description": "<p>Inventory of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "Color",
            "description": "<p>Color of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "Price",
            "description": "<p>Price of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "Boolean",
            "optional": false,
            "field": "Melts",
            "description": "<p>Melts of the Widget.</p>"
          }
        ]
      }
    },
    "sampleRequest": [
      {
        "url": "/api/widgets"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p><code>Bearer {TOKEN}</code> Bearer token for authentication.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\t\"message\": \"Widget created successfully!\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "WidgetBadRequest",
            "description": "<p>The Widget bad request.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 400 Bad Request\n{\n\t\"message\": \"Incorrect param\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Widgets"
  },
  {
    "type": "get",
    "url": "/api/widgets/:id",
    "title": "Get",
    "name": "Get_Widget",
    "group": "Widgets",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>Id of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "inventory",
            "description": "<p>Inventory of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "color",
            "description": "<p>Color of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "price",
            "description": "<p>Price of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "melts",
            "description": "<p>Melts of the Widget.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\t\"id\": 1,\n\t\"color\": \"blue\",\n\t\"inventory\": 51,\n\t\"name \": \"Losenoid\",\n\t\"price\": \"9.99\"\n}",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/api/widgets/1"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p><code>Bearer {TOKEN}</code> Bearer token for authentication.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "WidgetNotFound",
            "description": "<p>The Widget was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Widget Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Widgets"
  },
  {
    "type": "get",
    "url": "/api/widgets/",
    "title": "List",
    "name": "List_Widget",
    "group": "Widgets",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>Id of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "inventory",
            "description": "<p>Inventory of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "color",
            "description": "<p>Color of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "price",
            "description": "<p>Price of the Widget.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "melts",
            "description": "<p>Melts of the Widget.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n[\n    {\n        \"id\": 1,\n        \"color\": \"blue\",\n        \"inventory\": 51,\n        \"name \": \"Losenoid\",\n        \"price\": \"9.99\"\n    },\n    {\n        \"id\": 2,\n        \"color\": \"red\",\n        \"inventory\": 7,\n        \"name \": \"Rowlow\",\n        \"price\": \"4.00\"\n    }\n]",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "/api/widgets"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p><code>Bearer {TOKEN}</code> Bearer token for authentication.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "WidgetsNotFound",
            "description": "<p>The Widgets was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"message\": \"Widgets Not Found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Widgets"
  },
  {
    "type": "put",
    "url": "/api/widgets/:id",
    "title": "Update",
    "name": "Update_Widget",
    "group": "Widgets",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "inventory",
            "description": "<p>Inventory of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "color",
            "description": "<p>Color of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "price",
            "description": "<p>Price of the Widget.</p>"
          },
          {
            "group": "Parameter",
            "type": "Boolean",
            "optional": false,
            "field": "melts",
            "description": "<p>Melts of the Widget.</p>"
          }
        ]
      }
    },
    "sampleRequest": [
      {
        "url": "/api/widgets/1"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p><code>Bearer {TOKEN}</code> Bearer token for authentication.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\t\"message\": \"Widget updated successfully!\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "WidgetNotFound",
            "description": "<p>The Widget was not found.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "WidgetBadRequest",
            "description": "<p>The Widget bad request.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n\t\"message\": \"Widget updated successfully!\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 400 Bad Request\n{\n\t\"message\": \"Incorrect param\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./doc.go",
    "groupTitle": "Widgets"
  }
] });
