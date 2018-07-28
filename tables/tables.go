package tables

import "github.com/tochk/cockroachdb-admin-backend/connections_manager"

type Tables struct {
	Table string `db:"Table" json:"table"`
}

type Query struct {
	Token    string `json:"token"`
	Database string `json:"database"`
}

type SchemaQuery struct {
	Token    string `json:"token"`
	Database string `json:"database"`
	Table    string `json:"table"`
}

type SchemaAnswer struct {
	Rows []map[string]interface{} `json:"rows"`
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

func GetSchema(query SchemaQuery) (answer SchemaAnswer, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return
	}
	_, err = conn.Exec("USE " + query.Database)
	if err != nil {
		return
	}
	q := "SELECT column_name, ordinal_position, column_default, is_nullable, data_type, character_maximum_length, character_octet_length, numeric_precision, numeric_scale, datetime_precision FROM kirino.information_schema.columns WHERE table_name = '" + query.Table + "'"
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
