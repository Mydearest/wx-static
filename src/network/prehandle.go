package network

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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
	// 上传大小检查 由于ParseMultipartForm性质问题，需要提交调用，往后不再调用该方法
	if err := req.ParseMultipartForm(MaxUploadSize);err != nil{
		return err
	}

	ssid := req.MultipartForm.Value[Sessionid]
	usrid := req.MultipartForm.Value[UserId]
	// 参数污染检查
	if len(ssid) != 1 || len(usrid) != 1{
		return errors.New("not allowed to set multiple params with same key at the same time")
	}
	// 参数缺失
	if ssid[0] == "" || usrid[0] == ""{
		return errors.New("not allowed to call this api with empty param")
	}
	// 暂时的有效性检查
	if ssid[0] != "810"{
		return errors.New("no permission to call the API")
	}
	return nil
}

func checkFile(req *http.Request) error{
	files := req.MultipartForm.File["file"]
	// 上传数量检查
	if len(files) != 1{
		return errors.New("not allowed to upload multiple files at the same time")
	}
	// 上传无文件名或空文件
	if files[0].Filename == "" || files[0].Size == 0{
		return errors.New("not allowed to upload empty file")
	}
	//	上传文件格式检查 ,仅允许jpg,png
	file := files[0]
	if !strings.HasSuffix(file.Filename ,".jpg") && !strings.HasSuffix(file.Filename ,".png"){
		return errors.New("upload file of this type are not supported,only [png ,jpg] are allowed")
	}
	return nil
}