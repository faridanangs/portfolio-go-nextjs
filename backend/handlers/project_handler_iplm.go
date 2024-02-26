package handlers

import (
	"context"
	"portfolio/helpers"
	"portfolio/models/entities"
	"portfolio/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewProjectHandlerCreate(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	project := entities.CreateAndUpdateProject{}
	if err := c.BodyParser(&project); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at body parser create project", 400, "Bad Request", nil)
	}
	if err := validate.Struct(&project); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at validate.struct create project", 400, "Bad Request", nil)
	}

	// cloudinary proses
	ctx := context.Background()
	fileHeader, err := c.FormFile("image")
	helpers.PanicIfError(err, "error at formfile in create project")
	cloudyRes := helpers.CreateImageToCloudinary(fileHeader, ctx)

	project.Id = uuid.NewString()
	project.Image = cloudyRes.SecureURL

	projReq := entities.ProjectEntities{
		Id:            project.Id,
		Image:         project.Image,
		Title:         project.Title,
		Description:   project.Description,
		Tech:          project.Tech,
		PublicIdImage: cloudyRes.PublicID,
	}

	projResEntity, err := repositories.NewProjectRepository().Save(ctx, db, projReq)
	if err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at projectRepo.Save in create project", 400, "Bad Request", nil)
	}
	projRes := entities.ProjectResponse{
		Id:          projResEntity.Id,
		Image:       projResEntity.Image,
		Title:       projResEntity.Title,
		Description: projResEntity.Description,
		Tech:        projResEntity.Tech,
		CreatedAt:   projResEntity.CreatedAt,
		UpdatedAt:   projResEntity.UpdatedAt,
	}
	return helpers.ResponseJson(c, "-", "create project success", 200, "OK", projRes)
}
func NewProjectHandlerUpdate(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	projectId := c.Params("id")
	projectParser := entities.CreateAndUpdateProject{}
	if err := c.BodyParser(&projectParser); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at body parser update project", 400, "Bad Request", nil)
	}
	if err := validate.Struct(&projectParser); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at validate.struct update project", 400, "Bad Request", nil)
	}

	projectEntity := entities.ProjectEntities{}
	if err := db.Model(&entities.ProjectEntities{}).Where("id = ?", projectId).Take(&projectEntity).Error; err != nil {
		db.Rollback()
		return helpers.ResponseJson(c, err.Error(), "error at get data by id in handler project update", 404, "Not Found", nil)
	}

	ctx := context.Background()
	// cloudinary prosses deleteimage
	err := helpers.DeleteImageToCloudinary(projectEntity.PublicIdImage, ctx)
	if err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at delete image project incloudinary", fiber.StatusBadRequest, "Bad request", nil)
	}
	// cloudinary proses save image
	fileHeader, err := c.FormFile("image")
	helpers.PanicIfError(err, "error at formfile in update project")
	cloudyRes := helpers.CreateImageToCloudinary(fileHeader, ctx)

	// update project
	projectEntity.Image = cloudyRes.SecureURL
	projectEntity.Description = projectParser.Description
	projectEntity.Title = projectParser.Title
	projectEntity.Tech = projectParser.Tech

	projReq := entities.ProjectEntities{
		Id:            projectEntity.Id,
		Image:         projectEntity.Image,
		Title:         projectEntity.Title,
		Description:   projectEntity.Description,
		Tech:          projectEntity.Tech,
		PublicIdImage: cloudyRes.PublicID,
	}

	projResEntity, err := repositories.NewProjectRepository().Update(ctx, db, projReq)
	if err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at projectRepo.Save in update project", 400, "Bad Request", nil)
	}
	projRes := entities.ProjectResponse{
		Id:          projResEntity.Id,
		Image:       projResEntity.Image,
		Title:       projResEntity.Title,
		Description: projResEntity.Description,
		Tech:        projResEntity.Tech,
		CreatedAt:   projectEntity.CreatedAt,
		UpdatedAt:   projResEntity.UpdatedAt,
	}
	return helpers.ResponseJson(c, "-", "update project success", 200, "OK", projRes)
}

func NewProjectHandlerDelete(c *fiber.Ctx, db *gorm.DB) error {
	projectId := c.Params("id")
	projectEntity := entities.ProjectEntities{}
	if err := db.Model(&entities.ProjectEntities{}).Where("id = ?", projectId).Take(&projectEntity).Error; err != nil {
		db.Rollback()
		return helpers.ResponseJson(c, err.Error(), "error at get data by id in handler project delete", 404, "Not Found", nil)
	}
	ctx := context.Background()

	// cloudinary prosses deleteimage
	if err := helpers.DeleteImageToCloudinary(projectEntity.PublicIdImage, ctx); err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at delete image project incloudinary", fiber.StatusBadRequest, "Bad request", nil)
	}

	if err := repositories.NewProjectRepository().Delete(ctx, db, projectEntity); err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at delete data project in repository", 400, "Bad Request", nil)
	}

	return helpers.ResponseJson(c, "-", "delete project success", 200, "OK", nil)
}

func NewProjectHandlerGetAllProject(c *fiber.Ctx, db *gorm.DB) error {
	projectEntity := []entities.ProjectEntities{}
	if err := db.Model(&entities.ProjectEntities{}).Find(&projectEntity).Error; err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at get data all data in handler project", 404, "Not Found", nil)
	}
	projResponses := []entities.ProjectResponse{}
	for _, data := range projectEntity {
		projResponse := entities.ProjectResponse{
			Id:          data.Id,
			Image:       data.Image,
			Title:       data.Title,
			Description: data.Description,
			Tech:        data.Tech,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		}
		projResponses = append(projResponses, projResponse)
	}
	return helpers.ResponseJson(c, "-", "get all data project success", 200, "OK", projResponses)
}
