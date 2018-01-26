package messenger_template

import "encoding/json"

type FacebookPayload struct {
	Facebook Facebook `json:"facebook,omitempty"`
}

type Facebook struct {
	Attachment Attachment `json:"attachment,omitempty"`
}

func ConstructFacebookPayLoad(t AttachmentType, payload Payload) ([]byte, error) {
	fb := FacebookPayload{
		Facebook: Facebook{
			Attachment: Attachment{
				Type: AttachmentTypeTemplate,
				Payload: payload,
			},
		},
	}

	return json.Marshal(fb)
}