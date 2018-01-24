package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
	hg "vn.holayoga.dialogflow.service"
	"vn.holayoga.dialogflow.service/actor"
	"google.golang.org/appengine"
	"vn.holayoga.dialogflow.service/service/dao"
	"vn.holayoga.dialogflow.service/service"
	"flag"
)

var (
	projectId = flag.String("project-id", "newagent-4790c", "Project Id to populate store")
)

func init() {
	flag.Parse()
}

func main() {
	router := mux.NewRouter()

	// Init DAO for datastore without context
	// context to be set
	dao, err := dao.NewYogaCategoryCache(12, *projectId, dao.CategoryDataStoreEntity)
	if err != nil {
		panic(err)
	}

	// Init Service
	svc, err := service.NewYogaService(dao)
	if err != nil {
		panic(err)
	}

	// Init Action
	a := actor.NewActionManager("FACEBOOK", svc)

	// Init Handler
	h := hg.NewWebHookHandler(a)

	router.HandleFunc("/refresh/cache", h.RefreshCache).Methods("POST")
	router.HandleFunc("/webhook", h.Serve).Methods("POST")

	// apache logging style
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	appengine.Main()
}
