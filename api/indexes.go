package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tochk/cockroachdb-admin-backend/appError"
	"github.com/tochk/cockroachdb-admin-backend/common"
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
	"github.com/tochk/cockroachdb-admin-backend/indexes"
)

func IndexesHandler(w http.ResponseWriter, r *http.Request) {
	common.CORS(&w)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var query indexes.Query
	err := decoder.Decode(&query)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	indexList, err := indexes.GetIndexes(query)
	if err != nil {
		if err == connections_manager.InvalidTokenError {
			fmt.Fprint(w, appError.GetJsonError(4, "Invalid token", err))
			return
		}
		fmt.Fprint(w, appError.GetJsonError(11, "Get indexes error", err))
		return
	}
	result, err := json.Marshal(indexList)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}