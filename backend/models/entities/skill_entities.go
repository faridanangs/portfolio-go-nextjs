package entities

type SkillEntities struct {
	Id            string `gorm:"column:id;primaryKey;autoIncrement"`
	Image         string `gorm:"column:image"`
	Name          string `gorm:"column:name"`
	Category      string `gorm:"column:category"`
	PublicIdImage string `gorm:"column:public_id_image"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (a *SkillEntities) TableName() string {
	return "portfolio.skills"
}

type SkillResponse struct {
	Id        string `json:"id"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateAndUpdateSkill struct {
	Id       string
	Image    string `form:"image"`
	Name     string `validate:"required,max=100,min=5" form:"name"`
	Category string `validate:"required" form:"category"`
}
