package handlers

import (
	"context"
	"portfolio/helpers"
	"portfolio/models/entities"
	"portfolio/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewSkillHandlerCreate(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	skill := entities.CreateAndUpdateSkill{}
	if err := c.BodyParser(&skill); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at body parser create skill", 400, "Bad Request", nil)
	}
	if err := validate.Struct(&skill); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at validate.struct create skill", 400, "Bad Request", nil)
	}

	// cloudinary proses
	ctx := context.Background()
	fileHeader, err := c.FormFile("image")
	helpers.PanicIfError(err, "error at formfile in create skill")
	cloudyRes := helpers.CreateImageToCloudinary(fileHeader, ctx)

	skillReq := entities.SkillEntities{
		Image:         cloudyRes.SecureURL,
		Name:          skill.Name,
		Category:      skill.Category,
		PublicIdImage: cloudyRes.PublicID,
	}

	skillResEntity, err := repositories.NewSkillRepository().Save(ctx, db, skillReq)
	if err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at skill.Repo.Save in create skill", 400, "Bad Request", nil)
	}
	skillRes := entities.SkillResponse{
		Id:        skillResEntity.Id,
		Image:     skillResEntity.Image,
		Name:      skillResEntity.Name,
		Category:  skillResEntity.Category,
		CreatedAt: skillResEntity.CreatedAt,
		UpdatedAt: skillResEntity.UpdatedAt,
	}
	return helpers.ResponseJson(c, "-", "create skill success", 200, "OK", skillRes)
}
func NewSkillHandlerUpdate(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	skillId := c.Params("id")
	skillParser := entities.CreateAndUpdateSkill{}
	if err := c.BodyParser(&skillParser); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at body parser update skill", 400, "Bad Request", nil)
	}
	if err := validate.Struct(&skillParser); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at validate.struct update skill", 400, "Bad Request", nil)
	}

	skillEntity := entities.SkillEntities{}
	if err := db.Model(&entities.SkillEntities{}).Where("id = ?", skillId).Take(&skillEntity).Error; err != nil {
		db.Rollback()
		return helpers.ResponseJson(c, err.Error(), "error at get data by id in hNDLER update", 404, "Not Found", nil)
	}

	ctx := context.Background()
	// cloudinary prosses deleteimage
	err := helpers.DeleteImageToCloudinary(skillEntity.PublicIdImage, ctx)
	if err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at delete image skill incloudinary handler", fiber.StatusBadRequest, "Bad request", nil)
	}
	// cloudinary proses save image
	fileHeader, err := c.FormFile("image")
	helpers.PanicIfError(err, "error at formfile in update skill")
	cloudyRes := helpers.CreateImageToCloudinary(fileHeader, ctx)

	// update skill
	skillReq := entities.SkillEntities{
		Id:            skillEntity.Id,
		Image:         cloudyRes.SecureURL,
		Name:          skillParser.Name,
		Category:      skillParser.Category,
		PublicIdImage: cloudyRes.PublicID,
	}

	skillResEntity, err := repositories.NewSkillRepository().Update(ctx, db, skillReq)
	if err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at skillRepo.Save in update skill", 400, "Bad Request", nil)
	}
	skillRes := entities.SkillResponse{
		Id:        skillResEntity.Id,
		Image:     skillResEntity.Image,
		Name:      skillResEntity.Name,
		Category:  skillResEntity.Category,
		CreatedAt: skillResEntity.CreatedAt,
		UpdatedAt: skillResEntity.UpdatedAt,
	}
	return helpers.ResponseJson(c, "-", "update skill success", 200, "OK", skillRes)
}

func NewSkillHandlerDelete(c *fiber.Ctx, db *gorm.DB) error {
	skillId := c.Params("id")
	skillEntity := entities.SkillEntities{}
	if err := db.Model(&entities.SkillEntities{}).Where("id = ?", skillId).Take(&skillEntity).Error; err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at handler get data all data skill", 404, "Not Found", nil)
	}
	ctx := context.Background()

	// cloudinary prosses deleteimage
	if err := helpers.DeleteImageToCloudinary(skillEntity.PublicIdImage, ctx); err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at delete image skill incloudinary skill", fiber.StatusBadRequest, "Bad request", nil)
	}

	if err := repositories.NewSkillRepository().Delete(ctx, db, skillEntity); err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at delete data skill in repository", 400, "Bad Request", nil)
	}

	return helpers.ResponseJson(c, "-", "delete skill success", 200, "OK", nil)
}

func NewSkillHandlerGetAll(c *fiber.Ctx, db *gorm.DB) error {
	skillEntity := []entities.SkillEntities{}
	if err := db.Model(&entities.SkillEntities{}).Find(&skillEntity).Error; err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at get data by id in handler get all data project", 404, "Not Found", nil)
	}
	projResponses := []entities.SkillResponse{}
	for _, data := range skillEntity {
		projResponse := entities.SkillResponse{
			Id:        data.Id,
			Image:     data.Image,
			Name:      data.Name,
			Category:  data.Category,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		projResponses = append(projResponses, projResponse)
	}
	return helpers.ResponseJson(c, "-", "get all data skill success", 200, "OK", projResponses)
}
