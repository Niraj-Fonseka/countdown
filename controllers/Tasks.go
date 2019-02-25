package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"countdown/requests"
)
func HandleTasks(w http.ResponseWriter, r *http.Request){
	fmt.Println("Executing allTasks func")
}

func HandleTask(w http.ResponseWriter , r *http.Request) {
	
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

		
		return
	}else{
		http.Error(w, "Unsupported Request", http.StatusBadRequest)

	}


	}

	
	//marshal json

