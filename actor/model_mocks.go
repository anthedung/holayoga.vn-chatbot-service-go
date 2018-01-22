package actor

// [Mock]
var categoryMap = map[YogaCategory][]YogaPose{
	CatCoBan: posesMock,
}

var posesMock = []YogaPose{
	{
		Name:   "con-meo",
		Images: []ImageResource{{Url: "http://dl.dropboxusercontent.com/s/r9o75wloy6ng3jg/HolaYogaLogo_full.png"}},
	},
}
