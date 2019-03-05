package network

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	MaxUploadSize = 1024*1024*3
	UserId = "user_id"
	Sessionid = "ss_id"
)

type Result string

var(
	Success Result = "success"
	Error Result = "error"
	Refused Result = "refused"
)

type Response struct {
	Result Result	`json:"result"`
	Message string	`json:"message"`
	Time string	`json:"time"`
	Url string	`json:"url"`
}

func dispatcher(mux *http.ServeMux){
	mux.HandleFunc("/" ,StartHandle)
}

func writeRefused(writer http.ResponseWriter ,msg string ,req *http.Request){
	writeRes(writer ,Refused ,msg ,req)
}

func writeOk(writer http.ResponseWriter ,msg string ,req *http.Request){
	writeRes(writer ,Success ,msg ,req)
}

func writeErr(writer http.ResponseWriter ,msg string ,req *http.Request){
	writeRes(writer ,Error ,msg ,req)
}

func writeRes(writer http.ResponseWriter ,result Result ,msg string ,req *http.Request){
	writer.Header().Set("Content-Type" ,"application/json; charset=utf-8")
	switch result {
	case Refused:
		writer.WriteHeader(403)
	case Error:
		writer.WriteHeader(400)
	default:
		writer.WriteHeader(200)
	}

	res := &Response{
		Result:result,
		Message:msg,
		Url:req.URL.Path,
	}
	now := time.Now().String()
	index := strings.Index(now ,".")
	res.Time = now[:index]
	byts ,err := json.Marshal(res)
	if err != nil{
		log.Println(err)
		return
	}
	if _ ,err := writer.Write(byts);err != nil{
		log.Println(err)
	}
}
