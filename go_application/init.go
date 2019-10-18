package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

// Configuration is a data structure for marshaling configuration from file convenience.
type Configuration struct {
	AppListIP        string `json:"app_list_ip"`
	AppPort          uint16 `json:"app_port" default:"8000"`
	AppAdminUser     string `json:"app_admin_user" default:"admin"`
	AppAdminPassword string `json:"app_admin_password" default:"admin"`
	AppAuthRealm     string `json:"app_auth_realm" default:"Log in!"`
	PGhost           string `json:"pg_host" default:"localhost"`
	PGport           uint16 `json:"pg_port" default:"5432"`
	PGuser           string `json:"pg_user" default:"root"`
	PGpassword       string `json:"pg_password"`
	PGdatabase       string `json:"pg_database"`
	PGmaxConnection  int    `json:"pg_max_con" default:"32"`
}

// Config contains setiings for application work.
var Config Configuration

func init() {
	configPath := flag.String("config", "mwh-config.json",
		"the path to configuration file")
	flag.Parse()
	configFile, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal("Open configuration file FAIL :: ", err)
	}
	if err := json.Unmarshal([]byte(configFile), &Config); err != nil {
		log.Fatal("Parse configuration file FAIL :: ", err)
	}
	log.Println(Config)
}

const (
	driverName = "postgres"
	dbSchema   = "postgres"
)

var db *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable", dbSchema,
		Config.PGuser, Config.PGpassword, Config.PGhost, Config.PGport, Config.PGdatabase)
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Database connection error :: ", err)
	}
	db.SetMaxOpenConns(Config.PGmaxConnection)
	if err = db.Ping(); err != nil {
		log.Fatal("Database error :: ", err)
	} else {
		log.Println("Success initialization")
	}
}
