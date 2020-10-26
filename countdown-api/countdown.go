package main

import (
	"countdown/countdown-api/controllers"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting app ..")
	http.HandleFunc("/tasks", controllers.GetAllTasks)
	http.HandleFunc("/task", controllers.HandleTask)
	http.HandleFunc("/", controllers.Hello)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
