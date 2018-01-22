package actor

type YogaCategory string

const CatCoBan YogaCategory = "co-ban"

type YogaPose struct {
	Name   string
	Images []ImageResource
}

type ImageResource struct {
	Id   int    //
	Name string // con-meo
	Url  string //
}

func getYogaPosesByCategory(cat YogaCategory) []YogaPose {
	return categoryMap[cat]
}


