package utils

import (
	"encoding/json"
	"context"
	"google.golang.org/appengine/log"
	"github.com/sirupsen/logrus"
)

// utility function for pretty print struct
func DebugfPrettyPrintWithCtx(ctx context.Context, title string, v interface{}) error {
	out, err := json.MarshalIndent(v, "", "    ")

	if err == nil {
		log.Debugf(ctx, title+": "+string(out))
	}

	return err
}

func DebugfPrettyPrint(title string, v interface{}) error {
	out, err := json.MarshalIndent(v, "", "    ")

	if err == nil {
		logrus.Debug(title + ": " + string(out))
	}

	return err
}
