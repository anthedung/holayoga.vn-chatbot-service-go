package actor

import (
	"google.golang.org/api/dialogflow/v2beta1"
	"encoding/json"
)

func ActionShowPosesInCategory(req dialogflow.WebhookRequest) dialogflow.WebhookResponse {
	params := map[DialogFlowParameter]YogaCategory{}
	json.Unmarshal(req.QueryResult.Parameters, params)

	// get value from yoga_category param
	yogaCat := params[YogaCategoryParam]

	poses := getYogaPosesByCategory(yogaCat)

	replies := &dialogflow.IntentMessage{
		QuickReplies: &dialogflow.IntentMessageQuickReplies{
			QuickReplies: []string{},
		},
	}

	for _, p := range poses {
		replies.QuickReplies.QuickReplies = append(replies.QuickReplies.QuickReplies, p.Name)
	}

	return dialogflow.WebhookResponse{
		FulfillmentMessages: []*dialogflow.IntentMessage{replies},
	}
}
