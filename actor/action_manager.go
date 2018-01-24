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
	case showPosesInCategory:
		return a.ShowPosesInCategory(ctx, req)
	case showOriginalPoseImg:
		return a.ActionShowOriginalPoseImage(ctx, req)
	case showRightPoseImg:
		return a.ActionShowRightPoseImage(ctx, req)
	case showWrongPoseImg:
		return a.ActionShowWrongPoseImage(ctx, req)
	default:
		return nil, errors.New("no action found")
	}
}

func (a *ActionManager) ShowPosesInCategory(ctx context.Context, req dialogflow.WebhookRequest) (*dialogflow.WebhookResponse, error) {
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
