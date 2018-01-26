package model

import (
	"google.golang.org/appengine/datastore"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type YogaCategory struct {
	Name            string                                           // yoga co ban
	Poses           []YogaPose    `datastore:"-" `                   // list of poses in this category
	PosesStr        string        `datastore:"Poses," json:"-"`      // [PropertyLoadSaver]
	VideoCourses    []VideoCourse `datastore:"-"`                    // khoa co ban, khoa 30 ngay
	VideoCoursesStr string        `datastore:"VideoCourse" json:"-"` // [PropertyLoadSaver]
	Articles        []ArticleResource                                // blog post
	ID              int64                                            // The integer ID used in the datastore.
}

type VideoCourse struct {
	Name        string          `datastore:",omitempty"` // yoga co ban
	Videos      []VideoResource `datastore:",omitempty"` // key of videos
	Description string
}

type YogaPose struct {
	Name        string          `datastore:",omitempty"` // cơ bản (display)
	Images      []ImageResource `datastore:",omitempty"` // right,wrong, etc.
	Videos      []VideoResource `datastore:",omitempty"` // key of videos
	Description string          `datastore:",omitempty"`
}

// [PropertyLoadSaver] due to AppEngine SDK for go not supporting slices of slices
func (cat *YogaCategory) Load(ps []datastore.Property) error {
	datastore.LoadStruct(cat, ps)

	// Load []YogaPose
	var poses []YogaPose
	json.Unmarshal([]byte(cat.PosesStr), &poses)
	cat.Poses = poses

	// Load []VideoCourse
	var courses []VideoCourse
	json.Unmarshal([]byte(cat.VideoCoursesStr), &courses)
	cat.VideoCourses = courses

	return nil
}

func (cat *YogaCategory) Save() ([]datastore.Property, error) {
	logrus.Info("YogaCategory data store initial data populating ..")

	// save poses
	p, _ := json.Marshal(cat.Poses)
	cat.PosesStr = string(p)

	// save courses
	c, _ := json.Marshal(cat.VideoCourses)
	cat.VideoCoursesStr = string(c)

	return datastore.SaveStruct(cat)
}

type ImageResource struct {
	Name   string   `datastore:",omitempty"` // con mèo
	Url    string   `datastore:",omitempty"` //
	Tag    ImageTag `datastore:",omitempty"` // dung, sai goc
	Remark string   `datastore:",omitempty"` // e.g. dong tac nay tot cho blah blah => this should be in another table, dictionary
}

type VideoResource struct {
	Name         string `datastore:",omitempty"` // con mèo
	Url          string `datastore:",omitempty"` //
	Tag          string `datastore:",omitempty"` // dung, sai goc
	Remark       string `datastore:",omitempty"` // e.g. dong tac nay tot cho blah blah => this should be in another table, dictionary
	Duration     int    `datastore:",omitempty"` // duration in minute
	ThumbnailUrl string `datastore:",omitempty"` //video thumbnail
}

type ArticleResource struct {
	Name     string `datastore:",omitempty"` // con mèo
	Url      string `datastore:",omitempty"` //
	Tag      string `datastore:",omitempty"` // dung, sai goc
	Remark   string `datastore:",omitempty"` // e.g. dong tac nay tot cho blah blah => this should be in another table, dictionary
	Duration int    `datastore:",omitempty"` // duration in minute
}

func (pose *YogaPose) FindImageByTag(tag ImageTag) *ImageResource {
	for _, img := range pose.Images {
		println(img.Tag)
		if img.Tag == tag {
			return &img
		}
	}

	return nil
}

type ImageTag string

const (
	TagOriginalPoseImage ImageTag = "original" //0
	TagRightPoseImage    ImageTag = "right"
	TagWrongPoseImage    ImageTag = "wrong"
)
