package service

import (
	"context"
	"math"

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

	for _, pv := range req.GetPostValue() {
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

	return &pb.UpdatePostValueResponse{PostValue: updatedPostValues}, nil
}

func (s *PostInfoService) GetPostValues(ctx context.Context, req *pb.GetPostValuesRequest) (*pb.GetPostValuesResponse, error) {

	pageSize := 50

	page := int(req.GetPage())
	if page < 1 {
		page = 1
	}

	postValues, totalCount, err := s.repo.GetPostValues(ctx, int(req.GetCode()), page, pageSize)
	if err != nil {
		return nil, err
	}

	var pbPostValues []*pb.PostValue

	for _, pv := range postValues {
		pbPostValues = append(pbPostValues, &pb.PostValue{
			Id:    pv.ID,
			Code:  pv.Code,
			Name:  pv.Name,
			River: pv.River,
		})
	}

	maxPage := uint32(math.Ceil(float64(totalCount) / float64(pageSize)))

	return &pb.GetPostValuesResponse{
		Page:      uint32(page),
		MaxPage:   maxPage,
		PostValue: pbPostValues,
	}, nil
}

func (s *PostInfoService) GetPostValuesByCodeOrId(ctx context.Context, req *pb.GetPostValuesByCodeOrIdRequest) (*pb.GetPostValuesByCodeOrIdResponse, error) {

	postValues, err := s.repo.GetPostValuesByCodeOrId(ctx, req.GetId(), int(req.GetCode()))
	if err != nil {
		return nil, err
	}

	var pbPostValues []*pb.PostValue

	for _, pv := range postValues {
		pbPostValues = append(pbPostValues, &pb.PostValue{
			Id:    pv.ID,
			Code:  pv.Code,
			Name:  pv.Name,
			River: pv.River,
		})
	}

	return &pb.GetPostValuesByCodeOrIdResponse{
		PostValue: pbPostValues,
	}, nil
}
