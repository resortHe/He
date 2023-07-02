package menu_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (MenuApi) MenuNameList(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	res.OkWithData(menuNameList, c)
}
