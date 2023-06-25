package models

// MenuImageModel 自定义菜单和背景图的连接表，方便排序
type MenuImageModel struct {
	MenuID      int         `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    int         `json:"BannerID"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10" json:"Sort"`
}
