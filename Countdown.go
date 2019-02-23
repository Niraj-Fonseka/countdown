package main

import (
	"countdown/controllers"
	"net/http"
)






func main(){

	http.HandleFunc("/alltasks" , controllers.GetAllTasks)
	http.ListenAndServe(":8080", nil)
}