package controllers

import (
	"fmt"
	"net/http"
)
func GetAllTasks(w http.ResponseWriter, r *http.Request){
	fmt.Println("Executing allTasks func")
}

func CreateTask(w http.ResponseWriter , r *http.Request) {
		
}