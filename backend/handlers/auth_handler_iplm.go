package handlers

import (
	"context"
	"os"
	"portfolio/helpers"
	"portfolio/models/entities"
	"portfolio/repositories"
	"portfolio/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewAuthHandlerCreate(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	userSingIn := entities.AuthSignUp{}
	if err := c.BodyParser(&userSingIn); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at body parser auth-user SignUp", 400, "Bad Request", nil)
	}
	if err := validate.Struct(&userSingIn); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at validate.struct auth-user SignUp", 400, "Bad Request", nil)
	}
	authUser := entities.AuthUser{
		Id:       uuid.NewString(),
		Username: userSingIn.Username,
		Email:    userSingIn.Email,
		Password: utils.GeneratePassword(userSingIn.Password),
	}
	authUserRes, err := repositories.NewAuthRepository().Save(context.Background(), db, authUser)
	if err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at authRepo.Save auth-user SignUp", 400, "Bad Request", nil)
	}
	claims := jwt.MapClaims{
		"id":         authUserRes.Id,
		"username":   authUserRes.Username,
		"email":      authUserRes.Email,
		"created_at": authUserRes.CreatedAt,
		"exp":        -1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(os.Getenv("JWT_SECREET_KEY")))
	helpers.PanicIfError(err, "error in signin at token.SignedString auth-handler")
	return helpers.ResponseToken(c, 200, "Ok", strToken)

}

func NewAuthHandlerDelete(c *fiber.Ctx, db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			c.Status(200).JSON(fiber.Map{
				"Error": err,
			})
		}
	}()
	userId := c.Params("id")
	authUser := entities.AuthUser{}
	if err := tx.Model(&entities.AuthUser{}).Where("id = ?", userId).Take(&authUser).Error; err != nil {
		tx.Rollback()
		return helpers.ResponseJson(c, err.Error(), "error at get data by id in auth-delete", 404, "Not Found", nil)
	}
	if err := tx.Model(&entities.AuthUser{}).Delete(&authUser).Error; err != nil {
		tx.Rollback()
		return helpers.ResponseJson(c, err.Error(), "error at delete in auth-delete", 500, "InternalServerError", nil)
	}
	tx.Commit()
	return helpers.ResponseJson(c, "-", "delete acount success", 200, "OK", nil)
}

func NewAuthHandlerSignIn(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	userSingIn := entities.AuthSignIn{}
	if err := c.BodyParser(&userSingIn); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at body parser auth-user SignUp", 400, "Bad Request", nil)
	}
	if err := validate.Struct(&userSingIn); err != nil {
		return helpers.ResponseJson(c, err.Error(), "Error at validate.struct auth-user SignUp", 400, "Bad Request", nil)
	}
	authUser := entities.AuthUser{}
	if err := db.Model(&authUser).Where("email = ?", userSingIn.Email).Take(&authUser).Error; err != nil {
		return helpers.ResponseJson(c, err.Error(), "error at get data by id in auth-signup", 404, "Not Found", nil)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(userSingIn.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return helpers.ResponseJson(c, err.Error(), "Password doesn't match in auth-signup", 401, "Unauthorized", nil)
		}
		return helpers.ResponseJson(c, err.Error(), "Error at compare passwor in auth-signup", 500, "Internal Server Error", nil)
	}
	claims := jwt.MapClaims{
		"id":         authUser.Id,
		"username":   authUser.Username,
		"email":      authUser.Email,
		"created_at": authUser.CreatedAt,
		"exp":        time.Now().Add(168 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(os.Getenv("JWT_SECREET_KEY")))
	helpers.PanicIfError(err, "error in signin at token.SignedString auth-handler")
	return helpers.ResponseToken(c, 200, "Ok", strToken)
}
