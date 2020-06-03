package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	Port            string
	DbUsername   string
	DbPassword string
	ConnectionString string
}

var Config Configuration

func LoadConfig(){
	Config = Configuration{}
	err := gonfig.GetConf("./config/config.json", &Config)
	Config.ConnectionString = "mongodb+srv://"+Config.DbUsername+":"+Config.DbPassword+"@test-cluster-dvxai.mongodb.net/meetup"
	if err!=nil {
		log.Fatal(err)
	}
}

