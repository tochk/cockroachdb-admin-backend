package indexes

import (
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
)

type Query struct {
	Token    string `json:"token"`
	Database string `json:"database"`
	Table    string `json:"table"`
}

type Index struct {
	Table     string `db:"Table" json:"table"`
	Name      string `db:"Name" json:"name"`
	Unique    bool   `db:"Unique" json:"unique"`
	Seq       int    `db:"Seq" json:"seq"`
	Column    string `db:"Column" json:"column"`
	Direction string `db:"Direction" json:"direction"`
	Storing   bool   `db:"Storing" json:"storing"`
	Implicit  bool   `db:"Implicit" json:"implicit"`
}

func GetIndexes(query Query) (indexes []Index, err error) {
	conn, err := connections_manager.GetConnection(query.Token)
	if err != nil {
		return
	}
	_, err = conn.Exec("USE " + query.Database)
	if err != nil {
		return
	}
	err = conn.Select(&indexes, "SHOW INDEXES FROM "+query.Table)
	if err != nil {
		return
	}
	return
}
