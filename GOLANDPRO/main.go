package main

import (
	"GOLANDPRO/database"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	database.DB()
	r := router.SetupRoutes()
	log.Println("Server running on port 8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
