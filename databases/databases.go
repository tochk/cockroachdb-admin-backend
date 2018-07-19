package databases

import "github.com/tochk/cockroachdb-admin-backend/connections_manager"

type Database struct {
	Database string `db:"Database" json:"database"`
}

func GetDatabases(token string) (db []Database, err error) {
	conn, err := connections_manager.GetConnection(token)
	if err != nil {
		return db, err
	}
	err = conn.Select(&db, "SHOW DATABASES")
	return
}
