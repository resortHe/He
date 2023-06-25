package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` //父级评论
	ParentCommentID    *uint           `gorm:"size:10" json:"parent_comment_id"`                //父评论ID
	Content            string          `gorm:"size:256" json:"content"`                         //评论内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`              //点赞数
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`           //子评论数
	Article            ArticleModel    `json:"article"`                                         //关联的文章
	ArticleID          int             `gorm:"size:10" json:"article_id"`                       //文章ID
	User               UserMole        `json:"user"`                                            //关联的用户
	UserID             uint            `gorm:"size:10" json:"user_id"`                          //评论的用户
}
