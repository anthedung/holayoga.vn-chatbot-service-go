package service

import (
	"vn.holayoga.dialogflow.service/model"
	"vn.holayoga.dialogflow.service/service/dao"
	"errors"
	"golang.org/x/net/context"
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

func (s *YogaService) GetCoursesByCategory(catName string) ([]model.VideoCourse, error) {
	cat, err := s.GetCategoryByName(catName)

	if err != nil {
		return nil, err
	}

	return cat.VideoCourses, nil
}

func (s *YogaService) GetCourseByCourseNameAndCategory(ctx context.Context, courseName string, catName string) (*model.VideoCourse, error) {
	cat, err := s.GetCategoryByName(catName)

	if err != nil {
		return nil, err
	}

	var course model.VideoCourse
	for _, c := range cat.VideoCourses {
		if c.Name == courseName {
			course = c
		}
	}

	return &course, nil
}

func (s *YogaService) GetAVideoByCourseNameAndCategory(ctx context.Context, videoName string, courseName string, catName string) (*model.VideoResource, error) {
	cat, err := s.GetCourseByCourseNameAndCategory(ctx, courseName, catName)

	if err != nil {
		return nil, err
	}

	var video model.VideoResource
	for _, v := range cat.Videos {
		if v.Name == videoName {
			video = v
		}
	}

	return &video, nil
}

func NewYogaService(dao *dao.YogaCacheDao) (*YogaService, error) {
	return &YogaService{
		YogaCacheDao: dao,
	}, nil
}
