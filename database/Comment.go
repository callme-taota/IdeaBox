package database

import (
	"idea-box/models"
	"idea-box/models/APIs/Response"
	"idea-box/tolog"
)

func CreateComment(ideaID, userID int, content string) (int, error) {
	comment := &models.CommentTable{
		Content: content,
		UserID:  userID,
		IdeaID:  ideaID,
	}
	result := DbEngine.Create(&comment)
	if result.Error != nil {
		tolog.Log().Infof("Error while CreateComment %e", result.Error).PrintAndWriteSafe()
		return -1, result.Error
	}
	commentID := comment.ID
	return commentID, nil
}

func DeleteComment(commentID int) (bool, error) {
	var comment models.CommentTable
	result := DbEngine.First(&comment, commentID)
	if result.Error != nil {
		tolog.Log().Infof("Error while DeleteComment %e", result.Error).PrintAndWriteSafe()
		return false, result.Error
	}
	result = DbEngine.Delete(&comment, commentID)
	if result.Error != nil {
		tolog.Log().Infof("Error while DeleteComment %e", result.Error).PrintAndWriteSafe()
		return false, result.Error
	}
	return true, nil
}

func GetCommentCountByIdeaID(ideaID int) (int, error) {
	var count int64
	res := DbEngine.Model(&models.CommentTable{}).Where("idea_id = ?", ideaID).Count(&count)
	if res.Error != nil {
		return -1, res.Error
	}
	return int(count), nil
}

func GetCommentByIdeaID(ideaID int) ([]Response.FullComment, error) {
	var comments []Response.FullComment
	res := DbEngine.Select("comment.*, user.name as UserName, user.color as UserColor").
		Joins("INNER JOIN user ON comment.user_id = user.id").
		Where("idea_id = ?", ideaID).
		Order("comment.created_at DESC").
		Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}
