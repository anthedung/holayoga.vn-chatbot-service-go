package main

import (
	"github.com/sirupsen/logrus"
	"flag"
	"google.golang.org/appengine"
	"net/http"
	"vn.holayoga.dialogflow.service/model"
)

var (
	entity    = flag.String("entity", "Category", "Entity to populate")
)

func init() {
	flag.Parse()
}

func main() {
	logrus.Info("data store initial data populating..")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		model.InitDataStore(appengine.NewContext(r), *entity)
	})

	appengine.Main()

	logrus.Info("data store inited")
}