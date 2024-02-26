package repositories

import (
	"context"
	"portfolio/models/entities"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Save(ctx context.Context, db *gorm.DB, request entities.ProjectEntities) (entities.ProjectEntities, error)
	Update(ctx context.Context, db *gorm.DB, request entities.ProjectEntities) (entities.ProjectEntities, error)
	Delete(ctx context.Context, db *gorm.DB, request entities.ProjectEntities) error
}
type ProjectRepositoryIPLM struct{}

func NewProjectRepository() ProjectRepository {
	return &ProjectRepositoryIPLM{}
}

func (repo *ProjectRepositoryIPLM) Save(ctx context.Context, db *gorm.DB, request entities.ProjectEntities) (entities.ProjectEntities, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Model(&entities.ProjectEntities{}).Create(&request).Error; err != nil {
		tx.Rollback()
		return entities.ProjectEntities{}, err
	}
	project := entities.ProjectEntities{}
	if err := tx.Model(&entities.ProjectEntities{}).Where("id = ?", request.Id).Take(&project).Error; err != nil {
		tx.Rollback()
		return entities.ProjectEntities{}, err
	}
	tx.Commit()
	return project, nil
}

func (repo *ProjectRepositoryIPLM) Update(ctx context.Context, db *gorm.DB, request entities.ProjectEntities) (entities.ProjectEntities, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Model(&entities.ProjectEntities{}).Where("id = ?", request.Id).Updates(&request).Error; err != nil {
		tx.Rollback()
		return entities.ProjectEntities{}, err
	}
	project := entities.ProjectEntities{}
	if err := tx.Model(&entities.ProjectEntities{}).Where("id = ?", request.Id).Take(&project).Error; err != nil {
		tx.Rollback()
		return entities.ProjectEntities{}, err
	}
	tx.Commit()
	return project, nil
}

func (repo *ProjectRepositoryIPLM) Delete(ctx context.Context, db *gorm.DB, request entities.ProjectEntities) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Model(&entities.ProjectEntities{}).Where("id = ?", request.Id).Delete(&request).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
