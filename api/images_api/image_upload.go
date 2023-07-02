package images_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models/res"
	"go_vue/service"
	"go_vue/service/image_ser"
	"io/fs"
	"os"
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

// ImageUploadView 上传多个图片并返回上传结果
// @Tags 图片管理
// @Summary 上传图片
// @Description  上传多个图片并返回上传结果
// @Param images formData file true "图片文件"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//上传多个图片
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("不存在的文件", c)
		return
	}
	//判断路径是否存在
	//不存在就创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		//递归创建
		err := os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		//上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		//成功
		if !global.Config.QiNiu.Enable {
			//本地还得保存
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)

	}
	res.OkWithData(resList, c)
}
