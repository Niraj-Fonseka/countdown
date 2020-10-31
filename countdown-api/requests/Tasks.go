package requests

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Task      string `json:"task"`
	Timestamp int64  `json:"timestamp"`
	Deadline  int64  `json:"deadline" gorm:"-"`
}

type DeleteTask struct {
	TaskID uint `json:"task_id"`
}
