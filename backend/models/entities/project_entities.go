package entities

type ProjectEntities struct {
	Id            string `gorm:"column:id;primaryKey"`
	Image         string `gorm:"column:image"`
	Title         string `gorm:"column:title"`
	Description   string `gorm:"column:description"`
	Tech          string `gorm:"column:tech"`
	PublicIdImage string `gorm:"column:public_id_image"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (a *ProjectEntities) TableName() string {
	return "portfolio.projects"
}

type ProjectResponse struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tech        string `json:"tech"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type CreateAndUpdateProject struct {
	Id          string
	Image       string `form:"image"`
	Title       string `validate:"required,max=100,min=5" form:"title"`
	Description string `validate:"required" form:"description"`
	Tech        string `validate:"required,max=300,min=5" form:"tech"`
}
