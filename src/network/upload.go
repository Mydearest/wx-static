package network

import (
	"config"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func getRootPath() string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dir ,config.Args.StaticRootDir)
}


// POST	/		params:user_id token_id file
func Upload(req *http.Request) (string ,error){
	params := req.MultipartForm.Value
	userId := params[UserId][0]
	fileEntry := req.MultipartForm.File["file"][0]
	file ,err := fileEntry.Open()
	defer file.Close()
	if err != nil{
		return "" ,err
	}

	filename := fileEntry.Filename
	fileSaveId := getFileId(filename)

	if err := os.Mkdir(getSaveDir(userId), 0755);err != nil{
		return "" ,err
	}
	saveFile ,err := os.OpenFile(getSaveDir(userId)+"/"+fileSaveId, os.O_WRONLY|os.O_CREATE, 0755)
	defer saveFile.Close()

	if _ ,err := io.Copy(saveFile ,file);err != nil{
		log.Println(err)
		return "" ,err
	}
	return fmt.Sprintf("Upload successfully ,User - %s ,fileId - %s ,to get this file ,use 'GET static.shinoha.cn/%s/%s'" ,userId ,fileSaveId ,userId ,fileSaveId) ,nil
}


func getSaveDir(userid string) string{
	return getRootPath()+"/"+userid;
}

func getFileId(filename string) string{
	sha := sha256.New()
	sha.Write([]byte(filename))
	byteArr := sha.Sum(nil)
	return hex.EncodeToString(byteArr)[:10]
}

func updateDatabase(userId string ,filename string) error{

	return nil
}
