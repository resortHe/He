package image_ser

import (
	"fmt"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/ctype"
	"go_vue/plugins/qinniu"
	"go_vue/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var WhiteImageList = []string{
	"jpg",
	"png",
	"jpeg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
}

type FileUploadResponse struct {
	FileName  string `json:"fileName"`  //文件名
	IsSuccess bool   `json:"isSuccess"` //是否上传成功
	Msg       string `json:"msg"`       //消息
}

// ImageUploadService 处理图片上传
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath
	//文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return

	}
	//判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size > float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2fMB,设定大小为:%dMB:", size, global.Config.Upload.Size)
		return
	}
	//读取文件内容，hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteDate, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteDate)
	//通过哈希值去数据库查图片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash=?", imageHash).Error
	if err == nil {
		//找到了
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType := ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	if global.Config.QiNiu.Enable {
		filePath, err = qinniu.UploadImage(byteDate, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu

	}
	//图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
