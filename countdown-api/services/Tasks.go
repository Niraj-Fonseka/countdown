package services

import (
	"countdown/countdown-api/repositories"
	"countdown/countdown-api/requests"
)

func CreateNewTask(task *requests.Task) error {

	err := repositories.CreateTask(task)

	if err != nil {
		return err
	}

	return nil
}

func GetAllTasks() ([]requests.Task, error) {

	payload, err := repositories.GetTasks()

	if err != nil {
		return nil, err
	}

	return payload, nil
}

func DeleteTask(t *requests.DeleteTask) error {
	return repositories.DeleteTask(t)

}
