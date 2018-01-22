package actor

/*
import "google.golang.org/api/dialogflow/v2beta1"

type ImageManager interface {
	// get url based on image name and category
	GetImageUrlByName(imgRes ImageResource) string
	GetImageResourcesByCategory(cat string) []ImageResource
}



type GoogleCloudImageManager struct{}

// get url based on image name and category
func (i *GoogleCloudImageManager) GetImageUrlByName(imgRes ImageResource) string {
	return "http://dl.dropboxusercontent.com/s/r9o75wloy6ng3jg/HolaYogaLogo_full.png"
}

// get Images based on category e.g. yoga-co-ban
func (i *GoogleCloudImageManager) GetImageResourcesByCategory(cat string) []ImageResource {
	return []ImageResource{
		{Url: "http://dl.dropboxusercontent.com/s/r9o75wloy6ng3jg/HolaYogaLogo_full.png"},
		{Url: "http://dl.dropboxusercontent.com/s/r9o75wloy6ng3jg/HolaYogaLogo_full.png"},
	}
}

type ImageServiceProvider struct {
}

func (i *ImageServiceProvider) Respond(request dialogflow.WebhookRequest) dialogflow.WebhookResponse {


	return dialogflow.WebhookResponse{
		FulfillmentText:"tam thoi the nay da",
	}
}

*/
