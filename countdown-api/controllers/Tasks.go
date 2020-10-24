package controllers

import (
	"countdown/countdown-api/requests"
	"countdown/countdown-api/services"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	fmt.Println(r.Method == "POST")
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var tsk requests.Task
		err := decoder.Decode(&tsk)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = services.CreateNewTask(&tsk)
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
	enableCors(&w)

	tasks, err := services.GetAllTasks()

	for i, t := range tasks {
		diff := t.Timestamp - time.Now().Unix()
		if diff < 0 {
			tasks[i].Deadline = 0
			continue
		}

		tasks[i].Deadline = t.Timestamp - time.Now().Unix()
	}

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
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
