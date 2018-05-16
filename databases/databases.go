package databases

import "github.com/tochk/cockroachdb-admin-backend/connections_manager"

type Db struct {
	Database string `db:"Database" json:"database"`
}

func GetDatabases(token string) (db []Db, err error) {
	conn, err := connections_manager.GetConnection(token)
	if err != nil {
		return db, err
	}
	err = conn.Select(&db, "SHOW DATABASES")
	return
}