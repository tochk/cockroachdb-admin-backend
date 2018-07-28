package data

import (
	"strconv"

	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
)

type Query struct {
	Token    string `json:"token"`
	Database string `json:"database"`
	Table    string `json:"table"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type Answer struct {
	Count int                      `json:"count"`
	Rows  []map[string]interface{} `json:"rows"`
}

func GetData(query Query) (answer Answer, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return
	}
	_, err = conn.Exec("USE " + query.Database)
	if err != nil {
		return
	}
	var count int
	countQuery := "SELECT COUNT (*) FROM " + query.Table
	conn.Get(&count, countQuery)
	answer.Count = count
	q := "SELECT * FROM " + query.Table
	if query.Limit != 0 {
		q += " LIMIT " + strconv.Itoa(query.Limit)
	}
	if query.Offset != 0 {
		q += " OFFSET " + strconv.Itoa(query.Offset)
	}
	rows, err := conn.Queryx(q)
	if err != nil {
		return
	}
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		answer.Rows = append(answer.Rows, results)
	}
	return
}
