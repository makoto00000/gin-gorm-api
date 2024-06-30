package repositories

import (
	"app/api/models"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	// GORMのFindメソッドを使用してusersテーブルからデータを取得する
	// 取得されたデータはusers変数にマッピングされる
	if err := db.Find(&users).Error; err != nil {
			// クエリ実行中にエラーが発生した場合、そのエラーを返す
		return nil, err
	}
	// エラーが発生しなかった場合、取得されたデータを返す
	return users, nil
}

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}