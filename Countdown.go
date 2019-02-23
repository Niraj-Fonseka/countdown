package main

import (
	"countdown/controllers"
	"net/http"
)






func main(){

	http.HandleFunc("/alltasks" , controllers.AllTasks)
	http.ListenAndServe(":8080", nil)
}