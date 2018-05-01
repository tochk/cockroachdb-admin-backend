package errors

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type err struct {
	Code  int
	Human string
	Err   string
}

var fatal = "{\"Code\":500,\"Human\":\"Fatal error\",\"Err\":\"idk what is happening\"}"

func GetJsonError(code int, human string, appError error) string {
	result, jsonError := json.Marshal(err{Code: code, Human: human, Err: appError.Error()})
	if jsonError != nil {
		log.Error(jsonError, appError)
		return fatal
	}
	return string(result)
}
