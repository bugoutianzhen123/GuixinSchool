package main

import (
	"GuiXinSchool/config"
	"GuiXinSchool/route"
	"flag"
)

type App struct{
	*route.Engine
}

var flagConfig string

func init() {
	flag.StringVar(&flagConfig, "conf", "../config", "config path, eg: -conf config.yaml")
}


func main(){
	var ac config.AppConfig

	if err := config.Load(&ac, flagConfig); err != nil {
		panic(err)
	}


}