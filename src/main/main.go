package main

import (
	"config"
	"log"
	"network"
)

func init(){
	log.SetFlags(log.Llongfile)
}

func main(){
	cmd := config.ParseCmd()
	if cmd.Help{
		config.Help()
	}else {
		network.StartServer()
	}
}
