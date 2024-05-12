package Response

import "idea-box/models"

type FullComment struct {
	models.CommentTable
	UserName  string
	UserColor string
}
