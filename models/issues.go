package models

type Issues struct {
	Id int64 `gorm:"column:id;primaryKey" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
}

func (Issues) TableName() string {
    return "issues"
}