package upload

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"main/sharelin/httppojo"
	"net/http"
	"os"
	"time"
)

type File struct {
	Name       string
	SourceName string
	Path       string
	CreateTime string
}

type FileHandler interface {
	HandUploadFile(file File)
}

type FileUpload struct {
	uploadPath string
}

func NewFileUpload(uploadPath string, uploadUrl string, fileHandler FileHandler) *FileUpload {
	http.Handle(uploadUrl, uploadFileHandler(uploadPath, fileHandler))
	return &FileUpload{uploadPath: uploadPath}
}

func uploadFileHandler(uploadPath string, fileHandler FileHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uploadResult := httppojo.ServerResponseBuilder.CreateSuccess("上传成功")
		defer func() {
			res, _ := json.Marshal(uploadResult)
			w.Write(res)
		}()

		reader, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			log.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())
			if part.FileName() == "" {
				data, _ := ioutil.ReadAll(part)
				fmt.Printf("FormData=[%s]\n", string(data))
			} else {
				timeObj := time.Now()
				date := timeObj.Format("2006-01-02")
				savePath := uploadPath + "/" + date
				mkDirIfNotExists(savePath)
				fileName := timeObj.Format("20060102150405") + "-" + part.FileName()
				dst, err := os.Create(savePath + "/" + fileName)
				if err != nil {
					log.Printf("create file error %s \n", err.Error())
					uploadResult = httppojo.ServerResponseBuilder.CreateError("上传失败:" + err.Error())
					return
				}
				defer dst.Close()
				_, err = io.Copy(dst, part)
				if err != nil {
					log.Printf("save file error %s \n", err.Error())
					uploadResult = httppojo.ServerResponseBuilder.CreateError("上传失败:" + err.Error())
					return
				}
				dateTime := timeObj.Format("2006-01-02 15:04:05")
				file := File{CreateTime: dateTime, Path: savePath, SourceName: part.FileName(), Name: fileName}
				fileHandler.HandUploadFile(file)
			}
		}

	}
}

func mkDirIfNotExists(path string) {
	_, err := os.Stat(path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Printf("auto create  folder %s success\n", path)
		} else {
			log.Printf("auto create  folder %s fail\n", path)
		}
	}
}
