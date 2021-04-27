## Research Sys

base on Gin+gin-jwt+GORM+MongoDB Go Driver+Swagger

```
.
├── config
│   └── application.yml
├── controllers
│   ├── request
│   │   ├── role.go
│   │   └── user.go
│   ├── research.go
│   ├── response
│   │   ├── response.go
│   │   └── user.go
│   ├── role.go
│   └── user.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── initialize
│   ├── config.go
│   ├── mongo.go
│   ├── mysql.go
│   └── router.go
├── main.go
├── middlewares
│   └── jwt.go
├── models
│   ├── research.go
│   ├── role.go
│   └── user.go
├── pkg
│   └── global
│       ├── config.go
│       └── global.go
├── routers
│   ├── research.go
│   ├── role.go
│   └── user.go
├── services
│   ├── research.go
│   ├── role.go
│   └── user.go
└── utils
    ├── parseValidator.go
    └── s.go

```