## Research Sys

base on Gin | GORM | gin-jwt | Casbin | MongoDB-Go-Driver | Swagger | Zap | Viper

```
.
├── README.md
├── config
│   ├── application.yml
│   └── rbac_model.conf
├── controllers
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
│   ├── router.go
│   ├── viper.go
│   └── zap.go
├── logs
│   └── 2021-05-03.log
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
│   ├── global
│   │   ├── config.go
│   │   ├── global.go
│   │   └── pagination.go
│   ├── log
│   │   ├── constant.go
│   │   └── log.go
│   ├── req
│   │   ├── request.go
│   │   ├── role.go
│   │   └── user.go
│   ├── res
│   │   ├── response.go
│   │   ├── role.go
│   │   └── user.go
│   └── utils
│       ├── json.go
│       ├── parseValidator.go
│       └── s.go
├── routers
│   ├── research.go
│   ├── role.go
│   └── user.go
└── services
    ├── permission.go
    ├── research.go
    ├── role.go
    └── user.go
```