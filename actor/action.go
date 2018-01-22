package actor

import (
	"google.golang.org/api/dialogflow/v2beta1"
)

type Action string

const showExercisesInCategory Action = "show_exercises_in_category"

type ActionFunc func(dialogflow.WebhookRequest) dialogflow.WebhookResponse

var actionMap = map[Action]ActionFunc{
	showExercisesInCategory: ActionShowPosesInCategory,
}

func GetActorByAction(action Action) ActionFunc {
	return actionMap[action]
}
