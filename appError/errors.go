package appError

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type err struct {
	Code  int `json:"code"`
	Human string `json:"human"`
	Err   string `json:"err"`
}

var fatal = "{\"code\":500,\"human\":\"Fatal error\",\"err\":\"idk what is happening\"}"

func GetJsonError(code int, human string, appError error) string {
	result, jsonError := json.Marshal(err{Code: code, Human: human, Err: appError.Error()})
	if jsonError != nil {
		log.Error(jsonError, appError)
		return fatal
	}
	return string(result)
}
