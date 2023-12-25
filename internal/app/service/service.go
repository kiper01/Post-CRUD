package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kiper01/Post-CRUD/internal/app/model"
	"github.com/kiper01/Post-CRUD/internal/app/repository"
	pb "github.com/kiper01/Post-CRUD/internal/proto"
)

type PostInfoRepository interface {
	AddPostValue(ctx context.Context, value model.PostValue) error
	RemovePostValue(ctx context.Context, id string) error
	UpdatePostValue(ctx context.Context, values []model.PostValue) error
}

type PostInfoService struct {
	repo *repository.PostInfoRepository
	pb.UnimplementedPostInfoServiceServer
}

func NewPostInfoService(repo *repository.PostInfoRepository) *PostInfoService {
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

func (s *PostInfoService) RemovePostValue(ctx context.Context, req *pb.RemovePostValueRequest) (*pb.RemovePostValueResponse, error) {

	err := s.repo.RemovePostValue(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.RemovePostValueResponse{Success: true}, nil
}

func (s *PostInfoService) UpdatePostValue(ctx context.Context, req *pb.UpdatePostValueRequest) (*pb.UpdatePostValueResponse, error) {

	var postValues []model.PostValue

	for _, pv := range req.GetPost() {
		postValues = append(postValues, model.PostValue{
			ID:    pv.GetId(),
			Code:  pv.GetCode(),
			Name:  pv.GetName(),
			River: pv.GetRiver(),
		})
	}

	err := s.repo.UpdatePostValues(ctx, postValues)
	if err != nil {
		return nil, err
	}

	var updatedPostValues []*pb.PostValue
	for _, pval := range postValues {
		updatedPostValues = append(updatedPostValues, &pb.PostValue{
			Id:    pval.ID,
			Code:  pval.Code,
			Name:  pval.Name,
			River: pval.River,
		})
	}

	return &pb.UpdatePostValueResponse{Post: updatedPostValues}, nil
}

func (s *PostInfoService) GetPostValuesByPage() {

}
