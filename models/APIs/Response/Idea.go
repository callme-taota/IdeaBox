package Response

import "idea-box/models"

type IdeaWithUser struct {
	IdeaID int `json:"IdeaID"`
	models.IdeaTable
	models.UserTable
}
