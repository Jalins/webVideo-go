package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"streamserver/helper"
	"streamserver/model"
	"time"

	"github.com/julienschmidt/httprouter"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	vid := p.ByName("vid-id")
	vl := model.VIDEO_DIR + vid
	fmt.Println("vl->", vl)

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("打开文件失败: %v", err)
		helper.SendErrorResponese(w, http.StatusInternalServerError, "内部错误")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")

	// response, request, name, time, file
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()

}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 限定能读到的最大文件的大小
	r.Body = http.MaxBytesReader(w, r.Body, model.MAX_UPLOAD_SIZE)
	// 校验从表单传过来的form表单是否超过限制
	err := r.ParseMultipartForm(model.MAX_UPLOAD_SIZE)
	if err != nil {

		helper.SendErrorResponese(w, http.StatusBadRequest, "文件过大")
		return
	}
	// 从form表单中获取文件
	file, _, err := r.FormFile("file")
	if err != nil {
		helper.SendErrorResponese(w, http.StatusInternalServerError, "内部错误")
		return
	}

	// 将文件读取进来转化成二进制
	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Printf("读取文件失败: %v", err)
		helper.SendErrorResponese(w, http.StatusInternalServerError, "内部错误")
	}

	fn := p.ByName("vid-id")

	// 将二进制写到目的文件中去，并设置文件的权限为0666
	err = ioutil.WriteFile(model.VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("写文件错误: %v", err)
		helper.SendErrorResponese(w, http.StatusInternalServerError, "内部错误")
		return
	}

	// 返回给前端状态码
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "上传成功")

}
