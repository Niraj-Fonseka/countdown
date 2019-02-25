package services

import (
	"encoding/json"
	cache "countdown/repositories"
	"countdown/requests"
)

func CreateNewTask(task requests.Task) error{

	taskToB, err := json.Marshal(task)

	if err != nil {
		return err
	}
	err = cache.Create("tasks" , string(taskToB))

	if err != nil {
		return err
	}

	return nil
}