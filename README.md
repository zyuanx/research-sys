## Research sys

### About The Project
> Dynamic form, survey questionnaire backend system. Front-end repo: [vue-research-admin](https://github.com/Pandalzy/vue-research-admin)

base on Gin | GORM | gin-jwt | Casbin | mongo-driver | Swagger | Zap | Viper

```
.
├── config
│   ├── application.yml
│   └── rbac_model.conf
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal
│   ├── conf
│   │   └── global.go
│   ├── controller
│   │   ├── req
│   │   │   ├── permission.go
│   │   │   ├── record.go
│   │   │   ├── request.go
│   │   │   ├── research.go
│   │   │   ├── role.go
│   │   │   └── user.go
│   │   ├── res
│   │   │   ├── response.go
│   │   │   ├── role.go
│   │   │   └── user.go
│   │   ├── permission.go
│   │   ├── record.go
│   │   ├── research.go
│   │   ├── role.go
│   │   └── user.go
│   ├── initialize
│   │   ├── casbin.go
│   │   ├── config.go
│   │   ├── mongo.go
│   │   ├── mysql.go
│   │   ├── redis.go
│   │   ├── router.go
│   │   ├── viper.go
│   │   └── zap.go
│   ├── middleware
│   │   ├── casbin.go
│   │   ├── cors.go
│   │   └── jwt.go
│   ├── model
│   │   ├── base.go
│   │   ├── permission.go
│   │   ├── record.go
│   │   ├── research.go
│   │   ├── role.go
│   │   └── user.go
│   ├── router
│   │   ├── permission.go
│   │   ├── record.go
│   │   ├── research.go
│   │   ├── role.go
│   │   └── user.go
│   ├── service
│   │   ├── casbin.go
│   │   ├── permission.go
│   │   ├── record.go
│   │   ├── research.go
│   │   ├── role.go
│   │   └── user.go
│   ├── util
│   │   ├── pagination.go
│   │   └── response.go
│   └── app.go
├── nginx
│   ├── Dockerfile
│   ├── nginx.conf
│   └── nginx_bak.conf
├── pkg
│   └── utils
│       ├── json.go
│       ├── parseValidator.go
│       └── s.go
├── Dockerfile
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go

```