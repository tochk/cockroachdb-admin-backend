package tables

import "github.com/tochk/cockroachdb-admin-backend/connections_manager"

type Tables struct {
	Table string `db:"Table" json:"table"`
}

type Query struct {
	Token    string `json:"token"`
	Database string `json:"database"`
}

func GetTables(query Query) (tables []Tables, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return tables, err
	}
	_, err = conn.Exec("USE " + query.Database)
	if err != nil {
		return tables, err
	}
	err = conn.Select(&tables, "SHOW TABLES")
	return
}
