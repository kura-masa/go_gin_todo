package config

import (
	"database/sql"
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	SQLDriver, DbName, LogFile string
}

var Db *sql.DB
var err error
var Config ConfigList

// initの役割
// main.go内のinitはmain関数より先に読む

func init() {
	LoadConfig()
}
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
