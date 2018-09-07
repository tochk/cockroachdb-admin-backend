package tables

import (
	"strings"

	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
)

type CreateQuery struct {
	Token    string `json:"token"`
	Database string `json:"database"`
	Table    string `json:"table"`
	Fields   []struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Length  string `json:"length"`
		Key     string `json:"key"`
		Default string `json:"default"`
		Null    bool   `json:"null"`
	} `json:"fields"`
	Constraints []struct {
		Name  string `json:"name"`
		Check string `json:"check"`
	} `json:"constraints"`
	Indexes []struct {
		Name   string `json:"name"`
		Fields string `json:"fields"`
	} `json:"indexes"`
}

func CreateTable(query CreateQuery) (tables []Tables, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return tables, err
	}
	_, err = conn.Exec("USE " + query.Database)
	if err != nil {
		return tables, err
	}

	q := "CREATE TABLE " + query.Table + "(\n"
	for _, e := range query.Fields {
		q += e.Name + " " + e.Type + " "
		if e.Length != "" {
			q += "(" + e.Length + ") "
		}
		if e.Key != "" {
			q += e.Key + " "
		}
		if !e.Null {
			q += "NOT NULL "
		}
		if e.Default != "" {
			q += "DEFAULT " + e.Default
		}
		q += ",\n"
	}
	for _, e := range query.Constraints {
		q += "CONSTRAINT " + e.Name + " CHECK " + e.Check + ",\n"
	}
	for _, e := range query.Indexes {
		q += "INDEX " + e.Name + " (" + e.Fields + "),\n"
	}
	q = strings.Trim(q, ",\n")
	q += "\n);"

	_, err = conn.Exec(q)
	if err != nil {
		return tables, err
	}
	err = conn.Select(&tables, "SHOW TABLES")
	return
}
