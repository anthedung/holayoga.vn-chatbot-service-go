package actor

import (
	"google.golang.org/api/dialogflow/v2beta1"
	"encoding/json"
	"context"
	svc "vn.holayoga.dialogflow.service/service"
	"vn.holayoga.dialogflow.service/model"
	"errors"
	"google.golang.org/appengine/log"
	"time"
	"math/rand"
	"vn.holayoga.dialogflow.service/utils"
	fb "vn.holayoga.dialogflow.service/model/messenger_template"
)

type ActionManager struct {
	Platform string

	// list of services
	*svc.YogaService
}

func NewActionManager(platform string, service *svc.YogaService) (*ActionManager) {
	return &ActionManager{
		Platform:    platform,
		YogaService: service,
	}
}

func (a *ActionManager) InvokeActionByName(action ActionName, ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	switch action {
	case showVideoCoursesInCategory:
		return a.ShowVideoCoursesByCategory(ctx, req)
	case showVideosByCourseNameAndCategory:
		return a.ShowVideosByCourseNameAndCategory(ctx, req)
	case showAVideoInCourseByCategory:
		return a.ShowAVideosByCourseNameAndCategory(ctx, req)
	case showPosesInCategory:
		return a.ShowPosesByCategory(ctx, req)
	case showOriginalPoseImg:
		return a.ActionShowOriginalPoseImage(ctx, req)
	case showRightPoseImg:
		return a.ActionShowRightPoseImage(ctx, req)
	case showWrongPoseImg:
		return a.ActionShowWrongPoseImage(ctx, req)
	case showRandomOptionInCategory:
		return a.ShowRandomOptionInCategory(ctx, req)
	default:
		return nil, errors.New("no action found")
	}
}

func (a *ActionManager) ShowPosesByCategory(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	rand.Seed(time.Now().Unix())
	params := make(map[DialogFlowParameter]string)
	err := json.Unmarshal(req.QueryResult.Parameters, &params)

	if err != nil {
		log.Errorf(ctx, "cannot unmarshal params ", err.Error())
		return nil, err
	}

	// get value from yoga_category param
	yogaCat := params[YogaCategoryParam]

	poses, err := a.GetYogaPosesByCategory(yogaCat)

	if err != nil {
		return nil, err
	}

	// list of quick reply variants
	quickRepTitles := []string{
		"Chọn 1 cái rồi mama chỉ kĩ cho, tốt lém đó :D",
		"Con lựa 1 động tác đi mama chỉ thêm cho :x",
		"Con thích động tác nào trong đống này?",
	}
	rep := quickRepTitles[rand.Intn(len(quickRepTitles))]

	replies := &dialogflow.IntentMessage{
		QuickReplies: &dialogflow.IntentMessageQuickReplies{
			QuickReplies: []string{},
			Title:        rep,
		},
		Platform: "FACEBOOK",
	}

	for _, p := range poses {
		replies.QuickReplies.QuickReplies = append(replies.QuickReplies.QuickReplies, "Động tác "+p.Name)
	}

	return &dialogflow.WebhookResponse{
		FulfillmentMessages: []*dialogflow.IntentMessage{replies},
	}, nil
}

func (a *ActionManager) ActionShowOriginalPoseImage(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	return a.actionShowPoseImageByTagHelper(ctx, req, model.TagOriginalPoseImage)
}

func (a *ActionManager) ActionShowRightPoseImage(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	return a.actionShowPoseImageByTagHelper(ctx, req, model.TagRightPoseImage)
}

func (a *ActionManager) ActionShowWrongPoseImage(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	return a.actionShowPoseImageByTagHelper(ctx, req, model.TagWrongPoseImage)
}

func (a *ActionManager) actionShowPoseImageByTagHelper(ctx context.Context, req dialogflow.WebhookRequest, tag model.ImageTag) (*dialogflow.WebhookResponse, error) {
	rand.Seed(time.Now().Unix())
	params := make(map[DialogFlowParameter]string)
	err := json.Unmarshal(req.QueryResult.Parameters, &params)

	if err != nil {
		log.Errorf(ctx, "cannot unmarshal params ", err.Error())
		return nil, err
	}

	// get value from yoga_category param
	poseName := params[YogaPoseParam]

	pose, err := a.GetYogaPoseByName(poseName)
	if err != nil {
		return nil, err
	}

	if pose == nil {
		log.Errorf(ctx, "no pose found for the pose name")
		return nil, errors.New("no pose found for the pose name")
	}

	img := pose.FindImageByTag(tag)

	if img == nil {
		log.Errorf(ctx, "no image found for the tag: "+string(tag)+" of "+pose.Name)
		return nil, errors.New("no image found for the tag: " + string(tag) + " of " + pose.Name)
	}

	intentImg := &dialogflow.IntentMessage{
		Image: &dialogflow.IntentMessageImage{
			ImageUri:          img.Url,
			AccessibilityText: img.Name,
		},
		Platform: a.Platform,
	}

	moreTitles := []string{
		"Tìm hiểu thêm coi?",
		"Bạn muốn biết thêm gì nào?",
		"Cần gì nữa mama tra giúp?",
		"Nữa?",
	}
	qReplies := &dialogflow.IntentMessage{
		QuickReplies: &dialogflow.IntentMessageQuickReplies{
			QuickReplies: []string{
				"Cách làm đúng",
				"Chỉ lỗi sai",
				"Thôi má, cái khác",
			},
			Title: moreTitles[rand.Intn(len(moreTitles))],
		},
		Platform: a.Platform,
	}

	return &dialogflow.WebhookResponse{
		FulfillmentMessages: []*dialogflow.IntentMessage{intentImg, qReplies},
	}, nil
}

