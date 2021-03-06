package main

import (
	"flag"
	"net/http"

	"github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
	"github.com/tochk/cockroachdb-admin-backend/api"
	"github.com/tochk/cockroachdb-admin-backend/configuration"
	"github.com/tochk/cockroachdb-admin-backend/connections_manager"
)

var (
	configFile  = flag.String("config", "conf.json", "Where to read the config from")
	servicePort = flag.String("port", ":5001", "Service port number")
)

func main() {
	log.Info("Starting application")
	flag.Parse()
	log.Info("Flags parsed")
	connections_manager.Init()
	log.Info("Initializing connection manager successful")
	if err := configuration.LoadConfig(*configFile); err != nil {
		log.Fatal(err)
	}
	log.Info("Configuration file loaded")

	http.HandleFunc("/api/connect/", api.ConnectHandler)

	http.HandleFunc("/api/databases/", api.DatabasesHandler)
	http.HandleFunc("/api/databases/create/", api.CreateDatabaseHandler)
	http.HandleFunc("/api/databases/drop/", api.DropDatabaseHandler)

	http.HandleFunc("/api/tables/", api.TablesHandler)
	http.HandleFunc("/api/tables/create/", api.CreateTableHandler)
	http.HandleFunc("/api/tables/drop/", api.DropTableHandler)
	http.HandleFunc("/api/tables/schema/", api.TableSchemaHandler)

	http.HandleFunc("/api/data/", api.DataHandler)

	http.HandleFunc("/api/indexes/", api.IndexesHandler)

	http.HandleFunc("/api/query/", api.QueryHandler)

	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))

	log.Info("Starting listen connections on ", *servicePort)
	http.ListenAndServe(*servicePort, nil)
}
