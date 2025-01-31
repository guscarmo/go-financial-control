package models

type Categoria struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Category string `json:"category"`
}
