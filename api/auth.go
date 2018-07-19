package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tochk/cockroachdb-admin-backend/appError"
	"github.com/tochk/cockroachdb-admin-backend/common"
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
	"github.com/tochk/cockroachdb-admin-backend/user"
)

type Auth struct {
	Token string `json:"token"`
}

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	common.CORS(&w)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var usr user.User
	err := decoder.Decode(&usr)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	token, err := connections_manager.Connect(usr)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(2, "Connection to database error", err))
		return
	}
	result, err := json.Marshal(Auth{Token: token})
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}
