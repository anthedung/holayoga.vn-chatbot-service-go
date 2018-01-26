package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
	hg "vn.holayoga.dialogflow.service"
	hghandlers "vn.holayoga.dialogflow.service/handlers"
	"vn.holayoga.dialogflow.service/actor"
	"google.golang.org/appengine"
	"vn.holayoga.dialogflow.service/service/dao"
	"vn.holayoga.dialogflow.service/service"
	"flag"
	"vn.holayoga.dialogflow.service/utils"
)

var (
	config *hg.Config
	entity = flag.String("entity", "", "Entity to populate")
)

func init() {
	flag.Parse()
	config = hg.NewDefaultConfig()
	if len(*entity) > 0 {
		config.Datastore.CategoryKind = *entity
	}

	utils.DebugfPrettyPrint("config", config)
}

//TODO: Authenticate service
func main() {
	router := mux.NewRouter()

	// Init DAO for datastore without context
	// context to be set
	dao, err := dao.NewYogaCategoryCache(12, *config)
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
	h := hghandlers.NewWebHookHandler(a, *config)

	router.HandleFunc("/cache/refresh", h.RefreshCache).Methods("POST")
	router.HandleFunc("/cache/init", h.Init).Methods("POST")
	router.HandleFunc("/webhook", h.Serve).Methods("POST")

	// apache logging style
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	appengine.Main()
}
