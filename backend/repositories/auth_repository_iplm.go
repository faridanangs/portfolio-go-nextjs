package repositories

import (
	"context"
	"portfolio/models/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(ctx context.Context, db *gorm.DB, request entities.AuthUser) (entities.AuthUser, error)
}
type AuthRepositoryIPLM struct{}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryIPLM{}
}

func (repo *AuthRepositoryIPLM) Save(ctx context.Context, db *gorm.DB, request entities.AuthUser) (entities.AuthUser, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Model(&request).Create(&request).Error; err != nil {
		tx.Rollback()
		return entities.AuthUser{}, err
	}
	user := entities.AuthUser{}
	if err := tx.Model(&request).Where("id = ?", request.Id).Take(&user).Error; err != nil {
		tx.Rollback()
		return entities.AuthUser{}, err
	}
	tx.Commit()
	return user, nil
}
