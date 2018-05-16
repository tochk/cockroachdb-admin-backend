package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tochk/cockroachdb-admin-backend/appError"
	"github.com/tochk/cockroachdb-admin-backend/databases"
)

func DatabasesHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var auth Auth
	err := decoder.Decode(&auth)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	db, err := databases.GetDatabases(auth.Token)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(3, "Get databases error", err))
		return
	}
	result, err := json.Marshal(db)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}
