package messenger_template

import (
	"google.golang.org/api/dialogflow/v2beta1"
	"google.golang.org/api/gensupport"
)

type ExtendedIntentMessage struct {
	dialogflow.IntentMessage

	Type string
}

func (s *ExtendedIntentMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ExtendedIntentMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}
