package controller

import "github.com/zyuanx/research-sys/internal/service"

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}
