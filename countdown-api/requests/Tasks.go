package requests

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Task     string `json:"task"`
	Deadline int    `json:"deadline"`
}
