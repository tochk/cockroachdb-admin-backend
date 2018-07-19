package tables

import "github.com/tochk/cockroachdb-admin-backend/connections_manager"

type CreateQuery struct {
	Token    string `json:"token"`
	Database string `json:"database"`
	Table    string `json:"table"`
}

//todo add normal create table
func CreateTable(query CreateQuery) (tables []Tables, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return tables, err
	}
	_, err = conn.Exec("USE " + query.Database)
	if err != nil {
		return tables, err
	}
	_, err = conn.Exec("CREATE TABLE " + query.Table)
	if err != nil {
		return tables, err
	}
	err = conn.Select(&tables, "SHOW TABLES")
	return
}
