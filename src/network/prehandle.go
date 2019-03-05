package network

import (
	"errors"
	"fmt"
	"net/http"
)

func StartHandle(writer http.ResponseWriter, req *http.Request){

	if err := checkMethod(req);err != nil{
		writeRefused(writer ,err.Error() ,req)
		return
	} else if err := checkPermission(req);err != nil{
		writeRefused(writer ,err.Error() ,req)
		return
	}else if err := checkFile(req);err != nil{
		writeErr(writer ,err.Error() ,req)
		return
	}else{
		if msg ,err := Upload(req);err != nil{
			writeErr(writer ,err.Error() ,req)
		}else {
			writeOk(writer ,msg ,req)
		}
	}
}

// 仅允许post
func checkMethod(req *http.Request) error{
	if req.Method != http.MethodPost{
		return errors.New(fmt.Sprintf("%s is not allowed" ,req.Method))
	}
	return nil
}

// user_id ,ss_id
// 参数污染/参数缺失/参数格式非法/参数无效
func checkPermission(req *http.Request) error{
	return nil
}

func checkFile(req *http.Request) error{
	return nil
}