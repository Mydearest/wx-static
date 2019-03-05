package network

import "net/http"

// POST	/upload		params:user_id token_id file
func Upload(req *http.Request) (string ,error){
	if err := req.ParseMultipartForm(MaxUploadSize);err != nil{
		return "" ,err
	}
	params := req.MultipartForm.Value
	user := params[UserId][0]
	ss := params[Sessionid][0]

	return user+"*"+ss ,nil
}
