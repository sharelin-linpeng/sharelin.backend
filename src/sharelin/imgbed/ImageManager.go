package imgbed

import (
	"main/sharelin/httppojo"
	"main/sharelin/upload"
	"net/http"
)

var resource = "/images"

var manager = newImageManager()

type imageManager struct {
}

func newImageManager() *imageManager {
	http.Handle(resource+"/query", query())
	http.Handle(resource+"/remove", remove())
	return &imageManager{}
}

func query() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res = httppojo.NewServerResponse()
		defer httppojo.WriteResponse(&w, res)

		list := ImageDataBase.queryFileList()
		if list == nil {
			newList := make([]ImageFile, 0)
			list = &newList
		}
		res.CreateSuccessData("查询成功", list)

	}
}

func remove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res = httppojo.NewServerResponse()
		defer httppojo.WriteResponse(&w, res)
		id := r.FormValue("id")
		err := ImageDataBase.deleteOne(id)
		if err != nil {
			res.CreateError("删除失败")
			return
		}
		res.CreateSuccess("删除成功")

	}
}

const (
	UploadPath = "/Users/sharelin/GolandProjects/sharelin.backend/static/images/upload"
	UploadUrl  = "/uploadImage"
)

type ImageFileHandler struct {
}

func NewImageFileHandler() *ImageFileHandler {
	return &ImageFileHandler{}
}

func (receiver ImageFileHandler) HandUploadFile(file upload.File) {
	imageFile := ImageFile{}
	imageFile.Name = file.Name
	imageFile.Path = file.Path[len("/Users/sharelin/GolandProjects/sharelin.backend/static"):] + "/" + file.Name

	imageFile.SourceName = file.SourceName
	imageFile.CreateTime = file.CreateTime
	ImageDataBase.saveFileInfo(imageFile)
}

var imageFileHandler = NewImageFileHandler()

var imageUpload = upload.NewFileUpload(UploadPath, resource+UploadUrl, imageFileHandler)
