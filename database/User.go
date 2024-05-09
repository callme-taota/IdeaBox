package database

import "idea-box/models"

func GetUserInfo(id int) (models.UserTable, error) {
	user := models.UserTable{}
	res := DbEngine.First(&user, id)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func GetUserList() ([]models.UserTable, error) {
	var user []models.UserTable
	res := DbEngine.Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
