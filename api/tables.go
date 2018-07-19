package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tochk/cockroachdb-admin-backend/appError"
	"github.com/tochk/cockroachdb-admin-backend/common"
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
	"github.com/tochk/cockroachdb-admin-backend/tables"
)

func TablesHandler(w http.ResponseWriter, r *http.Request) {
	common.CORS(&w)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var query tables.Query
	err := decoder.Decode(&query)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}

	tbl, err := tables.GetTables(query)
	if err != nil {
		if err == connections_manager.InvalidTokenError {
			fmt.Fprint(w, appError.GetJsonError(4, "Invalid token", err))
			return
		}
		fmt.Fprint(w, appError.GetJsonError(5, "Get tables error", err))
		return
	}
	result, err := json.Marshal(tbl)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}

func CreateTableHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var query tables.CreateQuery
	err := decoder.Decode(&query)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}

	tbl, err := tables.CreateTable(query)
	if err != nil {
		if err == connections_manager.InvalidTokenError {
			fmt.Fprint(w, appError.GetJsonError(4, "Invalid token", err))
			return
		}
		fmt.Fprint(w, appError.GetJsonError(7, "Create table error", err))
		return
	}
	result, err := json.Marshal(tbl)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}

func DropTableHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var query tables.DropQuery
	err := decoder.Decode(&query)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}

	tbl, err := tables.DropTable(query)
	if err != nil {
		if err == connections_manager.InvalidTokenError {
			fmt.Fprint(w, appError.GetJsonError(4, "Invalid token", err))
			return
		}
		fmt.Fprint(w, appError.GetJsonError(8, "Drop table error", err))
		return
	}
	result, err := json.Marshal(tbl)
	if err != nil {
		fmt.Fprint(w, appError.GetJsonError(1, "Parsing json error", err))
		return
	}
	fmt.Fprint(w, string(result))
}