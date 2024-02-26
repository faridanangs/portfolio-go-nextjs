package entities

type AuthUser struct {
	Id        string `gorm:"column:id;primaryKey"`
	Username  string `gorm:"column:username"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime"`
}

func (a *AuthUser) TableName() string {
	return "portfolio.users"
}

type AuthResponse struct {
	Id        string
	Username  string
	Email     string
	CreatedAt int64
}

type AuthSignUp struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,max=150"`
	Password string `json:"password" validate:"required"`
}

type AuthSignIn struct {
	Email    string `json:"email" validate:"required,max=150"`
	Password string `json:"password" validate:"required"`
}
