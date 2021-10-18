package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/entity"
	"go-gin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}
