package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tochk/cockroachdb-admin-backend/appError"
	"github.com/tochk/cockroachdb-admin-backend/common"
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
	 "github.com/tochk/cockroachdb-admin-backend/query"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	common.CORS(&w)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var q query.Query
	err := decoder.Decode(&q)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	indexList, err := query.Execute(q)
	if err != nil {
		if err == connections_manager.InvalidTokenError {
			fmt.Fprint(w, appError.GetJsonError(4, "Invalid token", err))
			return
		}
		fmt.Fprint(w, appError.GetJsonError(12, "Execute query error", err))
		return
	}
	result, err := json.Marshal(indexList)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}
