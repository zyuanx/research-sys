## Research Sys

base on Gin gin-jwt GORM MongoDB-Go-Driver Swagger

```
.
├── README.md
├── config
│   ├── application.yml
│   └── rbac_model.conf
├── controllers
│   ├── req
│   │   ├── request.go
│   │   ├── role.go
│   │   └── user.go
│   ├── res
│   │   ├── response.go
│   │   ├── role.go
│   │   └── user.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── initialize
│   ├── casbin.go
│   ├── config.go
│   ├── mongo.go
│   ├── mysql.go
│   └── router.go
├── main.go
├── middlewares
│   ├── casbin.go
│   └── jwt.go
├── models
│   ├── permission.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── pkg
│   └── global
│       ├── config.go
│       ├── global.go
│       └── pagination.go
├── routers
│   ├── research.go
│   ├── role.go
│   └── user.go
├── services
│   ├── permission.go
│   ├── research.go
│   ├── role.go
│   └── user.go
└── utils
    ├── parseValidator.go
    └── s.go
```