package todos

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string
	Done  bool `gorm:"default:false"`
}
