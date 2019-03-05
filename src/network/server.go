package network

import (
	"log"
	"net/http"
)

var Mux *http.ServeMux

func init(){
	Mux = http.NewServeMux()
}

func StartServer(){
	var server = http.Server{
		Handler:Mux,
		Addr:"127.0.0.1:8080"}
	dispatcher(Mux)
	if err := server.ListenAndServe();err != nil{
		log.Println(err.Error())
	}
}
