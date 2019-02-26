package services

import (
	cache "countdown/countdown-api/repositories"
	"countdown/countdown-api/requests"
	"encoding/json"
	"fmt"
	"log"
)

func CreateNewTask(task requests.Task) error {

	taskToB, err := json.Marshal(task)

	if err != nil {
		return err
	}
	err = cache.Create("tasks", string(taskToB))

	if err != nil {
		return err
	}
	return nil
}

func GetAllTasks() ([]requests.Task, error) {

	payload, err := cache.GetAll("tasks")

	if err != nil {
		return nil, err
	}

	out := make([]requests.Task, len(payload))

	for index, val := range payload {
		var v, ok = val.([]byte)
		if ok {

			task := requests.Task{}
			err := json.Unmarshal(v, &task)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(task)
			out[index] = task

		}
	}

	return out, nil
}
