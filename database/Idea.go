package database

import (
	"idea-box/models"
	"idea-box/models/APIs/Response"
	"idea-box/tolog"
)

func CreateIdea(name, description string, author int) error {
	idea := models.IdeaTable{
		IdeaName:    name,
		Description: description,
		Author:      author,
	}
	result := DbEngine.Create(&idea)
	if result.Error != nil {
		tolog.Log().Infof("Error while CreateIdea: %e", result.Error).PrintAndWriteSafe()
		return result.Error
	}
	return nil
}

func GetIdeaCount() (int, error) {
	var count int64
	result := DbEngine.Model(&models.IdeaTable{}).Count(&count)
	if result.Error != nil {
		tolog.Log().Infof("Error while GetIdeaCount: %e", result.Error).PrintAndWriteSafe()
		return 0, result.Error
	}
	return int(count), nil
}

func GetIdeaList() ([]Response.IdeaWithUser, error) {
	var ideas []Response.IdeaWithUser
	result := DbEngine.Model(&models.IdeaTable{}).
		Select("idea.id as IdeaID,idea.*, user.name, user.color").
		Joins("JOIN user ON idea.author = user.id").
		Order("idea.created_at DESC").
		Scan(&ideas)
	if result.Error != nil {
		tolog.Log().Infof("Error while GetIdeaList: %e", result.Error).PrintAndWriteSafe()
		return nil, result.Error
	}
	return ideas, nil
}

func DeleteIdea(id int) error {
	result := DbEngine.Delete(&models.IdeaTable{}, id)
	if result.Error != nil {
		tolog.Log().Infof("Error while DeleteIdea: %e", result.Error).PrintAndWriteSafe()
		return result.Error
	}
	return nil
}

func UpdateIdeaName(id int, newName string) error {
	result := DbEngine.Model(&models.IdeaTable{}).Where("id = ?", id).Update("idea_name", newName)
	if result.Error != nil {
		tolog.Log().Infof("Error while UpdateIdeaName: %e", result.Error).PrintAndWriteSafe()
		return result.Error
	}
	return nil
}

func UpdateIdeaDesc(id int, newDesc string) error {
	result := DbEngine.Model(&models.IdeaTable{}).Where("id = ?", id).Update("description", newDesc)
	if result.Error != nil {
		tolog.Log().Infof("Error while UpdateIdeaDesc: %e", result.Error).PrintAndWriteSafe()
		return result.Error
	}
	return nil
}

func UpdateIdea(id int, newName, newDesc string, author int) error {
	result := DbEngine.Model(&models.IdeaTable{}).Where("id = ?", id).Updates(map[string]interface{}{"idea_name": newName, "description": newDesc, "author": author})
	if result.Error != nil {
		tolog.Log().Infof("Error while UpdateIdea: %e", result.Error).PrintAndWriteSafe()
		return result.Error
	}
	return nil
}
