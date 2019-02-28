package main

import (
	"countdown/countdown-api/controllers"
	"fmt"
	"net/http"
	"github.com/rs/cors"

)

func main() {
    mux := http.NewServeMux()

	fmt.Println("running server")
	http.HandleFunc("/tasks", controllers.GetAllTasks)
	http.HandleFunc("/task", controllers.HandleTask)
	http.HandleFunc("/", controllers.Hello)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)
}
