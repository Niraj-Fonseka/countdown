package main

import (
	"countdown/controllers"
	"net/http"
	"fmt"
)






func main(){

	fmt.Println("running server")
	http.HandleFunc("/tasks" , controllers.HandleTasks)
	http.HandleFunc("/task" , controllers.HandleTask)
	http.HandleFunc("/", controllers.Hello)
	http.ListenAndServe(":8080", nil)
}