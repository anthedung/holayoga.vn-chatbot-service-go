package service

import (
	"vn.holayoga.dialogflow.service/model"
	"vn.holayoga.dialogflow.service/service/dao"
	"errors"
)

type Service interface{}

type YogaService struct {
	*dao.YogaCacheDao
}

func (s *YogaService) GetYogaPosesByCategory(catName string) ([]model.YogaPose, error) {
	cat, err := s.GetCategoryByName(catName)

	if err != nil {
		return nil, err
	}

	return cat.Poses, nil
}

func (s *YogaService) GetYogaPoseByName(poseName string) (*model.YogaPose, error) {
	poses := s.GetAllPoses()

	for _, p := range poses {
		if p.Name == poseName {
			return &p, nil
		}
	}
	// loop through all categories
	return nil, errors.New("pose not found: " + poseName)
}

func NewYogaService(dao *dao.YogaCacheDao) (*YogaService, error) {
	return &YogaService{
		YogaCacheDao: dao,
	}, nil
}