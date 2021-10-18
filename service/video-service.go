package service

import "go-gin/entity"

type VideoService interface {
	Save(entitiy.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	video []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(entitiy.Video) entity.Video {

}

func (service *videoService) FindAll() []entity.Video {

}
