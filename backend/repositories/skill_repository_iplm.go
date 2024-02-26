package repositories

import (
	"context"
	"portfolio/models/entities"

	"gorm.io/gorm"
)

type SkillRepository interface {
	Save(ctx context.Context, db *gorm.DB, request entities.SkillEntities) (entities.SkillEntities, error)
	Update(ctx context.Context, db *gorm.DB, request entities.SkillEntities) (entities.SkillEntities, error)
	Delete(ctx context.Context, db *gorm.DB, request entities.SkillEntities) error
}
type SkillRepositoryIPLM struct{}

func NewSkillRepository() SkillRepository {
	return &SkillRepositoryIPLM{}
}

func (repo *SkillRepositoryIPLM) Save(ctx context.Context, db *gorm.DB, request entities.SkillEntities) (entities.SkillEntities, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Model(&entities.SkillEntities{}).Create(&request).Error; err != nil {
		tx.Rollback()
		return entities.SkillEntities{}, err
	}
	project := entities.SkillEntities{}
	if err := tx.Model(&entities.SkillEntities{}).Where("id = ?", request.Id).Take(&project).Error; err != nil {
		tx.Rollback()
		return entities.SkillEntities{}, err
	}
	tx.Commit()
	return project, nil
}

func (repo *SkillRepositoryIPLM) Update(ctx context.Context, db *gorm.DB, request entities.SkillEntities) (entities.SkillEntities, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Model(&entities.SkillEntities{}).Where("id = ?", request.Id).Updates(&request).Error; err != nil {
		tx.Rollback()
		return entities.SkillEntities{}, err
	}
	project := entities.SkillEntities{}
	if err := tx.Model(&entities.SkillEntities{}).Where("id = ?", request.Id).Take(&project).Error; err != nil {
		tx.Rollback()
		return entities.SkillEntities{}, err
	}
	tx.Commit()
	return project, nil
}

func (repo *SkillRepositoryIPLM) Delete(ctx context.Context, db *gorm.DB, request entities.SkillEntities) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Model(&entities.SkillEntities{}).Where("id = ?", request.Id).Delete(&request).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
