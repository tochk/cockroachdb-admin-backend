package databases

import "github.com/tochk/cockroachdb-admin-backend/connections_manager"

type DropQuery struct {
	Token    string `json:"token"`
	Database string `json:"database"`
}

func DropDatabase(query DropQuery) (db []Database, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return db, err
	}
	_, err = conn.Exec("DROP DATABASE " + query.Database)
	if err != nil {
		return db, err
	}
	err = conn.Select(&db, "SHOW DATABASES")
	return
}
