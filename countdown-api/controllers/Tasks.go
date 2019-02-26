package controllers

import (
	"countdown/countdown-api/requests"
	"countdown/countdown-api/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method == "POST")
	if r.Method == "POST" {
		fmt.Println("Method is equal to post")
		decoder := json.NewDecoder(r.Body)
		var tsk requests.Task
		err := decoder.Decode(&tsk)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = services.CreateNewTask(tsk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "Task created successfully")

			return
		}

	} else {
		http.Error(w, "Unsupported Request", http.StatusBadRequest)

	}

}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := services.GetAllTasks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(tasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(bytes))

}

//marshal json
