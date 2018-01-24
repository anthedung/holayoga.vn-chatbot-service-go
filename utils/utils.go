package utils

import (
	"encoding/json"
	"context"
	"google.golang.org/appengine/log"
	"github.com/sirupsen/logrus"
)

// utility function for pretty print struct
func PrettyPrintWithCtx(ctx context.Context, title string, v interface{}) error {
	out, err := json.MarshalIndent(v, "", "    ")

	if err == nil {
		log.Infof(ctx, title+": "+string(out))
	}

	return err
}

func PrettyPrint(title string, v interface{}) error {
	out, err := json.MarshalIndent(v, "", "    ")

	if err == nil {
		logrus.Info(title + ": " + string(out))
	}

	return err
}
