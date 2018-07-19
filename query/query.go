package query

import (
	"strings"

	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
)

type Query struct {
	Token    string `json:"token"`
	Database string `json:"database"`
	Query    string `json:"query"`
}

type Answer map[string]interface{}

func Execute(query Query) (data interface{}, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return
	}
	_, err = conn.Exec("USE " + query.Database)
	if err != nil {
		return
	}
	query.Query = strings.TrimSpace(query.Query)
	splitQuery := strings.Split(query.Query, " ")
	var answer []Answer
	switch splitQuery[0] {
	case "SHOW", "SELECT":
		rows, err := conn.Queryx(query.Query)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			results := make(map[string]interface{})
			err = rows.MapScan(results)
			answer = append(answer, results)
		}
		return answer, nil
	default:
		_, err := conn.Exec(query.Query)
		if err != nil {
			return struct {
				Status int `json:"status"`
			}{Status: 200}, err
		}
	}
	return
}
