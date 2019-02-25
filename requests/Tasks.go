package requests


type Task struct {
	Task string  `json:"task"`
	Deadline int `json:"deadline"`
}