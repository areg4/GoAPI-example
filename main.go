package main

import (
	"GoAPI/db"
	"GoAPI/routes"
	"GoAPI/types"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hola mundo")
	bd, err := db.GetDB()
	if err != nil {
		log.Printf("Error with database" + err.Error())
		return
	}
	bd.AutoMigrate(&types.Movie{})

	router := mux.NewRouter()
	routes.SetupRoutes(router)

	port := ":8080"

	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())

}
