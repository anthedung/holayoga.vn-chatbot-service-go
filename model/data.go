package model

import (
	"google.golang.org/appengine/datastore"
	"context"
)

func InitDataStore(ctx context.Context, entity string) {
	basic := &YogaCategory{
		Name:         "cơ bản",
		VideoCourses: []VideoCourses{vidCourses},
		Poses:        posesMock,
		Articles: []ArticleResource{
			{Name: "yoga cho ba bau"},
		},
	}

	key := datastore.NewIncompleteKey(ctx, entity, nil)
	key, err := datastore.Put(ctx, key, basic)
	if err != nil {
		println("err putting: " + err.Error())
	} else {
		println(key)
	}

	basic = &YogaCategory{}
	err = datastore.Get(ctx, key, basic)

	if err != nil {
		println("err getting: " + err.Error())
	} else {
		println(key)
	}
}

// [Mock]
var posesMock = []YogaPose{
	{
		Name:   "con mèo",
		Images: conMeoImgResourcesMock,
	},
	{
		Name:   "chiến binh",
		Images: []ImageResource{{Url: "http://dl.dropboxusercontent.com/s/r9o75wloy6ng3jg/HolaYogaLogo_full.png"}},
	},
}

var conMeoImgResourcesMock = []ImageResource{
	{
		Url:  "https://thehealthorange.com/wp-content/uploads/2016/09/Resized-79F7Q-2-1.jpg",
		Name: "con-meo-dung",
		Tag:  TagRightPoseImage,
	},
	{
		Name: "con-meo",
		Url:  "http://evenmoreaboutyoga.com/wp-content/uploads/2015/12/cat-pose-620x350.jpg",
		Tag:  TagOriginalPoseImage,
	},
	{
		Name: "con-meo-sai",
		Url:  "https://thumb7.shutterstock.com/display_pic_with_logo/4030810/385833517/stock-vector-cartoon-dog-shows-yoga-pose-adho-mukha-svanasana-downward-facing-dog-surya-namaskara-san-385833517.jpg",
		Tag:  TagWrongPoseImage,
	},
}

var vid = VideoResource{
	Name:   "Bài 1",
	Remark: "Yoga cơ bản bài 1",
	Url:    "https://www.youtube.com/watch?v=7KBXzCp_U8I",
}
// videos courses
var vidCourses = VideoCourses{
	Name:   "10 ngày cơ bản",
	Videos: []VideoResource{vid},
}
