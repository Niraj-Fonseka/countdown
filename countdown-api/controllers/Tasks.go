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
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var tsk requests.Task
		err := decoder.Decode(&tsk)
		if err != nil {
			payload := map[string]string{
				"message": "invalid request body",
			}
			http.Error(w, json.NewEncoder(w).Encode(payload), http.StatusBadRequest)
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

	} else if r.Method == "DELETE" {
		decoder := json.NewDecoder(r.Body)
		var tsk requests.DeleteTask

		err := decoder.Decode(&tsk)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = services.DeleteTask(&tsk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "Task deleted successfully")
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
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
