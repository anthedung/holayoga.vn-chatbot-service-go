package main

import (
	"flag"
	"google.golang.org/appengine/log"
	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"encoding/json"
	"net/http"
	"google.golang.org/api/dialogflow/v2beta1"
	"io/ioutil"
	"vn.holayoga.dialogflow.service/actor"
	"golang.org/x/net/context"
	"github.com/gorilla/handlers"
	"os"
)

//command line
var (
	port = flag.String("port", "8080", "Port for holayoga dialogflow service")
	conf = flag.String("conf", "holayoga.conf", "Configuration file")
)

func init() {
	flag.Parse()
}

func main() {
	// Routing
	router := mux.NewRouter()
	h := NewWebHookHandler()
	router.HandleFunc("/webhook", h.Serve).Methods("POST")
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	appengine.Main()
}

// [Handler]
func NewWebHookHandler() *WebHookHandler {
	handler := new(WebHookHandler)

	return handler
}

type WebHookHandler struct {
}

func (h *WebHookHandler) Serve(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx := appengine.NewContext(req)

	// 1. Read WebhookRequest
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	webhookReq := dialogflow.WebhookRequest{}
	PrettyPrint(ctx, "webhookReq", webhookReq)

	err = json.Unmarshal(body, &webhookReq)
	if err != nil {
		log.Errorf(ctx, "err unmarshal ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 2. call corresponding actionable service to process and get response
	var webhookResp dialogflow.WebhookResponse
	if actionFunc := actor.GetActorByAction(actor.Action(webhookReq.QueryResult.Action)); actionFunc != nil {
		webhookResp = actionFunc(webhookReq)
	} else {
		//TODO: if service not found, return funny message
		log.Errorf(ctx, "err unmarshal ", err.Error())
		webhookResp = dialogflow.WebhookResponse{
			FulfillmentText: "[error] action missing " + err.Error(),
		}
	}

	// 3. return processed response
	webhookRespByte, err := json.Marshal(webhookResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(webhookRespByte)
}

func (h *WebHookHandler) ServeTest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	webhookResp := dialogflow.WebhookResponse{
		FulfillmentText: "Day la hola mama, chao ban 2!",
	}

	webhookRespByte, _ := json.Marshal(webhookResp)

	w.WriteHeader(http.StatusOK)
	w.Write(webhookRespByte)
}

// utility function for pretty print struct
func PrettyPrint(ctx context.Context, title string, v interface{}) error {
	out, err := json.MarshalIndent(v, "", "    ")

	if err == nil {
		log.Infof(ctx, title+": "+string(out))
	}

	return err
}
