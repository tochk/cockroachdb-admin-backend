package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tochk/cockroachdb-admin-backend/appError"
	"github.com/tochk/cockroachdb-admin-backend/common"
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
	"github.com/tochk/cockroachdb-admin-backend/data"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	common.CORS(&w)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var query data.Query
	err := decoder.Decode(&query)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}

	tbl, err := data.GetData(query)
	if err != nil {
		if err == connections_manager.InvalidTokenError {
			fmt.Fprint(w, appError.GetJsonError(4, "Invalid token", err))
			return
		}
		fmt.Fprint(w, appError.GetJsonError(6, "Get data error", err))
		return
	}
	result, err := json.Marshal(tbl)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}
