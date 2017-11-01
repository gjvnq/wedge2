package main

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigS struct {
	MySQL struct {
		DB       string
		User     string
		Password string
	}
}

func GetDBStringURI() string {
	var dat ConfigS

	byt, err := ioutil.ReadFile("config.json")
	if err != nil {
		Log.FatalF("Failed to open config.json file: %s", err.Error())
	}
	err = json.Unmarshal(byt, &dat)
	if err != nil {
		Log.FatalF("Failed to parse config.json file: %s", err.Error())
	}
	return "" + dat.MySQL.User + ":" + dat.MySQL.Password + "@/" + dat.MySQL.DB + "?charset=utf8&parseTime=True&loc=Local"
}
