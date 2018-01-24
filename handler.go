package vn_holayoga_dialogflow_service

import (
	"vn.holayoga.dialogflow.service/actor"
	"io/ioutil"
	"google.golang.org/api/dialogflow/v2beta1"
	"encoding/json"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"vn.holayoga.dialogflow.service/utils"
)

type WebHookHandler struct {
	actionManager *actor.ActionManager // manage all actions
}

func NewWebHookHandler(actionManager *actor.ActionManager) *WebHookHandler {
	h := new(WebHookHandler)
	h.actionManager = actionManager

	return h
}

func (h *WebHookHandler) Serve(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx := appengine.NewContext(req)

	//TODO: only 1 handler for now. put into a middleware.
	h.actionManager.YogaCacheDao.RefreshCacheIfUninitialized(ctx)

	// 1. Read WebhookRequest
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	webhookReq := dialogflow.WebhookRequest{}
	err = json.Unmarshal(body, &webhookReq)
	utils.PrettyPrintWithCtx(ctx, "webhookReq:", webhookReq)
	log.Errorf(ctx, "webhookReq:", webhookReq)

	if err != nil {
		log.Errorf(ctx, "err unmarshal ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 2. call corresponding actionable service to process and get response
	webhookResp, err := h.actionManager.InvokeActionByName(actor.ActionName(webhookReq.QueryResult.Action), ctx, webhookReq)

	if err != nil {
		//TODO: if action not found, return funny message
		log.Errorf(ctx, "[error] action missing", err.Error())
		webhookResp = &dialogflow.WebhookResponse{
			FulfillmentText: "Chít, mama không bít phải làm gì, bạn đòi cái khác dễ chơi hơn coi :*",
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

func (h *WebHookHandler) RefreshCache(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	err := h.actionManager.YogaCacheDao.RefreshCache(ctx)
	if err != nil {
		log.Errorf(ctx, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusResetContent)
}
