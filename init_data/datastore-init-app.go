package main

import (
	"github.com/sirupsen/logrus"
	"flag"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"context"
	"vn.holayoga.dialogflow.service/model"
	"net/http"
)

var (
	projectId = flag.String("project-id", "newagent-4790c", "Project Id to populate store")
	entity    = flag.String("entity", "Category", "Entity to populate")
)

func init() {
	flag.Parse()
}

func main() {
	logrus.Info("data store initial data populating..")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		InitDataStore(*entity, appengine.NewContext(r))
	})

	appengine.Main()

	logrus.Info("data store inited")
}

func InitDataStore(entity string, ctx context.Context) {
	// Print welcome message.

	basic := &model.YogaCategory{
		Name:         "cơ bản",
		VideoCourses: []model.VideoCourses{vidCourses},
		Poses:        posesMock,
		Articles: []model.ArticleResource{
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

	basic = &model.YogaCategory{}
	err = datastore.Get(ctx, key, basic)

	if err != nil {
		println("err getting: " + err.Error())
	} else {
		println(key)
	}
}

// [Mock]
var posesMock = []model.YogaPose{
	{
		Name:   "con mèo",
		Images: conMeoImgResourcesMock,
	},
	{
		Name:   "chiến binh",
		Images: []model.ImageResource{{Url: "http://dl.dropboxusercontent.com/s/r9o75wloy6ng3jg/HolaYogaLogo_full.png"}},
	},
}

var conMeoImgResourcesMock = []model.ImageResource{
	{
		Url:  "https://thehealthorange.com/wp-content/uploads/2016/09/Resized-79F7Q-2-1.jpg",
		Name: "con-meo-dung",
		Tag:  model.TagRightPoseImage,
	},
	{
		Name: "con-meo",
		Url:  "http://evenmoreaboutyoga.com/wp-content/uploads/2015/12/cat-pose-620x350.jpg",
		Tag:  model.TagOriginalPoseImage,
	},
	{
		Name: "con-meo-sai",
		Url:  "https://thumb7.shutterstock.com/display_pic_with_logo/4030810/385833517/stock-vector-cartoon-dog-shows-yoga-pose-adho-mukha-svanasana-downward-facing-dog-surya-namaskara-san-385833517.jpg",
		Tag:  model.TagWrongPoseImage,
	},
}

var vid = model.VideoResource{
	Name:   "Bài 1",
	Remark: "Yoga cơ bản bài 1",
	Url:    "https://www.youtube.com/watch?v=7KBXzCp_U8I",
}
// videos courses
var vidCourses = model.VideoCourses{
	Name:   "10 ngày cơ bản",
	Videos: []model.VideoResource{vid},
}
