package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kiper01/Post-CRUD/internal/app/model"
	"github.com/kiper01/Post-CRUD/internal/app/repository"
)

type PostInfoRepository interface {
	AddPostValue(ctx context.Context, value model.PostValue) error
}

type PostInfoService struct {
	repo *repository.PostInfoRepository
	pb.UnimplementedPostInfoServiceServer
}

func NewPostStatsService(repo *repository.PostInfoRepository) *PostInfoService {
	return &PostInfoService{repo: repo}
}

func (s *PostInfoService) AddPostValue(ctx context.Context, req *pb.AddPostValueRequest) (*pb.AddPostValueResponse, error) {

	postValue := model.PostValue{
		ID:    uuid.New().String(),
		Code:  req.GetCode(),
		Name:  req.GetName(),
		River: req.GetRiver(),
	}

	if err := s.repo.AddPostValue(ctx, postValue); err != nil {
		return nil, err
	}

	return &pb.AddPostValueResponse{
		PostValue: &pb.PostValue{
			Id:    postValue.ID,
			Code:  postValue.Code,
			Name:  postValue.Name,
			River: postValue.River,
		},
	}, nil
}
