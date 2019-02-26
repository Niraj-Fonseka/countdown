package main

import (
	"countdown/countdown-api/controllers"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("running server")
	http.HandleFunc("/tasks", controllers.GetAllTasks)
	http.HandleFunc("/task", controllers.HandleTask)
	http.HandleFunc("/", controllers.Hello)
	http.ListenAndServe(":8080", nil)
}