func (a *ActionManager) ShowAVideosByCourseNameAndCategory(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	rand.Seed(time.Now().Unix())
	params := make(map[DialogFlowParameter]string)
	err := json.Unmarshal(req.QueryResult.Parameters, &params)

	if err != nil {
		log.Errorf(ctx, "cannot unmarshal params ", err.Error())
		return nil, err
	}

	// get value from yoga_category param
	yogaCat := params[YogaCategoryParam]
	courseName := params[VideoCourseNameParam]
	videoName := params[VideoNameParam]
	video, err := a.YogaService.GetAVideoByCourseNameAndCategory(ctx, videoName, courseName, yogaCat)

	if err != nil {
		return nil, err
	}

	url := &dialogflow.IntentMessage{
		Text: &dialogflow.IntentMessageText{
			Text: []string{video.Url},
		},
		Platform: "FACEBOOK",
	}
	replies := &dialogflow.IntentMessage{
		QuickReplies: &dialogflow.IntentMessageQuickReplies{
			QuickReplies: []string{"Bài tập tiếp theo", "Khoá Videos khác"},
			Title:        "muốn xem bài khác hem?",
		},
		Platform: "FACEBOOK",
	}

	return &dialogflow.WebhookResponse{
		FulfillmentMessages: []*dialogflow.IntentMessage{url, replies},
	}, nil
}

func (a *ActionManager) ShowVideosByCourseNameAndCategory(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	rand.Seed(time.Now().Unix())
	params := make(map[DialogFlowParameter]string)
	err := json.Unmarshal(req.QueryResult.Parameters, &params)

	if err != nil {
		log.Errorf(ctx, "cannot unmarshal params ", err.Error())
		return nil, err
	}

	// get value from yoga_category param
	category := params[YogaCategoryParam]
	courseName := params[VideoCourseNameParam]
	course, err := a.YogaService.GetCourseByCourseNameAndCategory(ctx, courseName, category)

	if err != nil {
		return nil, err
	}

	var templates []fb.Template
	for _, v := range course.Videos {
		templates = append(templates, fb.GenericTemplate{
			Title:    v.Name,
			Subtitle: v.Remark,
			ImageURL: v.ThumbnailUrl,
			ItemURL:  v.Url,
			Buttons: []fb.Button{
				{
					Type:  fb.ButtonTypeWebURL,
					Title: "Xem",
					URL:   v.Url,
				},
			},
		})
	}

	rawMsg, _ := fb.ConstructFacebookPayLoad(fb.AttachmentTypeTemplate, fb.Payload{
		Elements: templates,
	})

	resp := &dialogflow.WebhookResponse{
		FulfillmentMessages: []*dialogflow.IntentMessage{{
			Payload: rawMsg,
		}},
	}
	utils.DebugfPrettyPrintWithCtx(ctx, "WebhookResponse", resp)

	return resp, nil
}

func (a *ActionManager) ShowVideoCoursesByCategory(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	rand.Seed(time.Now().Unix())
	params := make(map[DialogFlowParameter]string)
	err := json.Unmarshal(req.QueryResult.Parameters, &params)

	if err != nil {
		log.Errorf(ctx, "cannot unmarshal params ", err.Error())
		return nil, err
	}

	// get value from yoga_category param
	yogaCat := params[YogaCategoryParam]

	videoCourses, err := a.GetCoursesByCategory(yogaCat)

	if err != nil {
		return nil, err
	}

	// list of quick reply variants
	quickRepTitles := []string{
		"Chọn 1 khoá coi, tốt lém đó :D",
		"Lựa 1 khoá đi mama chỉ thêm cho :x",
		"Con thích khoá nào?",
	}
	rep := quickRepTitles[rand.Intn(len(quickRepTitles))]

	replies := &dialogflow.IntentMessage{
		QuickReplies: &dialogflow.IntentMessageQuickReplies{
			QuickReplies: []string{},
			Title:        rep,
		},
		Platform: "FACEBOOK",
	}

	for _, c := range videoCourses {
		replies.QuickReplies.QuickReplies = append(replies.QuickReplies.QuickReplies, "Khoá "+c.Name)
	}

	return &dialogflow.WebhookResponse{
		FulfillmentMessages: []*dialogflow.IntentMessage{replies},
	}, nil
}

func (a *ActionManager) ShowRandomOptionInCategory(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
	rand.Seed(time.Now().Unix())

	categoryOptions := []ActionName{
		showVideoCoursesInCategory,
		showPosesInCategory,
		showArticlesInCategory,
	}

	return a.InvokeActionByName(categoryOptions[rand.Intn(len(categoryOptions))], ctx, req)
}
