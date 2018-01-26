package messenger_template

import (
	"encoding/json"
	"errors"
)

type TemplateType string
type TopElementStyle string

type Template interface {
	Type() TemplateType
	//TopElementStyle() TopElementStyle
	SupportsButtons() bool
}

type Payload struct {
	Elements []Template `json:"elements"`
}

type rawPayload struct {
	Type            TemplateType    `json:"template_type"`
	Elements        []Template      `json:"elements"`
	//TopElementStyle TopElementStyle `json:"top_element_style"`
}

func (p *Payload) MarshalJSON() ([]byte, error) {
	rp := &rawPayload{}
	if len(p.Elements) < 1 {
		return []byte{}, errors.New("Elements slice cannot be empty")
	}
	rp.Elements = p.Elements
	rp.Type = p.Elements[0].Type()
	return json.Marshal(rp)
}
