package handlers

import (
	"api/db"
	"api/helper"
	"api/model"
	"api/session"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取请求体中的数据
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	// 将数据反序列化到结构体中
	ubody := &model.UserCredential{}
	err = json.Unmarshal(res, ubody)
	if err != nil {
		helper.SendErroeResponse(w, helper.ErrorRequestBodyParseFailed)
		return
	}

	// 进行数据库的操作
	err = db.AddUserCradential(ubody.UserName, ubody.PassWord)
	if err != nil {
		helper.SendErroeResponse(w, helper.ErrorDBError)
		return
	}

	// 创建一个新的sessionid，并将数据存储到数据库中
	id := session.GenerateNewSeesionId(ubody.UserName)
	// 将session信息进行序列化返回出去
	su := &model.SignedUp{Success: true, SessionId: id}

	resp, err := json.Marshal(su)
	if err != nil {
		helper.SendErroeResponse(w, helper.ErrorInternalFaults)
		return
	} else {
		helper.SendNormalResponse(w, string(resp), 200)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "登陆成功！")
}
