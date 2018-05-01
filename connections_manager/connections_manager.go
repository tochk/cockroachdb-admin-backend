package connections_manager

import "github.com/jmoiron/sqlx"

var Connections map[string]*sqlx.DB

func Init() {
	Connections = make(map[string]*sqlx.DB)
}

func Connect(login, password string) {

}
