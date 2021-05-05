## Research sys

### About The Project
> Dynamic form, survey questionnaire backend system. Front-end repo: [vue-research-admin](https://github.com/Pandalzy/vue-research-admin)

base on Gin | GORM | gin-jwt | Casbin | mongo-driver | Swagger | Zap | Viper

```
.
├── config
│   ├── application.yml
│   └── rbac_model.conf
├── controllers
│   ├── permission.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
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
├── middlewares
│   ├── casbin.go
│   ├── cors.go
│   └── jwt.go
├── models
│   ├── base.go
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
│   ├── permission.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── services
│   ├── permission.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── README.md
├── go.mod
├── go.sum
└── main.go
```