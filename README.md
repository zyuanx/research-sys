## Research sys

### About The Project
> Dynamic form, survey questionnaire backend system. Front-end repo: [vue-research-admin](https://github.com/Pandalzy/vue-research-admin)

base on Gin | GORM | gin-jwt | Casbin | mongo-driver | Swagger | Zap | Viper

```
.
├── config
│   ├── application-pro.yml
│   ├── application.yml
│   └── rbac_model.conf
├── controllers
│   ├── req
│   │   ├── permission.go
│   │   ├── record.go
│   │   ├── request.go
│   │   ├── research.go
│   │   ├── role.go
│   │   └── user.go
│   ├── res
│   │   ├── response.go
│   │   ├── role.go
│   │   └── user.go
│   ├── permission.go
│   ├── record.go
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
│   ├── redis.go
│   ├── router.go
│   ├── viper.go
│   └── zap.go
├── logs
│   ├── nginx
│   │   ├── access.log
│   │   └── error.log
│   └── zap
│       ├── 2021-05-12.log
├── middlewares
│   ├── casbin.go
│   ├── cors.go
│   └── jwt.go
├── models
│   ├── base.go
│   ├── permission.go
│   ├── record.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── nginx
│   ├── ssl
│   ├── Dockerfile
│   ├── nginx.conf
│   └── nginx_bak.conf
├── pkg
│   ├── global
│   │   ├── global.go
│   │   └── pagination.go
│   └── utils
│       ├── json.go
│       ├── parseValidator.go
│       └── s.go
├── routers
│   ├── permission.go
│   ├── record.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── services
│   ├── casbin.go
│   ├── permission.go
│   ├── record.go
│   ├── research.go
│   ├── role.go
│   └── user.go
├── Dockerfile
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```