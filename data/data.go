package data

import (
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
)

type Query struct {
	Token string `json:"token"`
	Db    string `json:"db"`
	Table string `json:"table"`
}

type Answer map[string]interface{}

func GetData(query Query) (answer []Answer, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec("USE " + query.Db)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Queryx("SELECT * FROM " + query.Table)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		answer = append(answer, results)
	}
	return
}
