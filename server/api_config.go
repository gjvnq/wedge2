package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigS struct {
	ListenOn string
	DevMode  bool
	Port     string
	MySQL    struct {
		DB       string
		User     string
		Password string
	}
}

func LoadConfigFile() ConfigS {
	var dat ConfigS

	path := os.Getenv("CONFIG_FILE")
	if path == "" {
		path = "config.json"
	}

	byt, err := ioutil.ReadFile(path)
	if err != nil {
		Log.FatalF("Failed to open '%s' file: %s", path, err.Error())
	}
	err = json.Unmarshal(byt, &dat)
	if err != nil {
		Log.FatalF("Failed to parse '%s' file: %s", path, err.Error())
	}

	if dat.Port == "" {
		Log.FatalF("Config file '%s' MUST set the port", path)
	}
	return dat
}

func GetDBStringURI() string {
	return "" + Config.MySQL.User + ":" + Config.MySQL.Password + "@/" + Config.MySQL.DB + "?charset=utf8&parseTime=True&loc=Local"
}
